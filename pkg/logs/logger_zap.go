package logs

import(
	"time"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"github.com/streadway/amqp"
	"context"
)

func Logger() *zap.Logger{

	cfg := zap.Config{
        Encoding:    "json", //encode kiểu json hoặc console
        Level:       zap.NewAtomicLevelAt(zap.InfoLevel),	//chọn InfoLevel có thể log ở cả 3 level
        OutputPaths: []string{"stderr"},

        EncoderConfig: zapcore.EncoderConfig{	//Cấu hình logging, sẽ không có stacktracekey
            MessageKey: "message",
						TimeKey: "time",
            LevelKey: "level",
						CallerKey:    "caller",
        		EncodeCaller: zapcore.FullCallerEncoder,	//Lấy dòng code bắt đầu log
			EncodeLevel: CustomLevelEncoder,	//Format cách hiển thị level log
			EncodeTime: SyslogTimeEncoder,	//Format hiển thị thời điểm log
        },
    }

    logger, _ := cfg.Build()
		return logger
}

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
    enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
    enc.AppendString("[" + level.CapitalString() + "]")
}

func SugarLog() *zap.SugaredLogger{
	writerSyncer := getLogWriter()
	encoder := getEncoder()

	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)

	logger := zap.New(core)
	return logger.Sugar()
}

func SugarLogWithRabitMQ(ctx *context.Context, rabbitmqConn *amqp.Connection) *zap.SugaredLogger{
	rabbitmqEncoder, err := newRabbitMQEncoder(rabbitmqConn)
	if err != nil {
		panic(err.Error())
	}
	writerSyncer := zapcore.AddSync(file)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)

	logger := zap.New(core)
	defer func(){
		select {
		case <-ctx.Done():
			return rabbitmqEncoder.Flush()
		}
	}()
	return logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getLogWriter() zapcore.WriteSyncer {
	env = os.Getenv("environment")
	if env == "" {
		env = "development"
	}
	nameFileLog = fmt.Sprinf("./%s.log",env)
	file, _ := os.Create(nameFileLog)
	return zapcore.AddSync(file)
}


type rabbitmqEncoder struct {
	conn *amqp.Connection
	channel *amqp.channel
	queue amqp.Queue
}
func newRabbitMQEncoder(conn *amqp.Connection) (*rabbitmqEncoder, error) {
	channel, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open RabbitMQ channel: %w", err)
	}
	queue, err := channel.QueueDeclare(
		"zap-log",
		false,
		false,
		false,
		false,
		nil
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare RabbitMQ queue: %w", err)
	}
	return &rabbitmqEncoder{
		conn: conn,
		channel: channel,
		queue: queue
	}, nil
}

func(e *rabbitmqEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) ([]byte, error){
	// Encode log entry data to JSON
	data , err := json.Marschal(zapcore.StructuredEntry(entry, fields))
	if err != nil {
		return nil , err
	}

	// Send message to RabbitMQ queue
	err = e.channel.Pushlis(
		"", // Exchange name ( using default exchange for fanout)
		e.queue.Name,
		false,
		false,
		amqp.Publishing{
			Body: data,
		}
		
	)
	if err != nil {
		return nil, fmt.Errorf("Faile to publish message to RabbitMQ:%w",err)
	}
	return nil, nil
}

func(e *rabbitmqEncoder) Flush() error {
	return e.channel.Close()
}
package apiserver

import (
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIserver почти главная структура для управления сервером
type APIserver struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	red    *redis.Client
}

// New struct принимает структуру конфиг и возвращает заполненную структуру apiServer
func New(config *Config) *APIserver {
	return &APIserver{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
		red:    NewRedis(),
	}
}

// NewRedis функция для определения значений реддиса
func NewRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

}

// Start Apiserver ...
func (s *APIserver) Start() error {
	if err := s.configurLogger(); err != nil {
		return err
	}
	// запускаем метод с горутинами
	s.configureRouter()
	s.logger.Info("Starting API server")
	// запускаем реддис первый раз использую простите если что-то неправильно сделал, надеюсь научите меня на работе как правильно юзать ))))))))
	s.logger.Info("++++++++++========== Starting Reddis Server")
	// добавляем в реддис ключ значение
	s.red.Set("counter", 0, 0)

	//  запускаем метод dlya прослушивания адресов
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIserver) configurLogger() error {
	//  парсим значение LOGlevel для дальнейшего присвоения переменной, нужную структуру
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	// задаём логгеру формат его работы (debug,error,info и тд. ) в данном случае debug
	s.logger.SetLevel(level)
	// конфигурация логгирования окончена +-
	return nil
}

func (s *APIserver) configureRouter() {
	// запускаем горутину-хэндлер
	s.router.HandleFunc("/rest/substr/find", s.handleSubString())
	s.router.HandleFunc("/rest/email/check", s.handleEmail())
	s.router.HandleFunc("/rest/counter/add/{i}", s.handleCountAdd())
	s.router.HandleFunc("/rest/counter/sub/{i}", s.handleCountSub())
	s.router.HandleFunc("/rest/counter/val/", s.handleCounterValue())
	s.router.HandleFunc("/rest/user", s.handleCreateUser())
}

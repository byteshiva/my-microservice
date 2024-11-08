package service

type Service struct{}

func NewService() *Service {
    return &Service{}
}

func (s *Service) GetGreeting(name string) string {
    return "Hello, " + name
}


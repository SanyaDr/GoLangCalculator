package config

const (
	DefaultLaunchPort = "8080" // Порт по умолчанию для запуска сервера

	CheckNewExpression_Timeout = 300 // Timeout для повторного обращения к Оркестратору (в миллисекундах)
)

// Названия переменных окружения системы
const (
	stockTimeExecution  = 10 // Время выполнение операции по умолчанию (если переменные среды не найдены)
	stockComputingPower = 3  // Кол-во горутин по умолчанию

	time_addition_name       = "TIME_ADDITION_MS"        // Время выполнения операции сложения в миллисекундах
	time_subtraction_name    = "TIME_SUBTRACTION_MS"     // Время выполнения операции вычитания в миллисекундах
	time_multiplication_name = "TIME_MULTIPLICATIONS_MS" // Время выполнения операции умножения в миллисекундах
	time_division_name       = "TIME_DIVISIONS_MS"       // Время выполнения операции деления в миллисекундах
	computing_power_name     = "COMPUTING_POWER"         // Количество горутин регулируется переменной среды
)

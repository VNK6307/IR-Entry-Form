package models

type QuestionType int

const (
	InputText QuestionType = iota
	InlineChoice
)

type Question struct {
	Key     string
	Text    string
	Type    QuestionType
	Options []string
}

func GetQuestions() []Question {
	return []Question{
		{"FullName", "👤 Введите Фамилию и Имя:", InputText, nil},
		{"CarBrand", "🚗 Введите марку и модель автомобиля:", InputText, nil},
		{"Group", "🏁 Выберите зачетную группу:", InlineChoice,
			[]string{"Standard 90", "Standard 130", "Standard 170", "Street 210", "Sport 2000", "Unlim", "Классика", "Спортпрототип"}},
		{"Classification", "👥 Выберите классификацию:", InlineChoice,
			[]string{"Junior", "Junior, до 16 лет", "Lady"}},
		{"EngineVolume", "⚙️ Введите объем двигателя:", InputText, nil},
		{"MaxPower", "💪 Введите максимальную мощность (л.с.):", InputText, nil},
		{"Weight", "⚖️ Введите вес автомобиля (кг):", InputText, nil},
		{"Turbo", "🌪️ Есть ли наддув?", InlineChoice,
			[]string{"Да", "Нет"}},
		{"TireModel", "🏎️ Введите марку и модель шин:", InputText, nil},
		{"Phone", "📱 Введите номер телефона:", InputText, nil},
	}
}

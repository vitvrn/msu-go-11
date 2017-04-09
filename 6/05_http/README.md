go test -coverprofile=cover.out - генерим отчет о coverage

go tool cover -html=cover.out -o cover.html - переводим отчет в html

go tool cover -func=cover.out - покрытие по функциям
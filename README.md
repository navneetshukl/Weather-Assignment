# Weather Application

This Go project fetches weather data for a given latitude and longitude using the OpenWeatherMap API.

## Features

- **GetForm**: This function renders a form page where users can enter latitude and longitude.
- **GetTemperature**: This function fetches the weather data for the given latitude and longitude from the OpenWeatherMap API.

## Usage

1. Clone the repository to your local machine.
2. Run `go get` to download the necessary Go packages.
3. Create a `.env` file in the root directory of the project, and add your OpenWeatherMap API key like so: `WEATHER_API_KEY=your_api_key_here`.
4. Run the server with `go run main.go`.

## Endpoints

- **GET /**: Renders a form for the user to enter latitude and longitude.
- **POST /**: Takes the latitude and longitude from the form, fetches the corresponding weather data, and renders it on a new page.

## Dependencies

- Gin: A HTTP web framework written in Go.
- godotenv: A Go package that reads from `.env` and adds them to `os.Getenv()`.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

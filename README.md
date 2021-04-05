# Dfinity Weather Oracle Example

[Dfinity Oracles](https://github.com/hyplabs/dfinity-oracles) is a framework for building [blockchain oracles](https://en.wikipedia.org/wiki/Blockchain_oracle) for the [Internet Computer](https://dfinity.org/).

Dfinity Weather Oracle Example is a sample project using Dfinity Oracles to retrieve the current weather from a number of different weather APIs, summarize them, and write the results to a software canister.

[Dfinity Crypto Oracle Example](https://github.com/hyplabs/dfinity-crypto-oracle) is another sample project, using Dfinity Oracles to retrieve current ETH and BTC prices.

Also see the [Dfinity Oracles tutorial](https://github.com/hyplabs/dfinity-oracles/blob/main/docs/tutorial.md) for step-by-step instructions about writing oracles from scratch.

## Quickstart

To build the weather oracle:

```bash
go build
```

To run the weather oracle, you'll need API keys for [WeatherAPI](https://www.weatherapi.com/), [WeatherBit](https://www.weatherbit.io/api), and [OpenWeatherMap](https://openweathermap.org/api). After obtaining these, you can then provide them to the oracle via environment variables:

```bash
export WEATHERAPI_API_KEY='WEATHERAPI_API_KEY_GOES_HERE'
export WEATHERBIT_API_KEY='WEATHERBIT_API_KEY_GOES_HERE'
export OPENWEATHERMAP_API_KEY='OPENWEATHERMAP_API_KEY_GOES_HERE'
./dfinity-weather-oracle
```

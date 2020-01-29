# mystock
My personal portfolio stock tracking.

* Configured from local configuration file (see example provided) or from an s3 object
* Using [QUANDL](https://www.quandl.com) web services to get the daily quotes.
* Relevant stock prices are saved in local cache (sqlite db)
* Regular checks are conducted on portfolio evolution
* Alerts and notification (SMS and/or stdout) are triggered if needed



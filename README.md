# Objective

Get more familiar with Go channels and Go software architecture in general.
I would love to receive any comments and suggestions as pull requests or issues how to construct this code better.

# Repository contents

For this objective I wrote a small program: daemon capable to schedule data querying from multiple sources and send this data as notifications over imaginary channel.

# Implementation requirements

* Each notification source has its own scheduling mechanism: from where data is fetched how frequently it is fetched and so on.
* Depending on data source notifications delivery status might be required to be reported. 

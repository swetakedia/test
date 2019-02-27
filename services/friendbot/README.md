# Friendbot Service for the Test Test Network

This calls out to testhorizon to submit the transaction

TestHorizon needs to be started with the following command line param: --friendbot-url="http://localhost:8004/"
This will forward any query params received against /friendbot to the friendbot instance.
The ideal setup for testhorizon is to proxy all requests to the /friendbot url to the friendbot service

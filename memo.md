gdpr.timeouts_ms.init_vendorlist_fetches: Timeout for the initial GVL fetch, expressed in milliseconds. This is what you're seeing when the application starts.

gdpr.timeouts_ms.active_vendorlist_fetch: Timeout for the subsequent GVL fetches while PBS is running, expressed in milliseconds.

GET https://vendor-list.consensu.org/vendorlist.json returned 403.
https://github.com/prebid/prebid-server/issues/1689 


Transparency and Consent String with Global Vendor & CMP List Formats
https://github.com/InteractiveAdvertisingBureau/GDPR-Transparency-and-Consent-Framework/blob/master/TCFv2/IAB%20Tech%20Lab%20-%20Consent%20string%20and%20vendor%20list%20formats%20v2.md#where-can-i-access-the-global-vendor-list



https://wiki.e-planning.net/wiki/spaces/ENG/pages/21874591/Prebid+Server
https://e-planning.atlassian.net/wiki/spaces/ENG/pages/21874596/Prebid+Server+in+AMP+Sites


passing entire PBS cookie in user.id field of request
https://github.com/prebid/prebid-server/issues/589

How solve it? gdpr.default_value #1897
https://github.com/prebid/prebid-server/issues/1897
PBS_GDPR_TIMEOUTS_MS_INIT_VENDORLIST_FETCHES
PBS_GDPR_TIMEOUTS_MS_ACTIVE_VENDORLIST_FETCH
pbs.yaml
gdpr:
timeouts_ms:
init_vendorlist_fetches: 5000
active_vendorlist_fetch: 5000

* https://github.com/prebid/prebid-server/issues/690
  https://github.com/prebid/prebid-server/issues/984
https://github.com/prebid/Prebid.js/issues/8668
  https://github.com/prebid/prebid.github.io/blob/master/faq/prebid-server-faq.md

## 설정 ##
* https://github.com/spf13/viper#why-viper



## 테스트 ##
*** https://github.com/prebid/prebid-server/issues/1409

https://go.dev/doc/articles/race_detector
https://github.com/prebid/prebid-server/blob/master/adapters/adapterstest/test_json.go

https://docs.prebid.org/prebid-server/endpoints/pbs-endpoint-overview.html
https://github.com/prebid/prebid-server-java/releases
https://docs.prebid.org/prebid-server/endpoints/openrtb2/pbs-endpoint-auction.html
https://github.com/prebid/prebid-server/blob/master/docs/developers/stored-requests.md
https://docs.prebid.org/prebid-server/features/pbs-storedreqs.html
https://github.com/prebid/prebid-server-java/pull/2038/files
https://docs.prebid.org/prebid-server/endpoints/pbs-endpoint-event.html
https://developers.smaato.com/marketers/openrtb-2-5-specifications/
https://support.google.com/authorizedbuyers/answer/7174589?hl=en
https://docs.prebid.org/prebid-mobile/prebid-mobile-pbs.html
https://docs.prebid.org/examples/video/server/ooyala/pbs-ve-ooyala.html
https://github.com/prebid/prebid-server-java/blob/master/README.md
https://docs.prebid.org/prebid-server/endpoints/openrtb2/pbs-endpoint-amp.html



https://github.com/prebid/prebid-server/issues/350
https://docs.prebid.org/dev-docs/modules/pubCommonId.html
https://docs.prebid.org/prebid-server/developers/pbs-cookie-sync.html#devquickstart



https://github.com/prebid/prebid-server/issues/1341
https://github.com/prebid/prebid-server-java/issues/338
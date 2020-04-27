#!/usr/bin/env bash

#swagger generate server -f ../spec/MEF_api_geographicAddressManagement_3.0.0.json -t ../ -s api/geographic_address_management -m models/geographic_address_management
#swagger generate server -f ../spec/MEF_api_geographicSiteManagement_3.0.0.json -t ../ -s api/geographic_site_management -m models/geographic_site_management
#swagger generate server -f ../spec/MEF_api_productInventoryManagement_3.0.0.json -t ../ -s api/product_inventory_management -m models/product_inventory_management
#swagger generate server -f ../spec/MEF_api_productOfferingQualificationManagement_3.0.1.json -t ../ -s api/product_offering_qualification_management -m models/product_offering_qualification_management
#swagger generate server -f ../spec/MEF_api_productOrderManagement_3.0.1.json -t ../ -s api/product_order_management -m models/product_order_management
#swagger generate server -f ../spec/MEF_api_productOrderNotification_3.0.0.json -t ../ -s api/product_order_notification -m models/product_order_notification
#swagger generate server -f ../spec/MEF_api_quoteManagement_2.0.0.json -t ../ -s api/quote_management -m models/quote_management
#swagger generate server -f ../spec/MEF_api_quoteNotification_1.0.0.yaml -t ../ -s api/quote_notification -m models/quote_notification
swagger generate server -f ../spec/all.yaml -t ../ -A sonata
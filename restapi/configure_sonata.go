// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"time"

	"github.com/brianvoe/gofakeit/v5"
	"github.com/go-openapi/strfmt"
	"github.com/jinzhu/gorm"

	"github.com/qlcchain/go-sonata-server/restapi/handler"
	"github.com/qlcchain/go-sonata-server/util"

	"github.com/qlcchain/go-sonata-server/restapi/handler/db"

	"github.com/qlcchain/go-sonata-server/schema"

	"github.com/dgrijalva/jwt-go"

	"github.com/go-openapi/swag"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"

	"github.com/qlcchain/go-sonata-server/config"

	"github.com/qlcchain/go-sonata-server/restapi/handler/mock"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/gorilla/handlers"
	"github.com/justinas/alice"

	"github.com/qlcchain/go-sonata-server/auth"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/restapi/operations"
)

const (
	sites = `[{"additionalSiteInformation":"Phased","description":"Legacy","fieldedAddress":{"city":"Anjou, QC","country":"Canada","geographicSubAddress":[{"buildingName":"Cologix MTL4","id":"brkvff18d3b2mjrlmorg","levelNumber":"BASEMENT 1","levelType":"1","privateStreetName":"university","privateStreetNumber":"01","subUnit":[{"subUnitIdentifier":"1237770961","subUnitType":"UNIT"}]}],"id":"brkvff18d3b2mjrlmos0","locality":"Locality","postcode":"H1N 3M2","stateOrProvince":"Idaho","streetName":"7171 Jean Talon East","streetNr":"841","streetSuffix":"fort","streetType":"Alley"},"formattedAddress":{"addrLine1":"841  fort , 7171 Jean Talon East fort","city":"Anjou, QC","country":"Canada","locality":"Locality","postcode":"H1N 3M2","stateOrProvince":"Idaho"},"geographicLocation":{"geographicPoint":[{"id":"brkvff18d3b2mjrlmp0g","latitude":"-39.868629","longitude":"-27.665208"}],"id":"brkvff18d3b2mjrlmp00","spatialRef":"AU"},"id":"brkvff18d3b2mjrlmp10","referencedAddress":{"id":"brkvff18d3b2mjrlmou0","referenceId":"brkvff18d3b2mjrlmoug","referenceType":"refer"},"relatedParty":[{"@referredType":"Organization","emailAddress":"gilbertokuneva@hartmann.io","id":"brkvff18d3b2mjrlmp2g","name":"Brody Walker","number":"9058835105","role":["Buyer","Seller","Site Contact"]}],"siteCompanyName":"Verdafero","siteCustomerName":"Toy Lueilwitz","siteName":"Brandy Grady","siteType":"PUBLIC","status":"planned"},{"additionalSiteInformation":"modular","description":"Customer","fieldedAddress":{"city":"Baie d Urfe, QC","country":"Canada","geographicSubAddress":[{"buildingName":"ROOT MTL-R2","id":"brkvff18d3b2mjrlmp30","levelNumber":"BASEMENT 1","levelType":"1","privateStreetName":"university","privateStreetNumber":"01","subUnit":[{"subUnitIdentifier":"377280272","subUnitType":"UNIT"}]}],"id":"brkvff18d3b2mjrlmp3g","locality":"Locality","postcode":"H9X3T1","stateOrProvince":"Massachusetts","streetName":"19701 Avenue Clark Graham","streetNr":"73193","streetSuffix":"burgh","streetType":"Alley"},"formattedAddress":{"addrLine1":"73193  burgh , 19701 Avenue Clark Graham burgh","city":"Baie d Urfe, QC","country":"Canada","locality":"Locality","postcode":"H9X3T1","stateOrProvince":"Massachusetts"},"geographicLocation":{"geographicPoint":[{"id":"brkvff18d3b2mjrlmp80","latitude":"-84.659170","longitude":"-179.314598"}],"id":"brkvff18d3b2mjrlmp7g","spatialRef":"MV"},"id":"brkvff18d3b2mjrlmp8g","referencedAddress":{"id":"brkvff18d3b2mjrlmp5g","referenceId":"brkvff18d3b2mjrlmp60","referenceType":"refer"},"relatedParty":[{"@referredType":"Organization","emailAddress":"santospowlowski@daugherty.name","id":"brkvff18d3b2mjrlmpa0","name":"Wallace Bins","number":"1960369125","role":["Buyer","Seller","Site Contact"]}],"siteCompanyName":"Quertle","siteCustomerName":"Godfrey Emmerich","siteName":"Yvette Hansen","siteType":"PUBLIC","status":"planned"},{"additionalSiteInformation":"radical","description":"International","fieldedAddress":{"city":"Boucherville, QC","country":"Canada","geographicSubAddress":[{"buildingName":"eStruxture MTL-3 (fka. Kolotek)","id":"brkvff18d3b2mjrlmpag","levelNumber":"BASEMENT 1","levelType":"1","privateStreetName":"university","privateStreetNumber":"01","subUnit":[{"subUnitIdentifier":"1658235342","subUnitType":"UNIT"}]}],"id":"brkvff18d3b2mjrlmpb0","locality":"Locality","postcode":"J4B 5H3","stateOrProvince":"Alaska","streetName":"1350 Nobel St","streetNr":"1445","streetSuffix":"land","streetType":"Alley"},"formattedAddress":{"addrLine1":"1445  land , 1350 Nobel St land","city":"Boucherville, QC","country":"Canada","locality":"Locality","postcode":"J4B 5H3","stateOrProvince":"Alaska"},"geographicLocation":{"geographicPoint":[{"id":"brkvff18d3b2mjrlmpfg","latitude":"-87.691337","longitude":"-168.954410"}],"id":"brkvff18d3b2mjrlmpf0","spatialRef":"MU"},"id":"brkvff18d3b2mjrlmpg0","referencedAddress":{"id":"brkvff18d3b2mjrlmpd0","referenceId":"brkvff18d3b2mjrlmpdg","referenceType":"refer"},"relatedParty":[{"@referredType":"Organization","emailAddress":"missourimoen@aufderhar.net","id":"brkvff18d3b2mjrlmphg","name":"Ford Blanda","number":"5290865606","role":["Buyer","Seller","Site Contact"]}],"siteCompanyName":"DataLogix","siteCustomerName":"Sadie Douglas","siteName":"Dax Kiehn","siteType":"PUBLIC","status":"planned"},{"additionalSiteInformation":"Ameliorated","description":"Future","fieldedAddress":{"city":"Cote Des Neiges, QC","country":"Canada","geographicSubAddress":[{"buildingName":"estruxture MTL-2","id":"brkvff18d3b2mjrlmpi0","levelNumber":"BASEMENT 1","levelType":"1","privateStreetName":"university","privateStreetNumber":"01","subUnit":[{"subUnitIdentifier":"473656281","subUnitType":"UNIT"}]}],"id":"brkvff18d3b2mjrlmpig","locality":"Locality","postcode":"H4B 3A2","stateOrProvince":"Maryland","streetName":"7001 Rue Saint-Jacques","streetNr":"2513","streetSuffix":"shire","streetType":"Alley"},"formattedAddress":{"addrLine1":"2513  shire , 7001 Rue Saint-Jacques shire","city":"Cote Des Neiges, QC","country":"Canada","locality":"Locality","postcode":"H4B 3A2","stateOrProvince":"Maryland"},"geographicLocation":{"geographicPoint":[{"id":"brkvff18d3b2mjrlmpn0","latitude":"15.090253","longitude":"-155.247770"}],"id":"brkvff18d3b2mjrlmpmg","spatialRef":"GB"},"id":"brkvff18d3b2mjrlmpng","referencedAddress":{"id":"brkvff18d3b2mjrlmpkg","referenceId":"brkvff18d3b2mjrlmpl0","referenceType":"refer"},"relatedParty":[{"@referredType":"Organization","emailAddress":"kristianpowlowski@christiansen.info","id":"brkvff18d3b2mjrlmpp0","name":"Breanna Windler","number":"1435310490","role":["Buyer","Seller","Site Contact"]}],"siteCompanyName":"VisualDoD, LLC","siteCustomerName":"Florida Towne","siteName":"Herta Kirlin","siteType":"PUBLIC","status":"planned"},{"additionalSiteInformation":"Polarised","description":"National","fieldedAddress":{"city":"Hamilton, ON","country":"Canada","geographicSubAddress":[{"buildingName":"Commerce Place I","id":"brkvff18d3b2mjrlmppg","levelNumber":"BASEMENT 1","levelType":"1","privateStreetName":"university","privateStreetNumber":"01","subUnit":[{"subUnitIdentifier":"487793581","subUnitType":"UNIT"}]}],"id":"brkvff18d3b2mjrlmpq0","locality":"Locality","postcode":"L8N 1E7","stateOrProvince":"North Carolina","streetName":"1 King Street West","streetNr":"911","streetSuffix":"port","streetType":"Alley"},"formattedAddress":{"addrLine1":"911  port , 1 King Street West port","city":"Hamilton, ON","country":"Canada","locality":"Locality","postcode":"L8N 1E7","stateOrProvince":"North Carolina"},"geographicLocation":{"geographicPoint":[{"id":"brkvff18d3b2mjrlmpug","latitude":"25.355667","longitude":"163.173151"}],"id":"brkvff18d3b2mjrlmpu0","spatialRef":"MV"},"id":"brkvff18d3b2mjrlmpv0","referencedAddress":{"id":"brkvff18d3b2mjrlmps0","referenceId":"brkvff18d3b2mjrlmpsg","referenceType":"refer"},"relatedParty":[{"@referredType":"Organization","emailAddress":"alberthahane@pagac.net","id":"brkvff18d3b2mjrlmq0g","name":"Trystan Ebert","number":"6726152987","role":["Buyer","Seller","Site Contact"]}],"siteCompanyName":"TuvaLabs","siteCustomerName":"Elnora Walter","siteName":"Mariela Armstrong","siteType":"PUBLIC","status":"planned"},{"additionalSiteInformation":"support","description":"National","fieldedAddress":{"city":"LaSalle, QC","country":"Canada","geographicSubAddress":[{"buildingName":"ROOT MTL-R1","id":"brkvff18d3b2mjrlmq10","levelNumber":"BASEMENT 1","levelType":"1","privateStreetName":"university","privateStreetNumber":"01","subUnit":[{"subUnitIdentifier":"628422945","subUnitType":"UNIT"}]}],"id":"brkvff18d3b2mjrlmq1g","locality":"Locality","postcode":"H8N2J8","stateOrProvince":"Connecticut","streetName":"2711 Dollard Ave","streetNr":"4778","streetSuffix":"chester","streetType":"Alley"},"formattedAddress":{"addrLine1":"4778  chester , 2711 Dollard Ave chester","city":"LaSalle, QC","country":"Canada","locality":"Locality","postcode":"H8N2J8","stateOrProvince":"Connecticut"},"geographicLocation":{"geographicPoint":[{"id":"brkvff18d3b2mjrlmq60","latitude":"-4.374473","longitude":"-54.911382"}],"id":"brkvff18d3b2mjrlmq5g","spatialRef":"GI"},"id":"brkvff18d3b2mjrlmq6g","referencedAddress":{"id":"brkvff18d3b2mjrlmq3g","referenceId":"brkvff18d3b2mjrlmq40","referenceType":"refer"},"relatedParty":[{"@referredType":"Organization","emailAddress":"alessandroherman@connelly.name","id":"brkvff18d3b2mjrlmq80","name":"Marquis Fisher","number":"1386525469","role":["Buyer","Seller","Site Contact"]}],"siteCompanyName":"Progressive Insurance Group","siteCustomerName":"Savion Price","siteName":"Logan Hayes","siteType":"PUBLIC","status":"planned"},{"additionalSiteInformation":"middleware","description":"Senior","fieldedAddress":{"city":"Longueuil, QC","country":"Canada","geographicSubAddress":[{"buildingName":"COLO-D2","id":"brkvff18d3b2mjrlmq8g","levelNumber":"BASEMENT 1","levelType":"1","privateStreetName":"university","privateStreetNumber":"01","subUnit":[{"subUnitIdentifier":"31602041","subUnitType":"UNIT"}]}],"id":"brkvff18d3b2mjrlmq90","locality":"Locality","postcode":"J4G 1S8","stateOrProvince":"New York","streetName":"530 Rue Beriault","streetNr":"9352","streetSuffix":"berg","streetType":"Alley"},"formattedAddress":{"addrLine1":"9352  berg , 530 Rue Beriault berg","city":"Longueuil, QC","country":"Canada","locality":"Locality","postcode":"J4G 1S8","stateOrProvince":"New York"},"geographicLocation":{"geographicPoint":[{"id":"brkvff18d3b2mjrlmqdg","latitude":"24.471591","longitude":"81.834837"}],"id":"brkvff18d3b2mjrlmqd0","spatialRef":"PA"},"id":"brkvff18d3b2mjrlmqe0","referencedAddress":{"id":"brkvff18d3b2mjrlmqb0","referenceId":"brkvff18d3b2mjrlmqbg","referenceType":"refer"},"relatedParty":[{"@referredType":"Organization","emailAddress":"parisbreitenberg@yost.io","id":"brkvff18d3b2mjrlmqfg","name":"Lisandro Morar","number":"6017198041","role":["Buyer","Seller","Site Contact"]}],"siteCompanyName":"Rand McNally","siteCustomerName":"Blair Wiza","siteName":"Anabel Abernathy","siteType":"PUBLIC","status":"planned"},{"additionalSiteInformation":"archive","description":"Legacy","fieldedAddress":{"city":"Markham, ON","country":"Canada","geographicSubAddress":[{"buildingName":"Metro Optic / I.C.E","id":"brkvff18d3b2mjrlmqg0","levelNumber":"BASEMENT 1","levelType":"1","privateStreetName":"university","privateStreetNumber":"01","subUnit":[{"subUnitIdentifier":"2103967337","subUnitType":"UNIT"}]}],"id":"brkvff18d3b2mjrlmqgg","locality":"Locality","postcode":"L6G 1B9","stateOrProvince":"Pennsylvania","streetName":"105 Clegg Road","streetNr":"569","streetSuffix":"mouth","streetType":"Alley"},"formattedAddress":{"addrLine1":"569  mouth , 105 Clegg Road mouth","city":"Markham, ON","country":"Canada","locality":"Locality","postcode":"L6G 1B9","stateOrProvince":"Pennsylvania"},"geographicLocation":{"geographicPoint":[{"id":"brkvff18d3b2mjrlmql0","latitude":"-71.606181","longitude":"125.383329"}],"id":"brkvff18d3b2mjrlmqkg","spatialRef":"CA"},"id":"brkvff18d3b2mjrlmqlg","referencedAddress":{"id":"brkvff18d3b2mjrlmqig","referenceId":"brkvff18d3b2mjrlmqj0","referenceType":"refer"},"relatedParty":[{"@referredType":"Organization","emailAddress":"keanupredovic@schneider.info","id":"brkvff18d3b2mjrlmqn0","name":"Price Koepp","number":"9530742877","role":["Buyer","Seller","Site Contact"]}],"siteCompanyName":"Bridgewater","siteCustomerName":"Cecil Kilback","siteName":"Lenore Watsica","siteType":"PUBLIC","status":"planned"},{"additionalSiteInformation":"Balanced","description":"International","fieldedAddress":{"city":"Mississauga, ON","country":"Canada","geographicSubAddress":[{"buildingName":"City Centre Plaza","id":"brkvff18d3b2mjrlmqng","levelNumber":"BASEMENT 1","levelType":"1","privateStreetName":"university","privateStreetNumber":"01","subUnit":[{"subUnitIdentifier":"1051386848","subUnitType":"UNIT"}]}],"id":"brkvff18d3b2mjrlmqo0","locality":"Locality","postcode":"L5B 1M2","stateOrProvince":"Vermont","streetName":"1 City Centre Dr","streetNr":"318","streetSuffix":"chester","streetType":"Alley"},"formattedAddress":{"addrLine1":"318  chester , 1 City Centre Dr chester","city":"Mississauga, ON","country":"Canada","locality":"Locality","postcode":"L5B 1M2","stateOrProvince":"Vermont"},"geographicLocation":{"geographicPoint":[{"id":"brkvff18d3b2mjrlmqsg","latitude":"-73.879599","longitude":"96.931947"}],"id":"brkvff18d3b2mjrlmqs0","spatialRef":"MU"},"id":"brkvff18d3b2mjrlmqt0","referencedAddress":{"id":"brkvff18d3b2mjrlmqq0","referenceId":"brkvff18d3b2mjrlmqqg","referenceType":"refer"},"relatedParty":[{"@referredType":"Organization","emailAddress":"winfieldkuhlman@funk.biz","id":"brkvff18d3b2mjrlmqug","name":"Bonnie Von","number":"2311393076","role":["Buyer","Seller","Site Contact"]}],"siteCompanyName":"Zillow","siteCustomerName":"Angelica Romaguera","siteName":"Bert Doyle","siteType":"PUBLIC","status":"planned"},{"additionalSiteInformation":"modular","description":"Direct","fieldedAddress":{"city":"Montreal, QC","country":"Canada","geographicSubAddress":[{"buildingName":"Place Ville Marie","id":"brkvff18d3b2mjrlmqv0","levelNumber":"BASEMENT 1","levelType":"1","privateStreetName":"university","privateStreetNumber":"01","subUnit":[{"subUnitIdentifier":"1606840919","subUnitType":"UNIT"}]}],"id":"brkvff18d3b2mjrlmqvg","locality":"Locality","postcode":"H3B 2E7","stateOrProvince":"Mississippi","streetName":"1 Place Ville Marie","streetNr":"82448","streetSuffix":"view","streetType":"Alley"},"formattedAddress":{"addrLine1":"82448  view , 1 Place Ville Marie view","city":"Montreal, QC","country":"Canada","locality":"Locality","postcode":"H3B 2E7","stateOrProvince":"Mississippi"},"geographicLocation":{"geographicPoint":[{"id":"brkvff18d3b2mjrlmr40","latitude":"-65.128580","longitude":"26.813906"}],"id":"brkvff18d3b2mjrlmr3g","spatialRef":"TH"},"id":"brkvff18d3b2mjrlmr4g","referencedAddress":{"id":"brkvff18d3b2mjrlmr1g","referenceId":"brkvff18d3b2mjrlmr20","referenceType":"refer"},"relatedParty":[{"@referredType":"Organization","emailAddress":"newtonwilkinson@wolf.org","id":"brkvff18d3b2mjrlmr60","name":"Akeem Schamberger","number":"7255804225","role":["Buyer","Seller","Site Contact"]}],"siteCompanyName":"LOGIXDATA, LLC","siteCustomerName":"Verla Romaguera","siteName":"Carrie Romaguera","siteType":"PUBLIC","status":"planned"},{"additionalSiteInformation":"service-desk","description":"Investor","fieldedAddress":{"city":"Ottawa, ON","country":"Canada","geographicSubAddress":[{"buildingName":"World Exchange Plaza Tower ll","id":"brkvff18d3b2mjrlmr6g","levelNumber":"BASEMENT 1","levelType":"1","privateStreetName":"university","privateStreetNumber":"01","subUnit":[{"subUnitIdentifier":"1522789417","subUnitType":"UNIT"}]}],"id":"brkvff18d3b2mjrlmr70","locality":"Locality","postcode":"K1P 1J9","stateOrProvince":"South Dakota","streetName":"100 Queen Street","streetNr":"312","streetSuffix":"side","streetType":"Alley"},"formattedAddress":{"addrLine1":"312  side , 100 Queen Street side","city":"Ottawa, ON","country":"Canada","locality":"Locality","postcode":"K1P 1J9","stateOrProvince":"South Dakota"},"geographicLocation":{"geographicPoint":[{"id":"brkvff18d3b2mjrlmrbg","latitude":"2.513707","longitude":"-113.125871"}],"id":"brkvff18d3b2mjrlmrb0","spatialRef":"PT"},"id":"brkvff18d3b2mjrlmrc0","referencedAddress":{"id":"brkvff18d3b2mjrlmr90","referenceId":"brkvff18d3b2mjrlmr9g","referenceType":"refer"},"relatedParty":[{"@referredType":"Organization","emailAddress":"abigailjerde@price.org","id":"brkvff18d3b2mjrlmrdg","name":"Tyshawn Pfannerstill","number":"9462487028","role":["Buyer","Seller","Site Contact"]}],"siteCompanyName":"Innovest Systems","siteCustomerName":"Cortez Koelpin","siteName":"Armand Lockman","siteType":"PUBLIC","status":"planned"},{"additionalSiteInformation":"cohesive","description":"Central","fieldedAddress":{"city":"Point-Claire, QC","country":"Canada","geographicSubAddress":[{"buildingName":"Hypertec","id":"brkvff18d3b2mjrlmre0","levelNumber":"BASEMENT 1","levelType":"1","privateStreetName":"university","privateStreetNumber":"01","subUnit":[{"subUnitIdentifier":"1398071123","subUnitType":"UNIT"}]}],"id":"brkvff18d3b2mjrlmreg","locality":"Locality","postcode":"H9R 4A9","stateOrProvince":"Michigan","streetName":"2800 Trans-Canada Highway","streetNr":"76844","streetSuffix":"land","streetType":"Alley"},"formattedAddress":{"addrLine1":"76844  land , 2800 Trans-Canada Highway land","city":"Point-Claire, QC","country":"Canada","locality":"Locality","postcode":"H9R 4A9","stateOrProvince":"Michigan"},"geographicLocation":{"geographicPoint":[{"id":"brkvff18d3b2mjrlmrj0","latitude":"-17.407951","longitude":"23.734543"}],"id":"brkvff18d3b2mjrlmrig","spatialRef":"NP"},"id":"brkvff18d3b2mjrlmrjg","referencedAddress":{"id":"brkvff18d3b2mjrlmrgg","referenceId":"brkvff18d3b2mjrlmrh0","referenceType":"refer"},"relatedParty":[{"@referredType":"Organization","emailAddress":"aliagibson@robel.org","id":"brkvff18d3b2mjrlmrl0","name":"Noelia Green","number":"2228489896","role":["Buyer","Seller","Site Contact"]}],"siteCompanyName":"GovTribe","siteCustomerName":"Nathanial Sauer","siteName":"Ellis Mills","siteType":"PUBLIC","status":"planned"},{"additionalSiteInformation":"groupware","description":"National","fieldedAddress":{"city":"Saint-Laurent, QC","country":"Canada","geographicSubAddress":[{"buildingName":"Cologix MTL6","id":"brkvff18d3b2mjrlmrlg","levelNumber":"BASEMENT 1","levelType":"1","privateStreetName":"university","privateStreetNumber":"01","subUnit":[{"subUnitIdentifier":"970720160","subUnitType":"UNIT"}]}],"id":"brkvff18d3b2mjrlmrm0","locality":"Locality","postcode":"H4S2A9","stateOrProvince":"New Hampshire","streetName":"2341 Alfred Nobel","streetNr":"1284","streetSuffix":"mouth","streetType":"Alley"},"formattedAddress":{"addrLine1":"1284  mouth , 2341 Alfred Nobel mouth","city":"Saint-Laurent, QC","country":"Canada","locality":"Locality","postcode":"H4S2A9","stateOrProvince":"New Hampshire"},"geographicLocation":{"geographicPoint":[{"id":"brkvff18d3b2mjrlmrqg","latitude":"-39.903613","longitude":"150.452158"}],"id":"brkvff18d3b2mjrlmrq0","spatialRef":"DO"},"id":"brkvff18d3b2mjrlmrr0","referencedAddress":{"id":"brkvff18d3b2mjrlmro0","referenceId":"brkvff18d3b2mjrlmrog","referenceType":"refer"},"relatedParty":[{"@referredType":"Organization","emailAddress":"ramonaberge@gerlach.com","id":"brkvff18d3b2mjrlmrsg","name":"Lauryn Lockman","number":"7318746655","role":["Buyer","Seller","Site Contact"]}],"siteCompanyName":"Plus-U","siteCustomerName":"Delores Hudson","siteName":"Freeda Huels","siteType":"PUBLIC","status":"planned"},{"additionalSiteInformation":"Cross-group","description":"Investor","fieldedAddress":{"city":"Saint-Leonard, QC","country":"Canada","geographicSubAddress":[{"buildingName":"iWeb-CL","id":"brkvff18d3b2mjrlmrt0","levelNumber":"BASEMENT 1","levelType":"1","privateStreetName":"university","privateStreetNumber":"01","subUnit":[{"subUnitIdentifier":"1206004879","subUnitType":"UNIT"}]}],"id":"brkvff18d3b2mjrlmrtg","locality":"Locality","postcode":"H1P 1A8","stateOrProvince":"California","streetName":"5945 Couture Blvd","streetNr":"57239","streetSuffix":"burgh","streetType":"Alley"},"formattedAddress":{"addrLine1":"57239  burgh , 5945 Couture Blvd burgh","city":"Saint-Leonard, QC","country":"Canada","locality":"Locality","postcode":"H1P 1A8","stateOrProvince":"California"},"geographicLocation":{"geographicPoint":[{"id":"brkvff18d3b2mjrlms20","latitude":"86.780870","longitude":"-92.096297"}],"id":"brkvff18d3b2mjrlms1g","spatialRef":"MP"},"id":"brkvff18d3b2mjrlms2g","referencedAddress":{"id":"brkvff18d3b2mjrlmrvg","referenceId":"brkvff18d3b2mjrlms00","referenceType":"refer"},"relatedParty":[{"@referredType":"Organization","emailAddress":"brendenschumm@moen.info","id":"brkvff18d3b2mjrlms40","name":"Carrie Mohr","number":"4430192799","role":["Buyer","Seller","Site Contact"]}],"siteCompanyName":"xDayta","siteCustomerName":"Ryder Brekke","siteName":"Russel Weimann","siteType":"PUBLIC","status":"planned"},{"additionalSiteInformation":"Self-enabling","description":"Product","fieldedAddress":{"city":"Toronto, ON","country":"Canada","geographicSubAddress":[{"buildingName":"One Financial Place","id":"brkvff18d3b2mjrlms4g","levelNumber":"BASEMENT 1","levelType":"1","privateStreetName":"university","privateStreetNumber":"01","subUnit":[{"subUnitIdentifier":"2107794013","subUnitType":"UNIT"}]}],"id":"brkvff18d3b2mjrlms50","locality":"Locality","postcode":"M5C 2V9","stateOrProvince":"Utah","streetName":"1 Adelaide St. E.","streetNr":"6156","streetSuffix":"stad","streetType":"Alley"},"formattedAddress":{"addrLine1":"6156  stad , 1 Adelaide St. E. stad","city":"Toronto, ON","country":"Canada","locality":"Locality","postcode":"M5C 2V9","stateOrProvince":"Utah"},"geographicLocation":{"geographicPoint":[{"id":"brkvff18d3b2mjrlms9g","latitude":"-36.581213","longitude":"-26.216415"}],"id":"brkvff18d3b2mjrlms90","spatialRef":"KG"},"id":"brkvff18d3b2mjrlmsa0","referencedAddress":{"id":"brkvff18d3b2mjrlms70","referenceId":"brkvff18d3b2mjrlms7g","referenceType":"refer"},"relatedParty":[{"@referredType":"Organization","emailAddress":"darionzemlak@breitenberg.org","id":"brkvff18d3b2mjrlmsbg","name":"Miles Gulgowski","number":"9549976022","role":["Buyer","Seller","Site Contact"]}],"siteCompanyName":"First Fuel Software","siteCustomerName":"Florencio Dooley","siteName":"Judson Bartell","siteType":"PUBLIC","status":"planned"},{"additionalSiteInformation":"paradigm","description":"Dynamic","fieldedAddress":{"city":"Vancouver, BC","country":"Canada","geographicSubAddress":[{"buildingName":"MNP Tower","id":"brkvff18d3b2mjrlmsc0","levelNumber":"BASEMENT 1","levelType":"1","privateStreetName":"university","privateStreetNumber":"01","subUnit":[{"subUnitIdentifier":"64035511","subUnitType":"UNIT"}]}],"id":"brkvff18d3b2mjrlmscg","locality":"Locality","postcode":"V6E 2E9","stateOrProvince":"Texas","streetName":"1021 W Hastings","streetNr":"33749","streetSuffix":"furt","streetType":"Alley"},"formattedAddress":{"addrLine1":"33749  furt , 1021 W Hastings furt","city":"Vancouver, BC","country":"Canada","locality":"Locality","postcode":"V6E 2E9","stateOrProvince":"Texas"},"geographicLocation":{"geographicPoint":[{"id":"brkvff18d3b2mjrlmsh0","latitude":"33.091626","longitude":"38.928348"}],"id":"brkvff18d3b2mjrlmsgg","spatialRef":"SH"},"id":"brkvff18d3b2mjrlmshg","referencedAddress":{"id":"brkvff18d3b2mjrlmseg","referenceId":"brkvff18d3b2mjrlmsf0","referenceType":"refer"},"relatedParty":[{"@referredType":"Organization","emailAddress":"ernestinaherman@mcglynn.org","id":"brkvff18d3b2mjrlmsj0","name":"Aletha Doyle","number":"5032053628","role":["Buyer","Seller","Site Contact"]}],"siteCompanyName":"Outline","siteCustomerName":"Louie Effertz","siteName":"Annetta Kovacek","siteType":"PUBLIC","status":"planned"},{"additionalSiteInformation":"Team-oriented","description":"Global","fieldedAddress":{"city":"Verdun, QC","country":"Canada","geographicSubAddress":[{"buildingName":"iWeb 3","id":"brkvff18d3b2mjrlmsjg","levelNumber":"BASEMENT 1","levelType":"1","privateStreetName":"university","privateStreetNumber":"01","subUnit":[{"subUnitIdentifier":"2146307654","subUnitType":"UNIT"}]}],"id":"brkvff18d3b2mjrlmsk0","locality":"Locality","postcode":"H3E 1Z6","stateOrProvince":"Missouri","streetName":"20 Place du Commerce","streetNr":"9155","streetSuffix":"bury","streetType":"Alley"},"formattedAddress":{"addrLine1":"9155  bury , 20 Place du Commerce bury","city":"Verdun, QC","country":"Canada","locality":"Locality","postcode":"H3E 1Z6","stateOrProvince":"Missouri"},"geographicLocation":{"geographicPoint":[{"id":"brkvff18d3b2mjrlmsog","latitude":"-19.736674","longitude":"-165.697568"}],"id":"brkvff18d3b2mjrlmso0","spatialRef":"ST"},"id":"brkvff18d3b2mjrlmsp0","referencedAddress":{"id":"brkvff18d3b2mjrlmsm0","referenceId":"brkvff18d3b2mjrlmsmg","referenceType":"refer"},"relatedParty":[{"@referredType":"Organization","emailAddress":"daxkeebler@sawayn.io","id":"brkvff18d3b2mjrlmsqg","name":"Laisha Beier","number":"5082016992","role":["Buyer","Seller","Site Contact"]}],"siteCompanyName":"Copyright Clearance Center","siteCustomerName":"Stuart Blanda","siteName":"Isobel Barrows","siteType":"PUBLIC","status":"planned"},{"additionalSiteInformation":"product","description":"Investor","fieldedAddress":{"city":"Chai Wan","country":"China","geographicSubAddress":[{"buildingName":"Evoque HNK1","id":"brkvff18d3b2mjrlmsr0","levelNumber":"BASEMENT 1","levelType":"1","privateStreetName":"university","privateStreetNumber":"01","subUnit":[{"subUnitIdentifier":"663773398","subUnitType":"UNIT"}]}],"id":"brkvff18d3b2mjrlmsrg","locality":"Locality","postcode":"0","stateOrProvince":"Connecticut","streetName":"399 Chai Wan Road","streetNr":"679","streetSuffix":"land","streetType":"Alley"},"formattedAddress":{"addrLine1":"679  land , 399 Chai Wan Road land","city":"Chai Wan","country":"China","locality":"Locality","postcode":"0","stateOrProvince":"Connecticut"},"geographicLocation":{"geographicPoint":[{"id":"brkvff18d3b2mjrlmt00","latitude":"77.654750","longitude":"-24.648909"}],"id":"brkvff18d3b2mjrlmsvg","spatialRef":"LB"},"id":"brkvff18d3b2mjrlmt0g","referencedAddress":{"id":"brkvff18d3b2mjrlmstg","referenceId":"brkvff18d3b2mjrlmsu0","referenceType":"refer"},"relatedParty":[{"@referredType":"Organization","emailAddress":"sammytremblay@senger.biz","id":"brkvff18d3b2mjrlmt20","name":"Merle Bechtelar","number":"7483945141","role":["Buyer","Seller","Site Contact"]}],"siteCompanyName":"Dun \u0026 Bradstreet","siteCustomerName":"Tressa Franecki","siteName":"Thalia Carroll","siteType":"PUBLIC","status":"planned"},{"additionalSiteInformation":"Monitored","description":"Internal","fieldedAddress":{"city":"Hong Kong","country":"China","geographicSubAddress":[{"buildingName":"Equinix HK3","id":"brkvff18d3b2mjrlmt2g","levelNumber":"BASEMENT 1","levelType":"1","privateStreetName":"university","privateStreetNumber":"01","subUnit":[{"subUnitIdentifier":"1631995181","subUnitType":"UNIT"}]}],"id":"brkvff18d3b2mjrlmt30","locality":"Locality","postcode":"999077","stateOrProvince":"Michigan","streetName":"1 Wang Wo Tsai Street, Tsuen Wan, New Territories","streetNr":"62820","streetSuffix":"borough","streetType":"Alley"},"formattedAddress":{"addrLine1":"62820  borough , 1 Wang Wo Tsai Street, Tsuen Wan, New Territories borough","city":"Hong Kong","country":"China","locality":"Locality","postcode":"999077","stateOrProvince":"Michigan"},"geographicLocation":{"geographicPoint":[{"id":"brkvff18d3b2mjrlmt7g","latitude":"79.397417","longitude":"26.501246"}],"id":"brkvff18d3b2mjrlmt70","spatialRef":"UY"},"id":"brkvff18d3b2mjrlmt80","referencedAddress":{"id":"brkvff18d3b2mjrlmt50","referenceId":"brkvff18d3b2mjrlmt5g","referenceType":"refer"},"relatedParty":[{"@referredType":"Organization","emailAddress":"evertmann@deckow.io","id":"brkvff18d3b2mjrlmt9g","name":"Orion Grimes","number":"2091926527","role":["Buyer","Seller","Site Contact"]}],"siteCompanyName":"Vizzuality","siteCustomerName":"Elenor Schaefer","siteName":"Donnell Gutkowski","siteType":"PUBLIC","status":"planned"}]`
)

func init() {
	dir := config.LogDir()
	fn := path.Join(dir, "sonata.log")

	lw, _ := rotatelogs.New(
		fn+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(fn),
	)

	lh := lfshook.NewHook(
		lw,
		&log.JSONFormatter{},
	)
	log.AddHook(lh)
	//log.SetFormatter(&log.TextFormatter{
	//	FullTimestamp: true,
	//	DisableColors: true,
	//})
	log.SetReportCaller(true)
	log.SetLevel(log.ErrorLevel)
}

//go:generate swagger generate server --target ../../go-sonata-server --name Sonata --spec ../spec/all.yaml --principal models.Principal

func configureFlags(api *operations.SonataAPI) {
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		{
			ShortDescription: "jwt",
			LongDescription:  "jwt options",
			Options:          config.Cfg.Jwt,
		}, {
			ShortDescription: "debug",
			LongDescription:  "debug options",
			Options:          config.Cfg.Debug,
		}, {
			ShortDescription: "server",
			LongDescription:  "server options",
			Options:          config.Cfg.Server,
		},
	}
}

func configureAPI(api *operations.SonataAPI) http.Handler {
	cfg := config.Cfg

	if cfg.Debug.Verbose {
		log.SetLevel(log.DebugLevel)
	}

	log.Debug(util.ToIndentString(cfg))

	var (
		err  error
		file string
	)
	if config.Cfg.Debug.IsFile {
		dir := config.DBDir()
		file = path.Join(dir, "sonata.db")

		log.Debugf("save data to db @ %s", file)
	} else {
		file = "file:mockdb?mode=memory&cache=shared"
		log.Debugf("use db memory mode...")
	}
	//Store, err = gorm.Open(sqlite.Open("file:mockdb?mode=memory&cache=shared"), &gorm.Config{
	//	SkipDefaultTransaction: false,
	//	NamingStrategy:         nil,
	//	Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
	//		SlowThreshold: 100 * time.Millisecond,
	//		LogLevel:      logger.Info,
	//		Colorful:      true,
	//	}),
	//})
	db.Store, err = gorm.Open("sqlite3", file)
	if err != nil {
		log.Fatalln(err)
	}

	if cfg.Debug.Verbose {
		db.Store.LogMode(true)
	}

	if err := db.CreateTables(); err != nil {
		log.Fatalln(err)
	}

	if cfg.Debug.IsMock {
		if err := mockData(cfg); err != nil {
			log.Fatal(err)
		} else {
			log.Debug("insert mock data successful")
		}
	}

	jwtCfg := cfg.Jwt
	if jwtCfg.Key != "" {
		var err error
		if jwtCfg.PrivateKey, err = auth.FromBase64(jwtCfg.Key); err != nil {
			log.Fatal(err)
		} else {
			jwtCfg.PublicKey = &jwtCfg.PrivateKey.PublicKey
		}
	}

	if jwtCfg.PrivateKey == nil {
		log.Debug("use default key...")
		jwtCfg.PrivateKey = auth.Decode(auth.DefaultKey)
		jwtCfg.PublicKey = &jwtCfg.PrivateKey.PublicKey
	}

	// configure the api here
	api.ServeError = errors.ServeError
	api.Logger = log.Printf
	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	api.BearerAuth = func(token string, scopes []string) (*models.Principal, error) {
		// TODO: verify scopes???
		if claims, err := auth.ParseAndCheckToken(token, jwtCfg.PublicKey); err == nil {
			return &models.Principal{
				Code:   0,
				Reason: "",
				Roles:  claims.Roles,
			}, nil
		} else {
			switch vErr, ok := err.(*jwt.ValidationError); ok {
			case vErr.Errors&jwt.ValidationErrorMalformed != 0:
				fallthrough
			case vErr.Errors&jwt.ValidationErrorUnverifiable != 0:
				fallthrough
			case vErr.Errors&jwt.ValidationErrorSignatureInvalid != 0:
				fallthrough
			case vErr.Errors&jwt.ValidationErrorClaimsInvalid != 0:
				return &models.Principal{
					Code: 41, Reason: "Invalid credentials",
				}, nil
			case vErr.Errors&jwt.ValidationErrorExpired != 0:
				return &models.Principal{
					Code: 42, Reason: "Expired credentials",
				}, nil
			default:
				return &models.Principal{
					Code: 400, Reason: err.Error(),
				}, nil
			}
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	binder := &mock.MockBinder{}
	binder.Bind(api)

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {
		// clear all subscriber
		db.Store.Delete(&schema.HubSubscriber{})
		if db.Store != nil {
			if err := db.Store.Close(); err != nil {
				log.Error(err)
			}
		}
	}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	writer := &logger{}
	cfg := config.Cfg.Server
	return alice.New(handlers.CORS(handlers.AllowedOrigins(cfg.AllowedOrigins), handlers.AllowCredentials()),
		func(handler http.Handler) http.Handler {
			return handlers.LoggingHandler(writer, handler)
		},
		handlers.ProxyHeaders,
		handlers.CompressHandler,
		handlers.RecoveryHandler(
			handlers.RecoveryLogger(writer),
			handlers.PrintRecoveryStack(true),
		)).Then(handler)
}

type logger struct {
}

func (l *logger) Write(p []byte) (n int, err error) {
	log.Debugln(string(p))
	return len(p), nil
}

func (l *logger) Println(args ...interface{}) {
	log.Error(args...)
}

func mockData(cfg *config.SonataCfg) error {
	//mock site
	var site []*schema.GeographicSite
	var china *schema.GeographicSite
	if err := json.Unmarshal([]byte(sites), &site); err != nil {
		return err
	} else {
		for _, geographicSite := range site {
			id := geographicSite.ID
			geographicSite.FormattedAddress.ID = id
			geographicSite.GeographicLocation.ID = id
			geographicSite.FieldedAddress.ID = id

			if china == nil && geographicSite.FieldedAddress.Country == "China" {
				china = geographicSite
			}
			if err := db.Store.Create(geographicSite).Error; err != nil {
				log.Error(err)
			}
		}
	}

	//mock address
	for _, s := range site {
		a := &schema.GeographicAddress{
			AtSchemaLocation: "",
			AtType:           "",
			AllowsNewSite:    false,
			//FieldedAddress:     s.FieldedAddress,
			//FormattedAddress:   s.FormattedAddress,
			//GeographicLocation: s.GeographicLocation,
			HasPublicSite:     true,
			ID:                s.FormattedAddress.ID,
			ReferencedAddress: nil,
		}
		if err := db.Store.Create(a).Error; err != nil {
			log.Error(err)
		}
	}

	//mock sample product
	id := util.NewIDPtr()
	agreementId := util.NewIDPtr()
	status := models.ProductStatusActive

	from := &models.Product{
		//AtBaseType:       "",
		//AtSchemaLocation: "",
		AtType: "E-Line Spec",
		Agreement: []*models.Agreement{
			&models.Agreement{
				ID:   agreementId,
				Name: "E-Line Spec Agreement",
				Path: handler.HrefToID(fmt.Sprintf("%s/%s", cfg.Server.Domain, "agreements"), swag.StringValue(agreementId)),
			},
		},
		//BillingAccount: []*models.BillingAccountRef{
		//	{
		//		ID: util.NewIDPtr(),
		//	},
		//},
		//BuyerProductID: util.NewID(),
		Href:           handler.HrefToID(fmt.Sprintf("%s/%s", cfg.Server.Domain, "api/mef/productInventoryManagement/v3/product"), *id),
		ID:             id,
		LastUpdateDate: strfmt.DateTime(time.Now()),
		ProductOffering: &models.ProductOfferingRef{
			ID: util.NewIDPtr(),
		},
		//ProductOrder: []*models.ProductOrderRef{
		//	{
		//		Href:        gofakeit.URL(),
		//		ID:          util.NewIDPtr(),
		//		OrderItemID: swag.String("01"),
		//	},
		//},
		ProductPrice: []*models.ProductPrice{
			handler.ProductPrice(),
		},
		//ProductRelationship: []*models.ProductRelationship{
		//	{
		//		Product: &models.ProductRef{
		//			BuyerProductID: util.NewID(),
		//			Href:           gofakeit.URL(),
		//			ID:             util.NewIDPtr(),
		//		},
		//		Type: swag.String("bundled"),
		//	},
		//},
		ProductSpecification: &models.ProductSpecificationRef{
			Describing: handler.Describing(),
			ID:         util.NewIDPtr(),
		},
		ProductTerm: []*models.ProductTerm{
			{
				Description: "ProductTerm",
				Duration: &models.Quantity{
					Amount: float32(gofakeit.Uint8()),
					Units:  "day",
				},
				Name: "Term",
				ValidFor: &models.TimePeriod{
					EndDateTime:   strfmt.DateTime(time.Now().AddDate(2, 0, 0)),
					StartDateTime: strfmt.DateTime(time.Now().AddDate(-1, 0, 0)),
				},
			},
		},
		RelatedParty: []*models.RelatedParty{
			{
				AtReferredType:  "Organization",
				EmailAddress:    "hi@cbccom.cn",
				ID:              util.NewIDPtr(),
				Name:            swag.String(gofakeit.Name()),
				Number:          gofakeit.Phone(),
				NumberExtension: "",
				Role: []string{
					"Buyer", "Seller", "Site Contact",
				},
			},
		},
		Site: []*models.GeographicSite{
			china.To(),
		},
		StartDate: handler.NewDatetime(time.Now().AddDate(-1, 0, 0)),
		Status:    &status,
		//StatusChange: []*models.StatusChange{
		//	StatusChange(),
		//},
		//TerminationDate: strfmt.DateTime(time.Now().AddDate(10, 0, 0)),
	}
	to := schema.FromProduct(from)
	if err := db.Store.Create(to).Error; err != nil {
		log.Error(err)
	}

	return nil
}

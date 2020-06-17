/*
 * Copyright (c) 2020. QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package main

import (
	"encoding/csv"
	"fmt"
	"strings"

	"github.com/go-openapi/swag"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/restapi/handler"
	"github.com/qlcchain/go-sonata-server/util"
)

const (
	from = `2,Australia,"Alexandria, Sydney",200 Bourke Road,Equinix SY4,NSW 2015,CNDC,COGENT,151.189,-33.9182,2020/5/24 11:50
3,Australia,"Alexandria, Sydney",47 Bourke Road,Equinix SY3,NSW 2015,CNDC,COGENT,151.193,-33.9118,2020/5/24 11:50
4,Australia,Macquarie Park,17-23 Talavera Road,Macquarie DC - IC2,2113,CNDC,COGENT,151.127,-33.7808,2020/5/24 11:50
5,Australia,Macquarie Park,4 Eden Park Drive,NEXTDC S1,2113,CNDC,COGENT,151.131,-33.7853,2020/5/24 11:50
6,Australia,Malaga,4 Millrose Drive,COI Hub - PER01 - P1 NextDC,6090,CNDC,COGENT,115.896,-31.8644,2020/5/24 11:50
7,Australia,"Mascot, Sydney",639 Gardeners Road-Unit B,Equinix SY1,NSW 2020,CNDC,COGENT,151.188,-33.9214,2020/5/24 11:50
8,Australia,"Mascot, Sydney",639 Gardeners Road-Unit C,Equinix SY2,NSW 2020,CNDC,COGENT,151.188,-33.9214,2020/5/24 11:50
9,Australia,Silverwater,8-14 Egerton Street,Equinix SY6 fka Metronode,2128,CNDC,COGENT,151.053,-33.8367,2020/5/24 11:50
10,Australia,Sydney,400 Harris Street,TechFlow Services,2007,CNDC,COGENT,151.198,-33.8756,2020/5/24 11:50
11,Australia,Sydney,400 Harris Street,Global Switch East/West,2007,CNDC,COGENT,151.198,-33.8756,2020/5/24 11:50
12,Australia,Sydney,400 Harris Street,Evoque SYD1,2007,CNDC,COGENT,151.198,-33.8756,2020/5/24 11:50
13,Australia,Sydney,400 Harris Street,AAPT,2007,CNDC,COGENT,151.198,-33.8756,2020/5/24 11:50
14,Australia,Sydney,477 Pitt Street,Macquarie DC - IC1,2000,CNDC,COGENT,151.206,-33.8805,2020/5/24 11:50
15,Australia,Sydney,506-518 Gardeners Road,Equinix SY5,2015,CNDC,COGENT,151.19,-33.9194,2020/5/24 11:50
16,Austria,Vienna,Computerstrasse 2-4,e-Shelter,1000,CNDC,COGENT,16.3377,48.1593,2020/5/24 11:50
17,Austria,Vienna,Louis-HÃ¤fliger-Gasse 10 Bldg. 50,Interxion VIE1 (VIX2),1210,CNDC,COGENT,16.4106,48.269,2020/5/24 11:50
18,Austria,Vienna,Universitaetsstrasse 7,VIX1 (Aconet),1010,CNDC,COGENT,16.3578,48.2141,2020/5/24 11:50
19,Belgium,Antwerp,Antwerpsesteenweg 211-223 unit 1 (Noorderlaan 147,Yellow Duck,2070,CNDC,COGENT,4.41632,51.2308,2020/5/24 11:50
20,Belgium,Antwerp,Haifastraat 6,Antwerp DC,B 2030,CNDC,COGENT,4.41457,51.2631,2020/5/24 11:50
21,Belgium,Antwerpen,Noorderlaan 133,LCL Antwerpen,2030,CNDC,COGENT,4.40326,51.2641,2020/5/24 11:50
22,Belgium,Brussels,2 Avenue Leon GrosJean,CenturyLink (fka. Level 3),1140,CNDC,COGENT,4.41212,50.8577,2020/5/24 11:50
23,Belgium,Diegem,Kouterveldstraat 13,LCL Diegem,1831,CNDC,COGENT,4.45562,50.8874,2020/5/24 11:50
24,Belgium,Erembodegem (Aalst),Industrielaan 3,LCL Aalst,9320,CNDC,COGENT,4.86927,51.1539,2020/5/24 11:50
25,Belgium,Gent,Poortakkerstraat 33,Data Center Gent (DCG),B-9051,CNDC,COGENT,3.68334,51.0268,2020/5/24 11:50
27,Belgium,"Nossegem, Brussels",Mercurusstraat 30,COLT Brussels 1 (Nossegem 1 & 2),B-1930,CNDC,COGENT,4.50422,50.8719,2020/5/24 11:50
28,Belgium,Zaventem,Excelsiorlaan 15,Brussels DC (fka. AXS),1930,CNDC,COGENT,4.45412,50.8814,2020/5/24 11:50
29,Belgium,Zaventem,Wezembeekstraat 2,Interxion BRU1,1930,CNDC,COGENT,4.47694,50.8705,2020/5/24 11:50
30,Brazil,Rio de Janeiro,"Estrada Adhemar Bebiano, 1380 Del Castilho",Equinix RJ2,21051-070,CNDC,COGENT,-43.2731,-22.8714,2020/5/24 11:50
31,Brazil,Tambore-Barueri/SP,"Avenida Ceci, 1900",Equinix SP4,06460-120,CNDC,COGENT,-46.8084,-23.4926,2020/5/24 11:50
33,Bulgaria,Sofia,"10 5030 Str. ,Druzhba 1 district",Equinix SO1 (fka. 3DC),1592,CNDC,COGENT,23.3954,42.6685,2020/5/24 11:50
34,Bulgaria,Sofia,122 Ovche Pole Street,Telepoint,1303,CNDC,COGENT,23.3062,42.7027,2020/5/24 11:50
35,Bulgaria,Sofia,135 Tsarigradsko Shosse,Daticum,1784,CNDC,COGENT,23.3904,42.6541,2020/5/24 11:50
36,Bulgaria,Sofia,3 Grigorii Gorbatenko Str.,Neterra - Teleport 2,1784,CNDC,COGENT,23.3693,42.6563,2020/5/24 11:50
37,Bulgaria,Sofia,bul Asen Yordanov 8 1592 gk Druzhba 1,Telepoint DC2,1592,CNDC,COGENT,23.3835,42.6732,2020/5/24 11:50
38,Bulgaria,Sofia,str. Poruchik Nedelcho Bonchev #35,Equinix SO2,1528,CNDC,COGENT,23.4076,42.6718,2020/5/24 11:50
39,Bulgaria,Varna,128 8 mi Primorski Polk Blvd.,Varna DC,9000,CNDC,COGENT,27.9288,43.215,2020/5/24 11:50
40,Canada,"Anjou, QC",7171 Jean Talon East,Cologix MTL4,H1N 3M2,CNDC,COGENT,-73.5696,45.5989,2020/5/24 11:50
41,Canada,"Baie d Urfe, QC",19701 Avenue Clark Graham,ROOT MTL-R2,H9X3T1,CNDC,COGENT,-73.9139,45.4241,2020/5/24 11:50
42,Canada,"Boucherville, QC",1350 Nobel St,eStruxture MTL-3 (fka. Kolotek),J4B 5H3,CNDC,COGENT,-73.4266,45.567,2020/5/24 11:50
43,Canada,"Cote Des Neiges, QC",7001 Rue Saint-Jacques,estruxture MTL-2,H4B 3A2,CNDC,COGENT,-73.6317,45.4576,2020/5/24 11:50
44,Canada,"Hamilton, ON",1 King Street West,Commerce Place I,L8N 1E7,COB,COGENT,-79.3781,43.6489,2020/5/24 11:50
45,Canada,"Hamilton, ON",21 King Street West,Commerce Place II,L8N 1E7,COB,COGENT,-79.8701,43.2567,2020/5/24 11:50
46,Canada,"LaSalle, QC",2711 Dollard Ave,ROOT MTL-R1,H8N2J8,CNDC,COGENT,-73.6437,45.4389,2020/5/24 11:50
47,Canada,"LaSalle, QC",8501 Elmslie St,Compass Datacenters MTL-R1,H8N 2W6,CNDC,COGENT,-73.6408,45.4389,2020/5/24 11:50
48,Canada,"Longueuil, QC",530 Rue Beriault,COLO-D2,J4G 1S8,CNDC,COGENT,-73.4754,45.5655,2020/5/24 11:50
49,Canada,"Markham, ON",105 Clegg Road,Metro Optic / I.C.E,L6G 1B9,CNDC,COGENT,-79.3425,43.8507,2020/5/24 11:50
50,Canada,"Markham, ON","3500 Steeles Ave. E., Tower 6",Toronto Stock Exchange - 2nd Floor Demarc,L3R 2Z1,CNDC,COGENT,-79.3393,43.8176,2020/5/24 11:50
51,Canada,"Markham, ON",4175 14th Ave,Cyxtera YYZ2 (fka. CenturyLink TR3),L3R 5R5,CNDC,COGENT,-79.3163,43.8383,2020/5/24 11:50
52,Canada,"Mississauga, ON",1 City Centre Dr,City Centre Plaza,L5B 1M2,COB,COGENT,-79.6396,43.596,2020/5/24 11:50
53,Canada,"Mississauga, ON",2920 Matheson Blvd. E.,Cogent Data Center - Mississauga,L4W 5J4,CDC,COGENT,-79.5912,43.6661,2020/5/24 11:50
54,Canada,"Mississauga, ON",2920 Matheson Blvd. E.,2920 Matheson Blvd. E.,L4W 5J4,COB,COGENT,-79.5912,43.6661,2020/5/24 11:50
55,Canada,"Mississauga, ON",50 Burnhamthorpe Rd West,Sussex Centre East,L5B 3C2,COB,COGENT,-79.6363,43.5918,2020/5/24 11:50
56,Canada,"Mississauga, ON",6800 Millcreek,Cyxtera YYZ1 (fka. CenturyLink TR1),L5N 4J9,CNDC,COGENT,-79.7543,43.5951,2020/5/24 11:50
57,Canada,"Mississauga, ON",90 Burnhamthorpe Rd West,Sussex Centre West,L5B 3C3,COB,COGENT,-79.6368,43.591,2020/5/24 11:50
58,Canada,"Montreal, QC",1 Place Ville Marie,Place Ville Marie,H3B 2E7,COB,COGENT,-73.5693,45.5015,2020/5/24 11:50
59,Canada,"Montreal, QC",1000 de la Gauchetier St W,1000 de la Gauchetier,H3B 0A2,COB,COGENT,-73.5661,45.498,2020/5/24 11:50
60,Canada,"Montreal, QC",1000 Sherbrooke Street West,1000 Sherbrooke West,H3A 3G4,COB,COGENT,-73.5752,45.5028,2020/5/24 11:50
61,Canada,"Montreal, QC",1002 Sherbrooke Street West,Tour Scotia,H3A 3L6,COB,COGENT,-73.5749,45.5023,2020/5/24 11:50
62,Canada,"Montreal, QC",1010 de la Gauchetier,1010 de la Gauchetier,H3B 2N2,COB,COGENT,-73.5672,45.4975,2020/5/24 11:50
63,Canada,"Montreal, QC",1100 Rene Levesque West,1100 Rene Levesque,H3B 5H5,COB,COGENT,-73.5703,45.498,2020/5/24 11:50
64,Canada,"Montreal, QC",1155 Metcalfe,Edifice Sun Life,H3B 2V6,COB,COGENT,-73.5702,45.5003,2020/5/24 11:50
65,Canada,"Montreal, QC",1155 Rene Levesque West,La Tour CIBC,H3B 2J6,COB,COGENT,-73.5709,45.4986,2020/5/24 11:50
67,Canada,"Montreal, QC",1250 Blvd Rene Levesque West,Cologix MTL3,H3B 4W8,CNDC,COGENT,-73.5705,45.4972,2020/5/24 11:50
68,Canada,"Montreal, QC",1250 Blvd Rene Levesque West,1250 Blvd Rene Levesque West,H3B 4W8,COB,COGENT,-73.5705,45.4972,2020/5/24 11:50
69,Canada,"Montreal, QC",1501 McGill College,Tour McGill College,H3A 3M8,COB,COGENT,-73.5717,45.5022,2020/5/24 11:50
70,Canada,"Montreal, QC",1751 Richardson,The Nordelec,H3K 1G6,COB,COGENT,-73.5621,45.4841,2020/5/24 11:50
71,Canada,"Montreal, QC",1800 McGill College Ave,Place Montreal Trust,H3A 3J6,COB,COGENT,-73.5722,45.502,2020/5/24 11:50
72,Canada,"Montreal, QC",1801 McGill College,1801 McGill College,H3A 2N4,COB,COGENT,-73.5723,45.5029,2020/5/24 11:50
73,Canada,"Montreal, QC",1981 McGill College Avenue,Le 1981 McGill College,H3A 3H3,COB,COGENT,-73.5728,45.5034,2020/5/24 11:50
74,Canada,"Montreal, QC",2 Place Ville Marie,2 Place Ville Marie,H3B 2E7,COB,COGENT,-73.5685,45.5017,2020/5/24 11:50
75,Canada,"Montreal, QC",2000 McGill College,Tour L Industrielle Vie,H3A 3H3,COB,COGENT,-73.5736,45.5026,2020/5/24 11:50
76,Canada,"Montreal, QC",2001 McGill College Avenue,2001 McGill College Avenu,H3A 1G1,COB,COGENT,-73.574,45.5036,2020/5/24 11:50
77,Canada,"Montreal, QC",3 Place Ville Marie,3 Place Ville Marie,H3B 2E7,COB,COGENT,-73.5693,45.5021,2020/5/24 11:50
78,Canada,"Montreal, QC",3000 Blvd Rene-Levesque,Cologix MTL2,H3E 1T9,CNDC,COGENT,-73.5413,45.467,2020/5/24 11:50
79,Canada,"Montreal, QC",3400 de Maisonneuve Blvd West,1 Place Alexis Nihon,H3Z3B8,COB,COGENT,-73.5856,45.4895,2020/5/24 11:50
80,Canada,"Montreal, QC",3500 de Maisonneuve Blvd West,2 Place Alexis Nihon,H3Z3B8,COB,COGENT,-73.5863,45.4897,2020/5/24 11:50
81,Canada,"Montreal, QC",4 Place Ville Marie,4 Place Ville Marie,H3B 2E7,COB,COGENT,-73.5692,45.5008,2020/5/24 11:50
82,Canada,"Montreal, QC",5 Place Ville Marie,5 Place Ville Marie,H3B 2E7,COB,COGENT,-73.5695,45.5007,2020/5/24 11:50
83,Canada,"Montreal, QC",600 de Maisonneuve West,Place de la Cathudrale (Tour KPMG),H3B 4W8,COB,COGENT,-73.571,45.5041,2020/5/24 11:50
84,Canada,"Montreal, QC",625 Blvd Rene-Levesque West,Cologix MTL1,H3B 1R2,CNDC,COGENT,-73.5672,45.5026,2020/5/24 11:50
85,Canada,"Montreal, QC",630 Rene Levesque West,Tour Telus,H3B 1S6,COB,COGENT,-73.5665,45.502,2020/5/24 11:50
86,Canada,"Montreal, QC",630 Rene Levesque West,Verizon,H3B 1S6,COB,COGENT,-73.5665,45.502,2020/5/24 11:50
87,Canada,"Montreal, QC",800 Rene Levesque,800 Rene Levesque,H3B 1X9,COB,COGENT,-73.6534,45.4869,2020/5/24 11:50
88,Canada,"Montreal, QC",800 Rue du Square Victoria,eStruxture MTL-1 (fka. Netelligent),H4Z 1J2,CNDC,COGENT,-73.5621,45.5003,2020/5/24 11:50
89,Canada,"Montreal, QC",800 Rue du Square Victoria,Tour de la Bourse,H4Z 1J2,COB,COGENT,-73.5621,45.5003,2020/5/24 11:50
90,Canada,"Montreal, QC",875 St. Antoine W,I.C.E. Data Centers,H3C 1A6,CNDC,COGENT,-73.5645,45.4989,2020/5/24 11:50
91,Canada,"Ottawa, ON",100 Queen Street,World Exchange Plaza Tower ll,K1P 1J9,COB,COGENT,-75.6974,45.4217,2020/5/24 11:50
92,Canada,"Ottawa, ON",111 Albert Street,World Exchange Plaza,K1P 1A4,COB,COGENT,-75.6975,45.4216,2020/5/24 11:50
93,Canada,"Ottawa, ON",160 Elgin St.,160 Elgin St.,K2P 2N8,COB,COGENT,-75.6929,45.4197,2020/5/24 11:50
94,Canada,"Ottawa, ON",250 Albert St,250 Albert st,K1P 6M1,COB,COGENT,-80.5326,43.4752,2020/5/24 11:50
95,Canada,"Ottawa, ON",340 Albert St,Constitution Square Tower lll,K1R 7Y6,COB,COGENT,-75.7025,45.4186,2020/5/24 11:50
96,Canada,"Ottawa, ON",350 Albert St,Constitution Square Tower,K1R 1A4,COB,COGENT,-75.7031,45.4184,2020/5/24 11:50
97,Canada,"Ottawa, ON",360 Albert St,Constitution Square Tower ll,K1R 7X7,COB,COGENT,-75.7036,45.4182,2020/5/24 11:50
98,Canada,"Ottawa, ON",45 OConnor,World Exchange Plaza Tower I,K1P 1A4,COB,COGENT,-75.6982,45.4212,2020/5/24 11:50
99,Canada,"Ottawa, ON",50 OConnor St,Sun Life Financial Centre,K1P6L2,COB,COGENT,-75.6989,45.4208,2020/5/24 11:50
100,Canada,"Ottawa, ON",99 Bank St.,Sun Life Financial Centre,K1P 1H4,COB,COGENT,-75.7001,45.4205,2020/5/24 11:50
101,Canada,"Ottawa, ON",99 Metcalfe,99 Metcalfe,K1P 6L7,COB,COGENT,-75.6952,45.4209,2020/5/24 11:50
102,Canada,"Point-Claire, QC",2800 Trans-Canada Highway,Hypertec,H9R 4A9,CNDC,COGENT,-79.47,46.3491,2020/5/24 11:50
103,Canada,"Saint-Laurent, QC",2341 Alfred Nobel,Cologix MTL6,H4S2A9,CNDC,COGENT,-73.7601,45.4825,2020/5/24 11:50
104,Canada,"Saint-Laurent, QC",2351 Alfred Nobel,Cologix MTL5,H4S2A9,CNDC,COGENT,-73.7606,45.4842,2020/5/24 11:50
105,Canada,"Saint-Laurent, QC","2900, av Marie-Curie",4 Degrees,4HS 2C3,CNDC,COGENT,-73.7555,45.4865,2020/5/24 11:50
106,Canada,"Saint-Laurent, QC",9300 Route Transcanadienne,Bell (fka. Hypertec 1),H4S 1K5,CNDC,COGENT,-74.7891,46.346,2020/5/24 11:50
107,Canada,"Saint-Leonard, QC",5945 Couture Blvd,iWeb-CL,H1P 1A8,CNDC,COGENT,-73.6078,45.5993,2020/5/24 11:50
108,Canada,"Toronto, ON",1 Adelaide St. E.,One Financial Place,M5C 2V9,COB,COGENT,-79.3778,43.6502,2020/5/24 11:50
109,Canada,"Toronto, ON",1 Dundas West,Toronto Eaton Centre,M5B 2H1,COB,COGENT,-79.6165,43.5802,2020/5/24 11:50
110,Canada,"Toronto, ON",1 Queen St East,1 Queen St East,M5C 2C5,COB,COGENT,-79.3787,43.6522,2020/5/24 11:50
111,Canada,"Toronto, ON",1 University Ave.,One University Avenue,M5J 2P1,COB,COGENT,-79.383,43.6459,2020/5/24 11:50
112,Canada,"Toronto, ON",1 Yonge St.,Amanah Tech Data Center,M5E 1E5,CNDC,COGENT,-79.3742,43.6429,2020/5/24 11:50
113,Canada,"Toronto, ON",1 Yonge St.,One Yonge Street,M5E 1E5,COB,COGENT,-79.3742,43.6429,2020/5/24 11:50
114,Canada,"Toronto, ON",10 Bay St.,"Phase II, WaterPark Place",M5J 2R8,COB,COGENT,-79.3774,43.6412,2020/5/24 11:50
115,Canada,"Toronto, ON",10 York Mills Rd.,"Building 3, York Mills Centre",M2P 2G4,COB,COGENT,-79.4062,43.7447,2020/5/24 11:50
116,Canada,"Toronto, ON",100 Adelaide Street West,EY Tower,M5H 1S3,COB,COGENT,-79.3825,43.65,2020/5/24 11:50
117,Canada,"Toronto, ON",100 King St W,1 First Canadian Place,M5X 1E5,COB,COGENT,-79.3817,43.6486,2020/5/24 11:50
118,Canada,"Toronto, ON",100 Wellington Street West,TD Centre,M5K 1E7,COB,COGENT,-79.3826,43.6468,2020/5/24 11:50
119,Canada,"Toronto, ON",100 Wellington West,Bell / Q9 Datacenter Toronto 1,M5K 1E7,COB,COGENT,-75.6987,45.4229,2020/5/24 11:50
120,Canada,"Toronto, ON",100 Yonge St.,100 Yonge St.,M5C 2W1,COB,COGENT,-79.3785,43.6498,2020/5/24 11:50
121,Canada,"Toronto, ON",105 Adelaide St. W.,Lombard Place,M5H 1P9,COB,COGENT,-79.3829,43.6492,2020/5/24 11:50
122,Canada,"Toronto, ON",111 Richmond St.,Richmond Adelaide Centre,M5H-2G4,COB,COGENT,-79.3736,43.6527,2020/5/24 11:50
123,Canada,"Toronto, ON",120 Adelaide Street W,Richmond Adelaide Complex,M5H 3P5,COB,COGENT,-79.3832,43.6499,2020/5/24 11:50
124,Canada,"Toronto, ON",121 Bloor St. E.,121 Bloor St. E.,M4W 3M5,COB,COGENT,-79.3839,43.6705,2020/5/24 11:50
125,Canada,"Toronto, ON",121 King St. W.,Standard Life Centre,M5H 3T9,COB,COGENT,-79.383,43.6476,2020/5/24 11:50
126,Canada,"Toronto, ON",123 Edward St.,123 Edward St.,M5G 1Z8,COB,COGENT,-79.3868,43.6556,2020/5/24 11:50
127,Canada,"Toronto, ON",123 Front St. W.,Citibank Place,M5J 2M2,COB,COGENT,-79.3831,43.6449,2020/5/24 11:50
128,Canada,"Toronto, ON",130 Adelaide St W,Oxford Tower,M5H 3P5,COB,COGENT,-79.3838,43.6497,2020/5/24 11:50
129,Canada,"Toronto, ON",130 Bloor St. W.,Air Canada Building,M5S 1N5,COB,COGENT,-79.3927,43.6692,2020/5/24 11:50
130,Canada,"Toronto, ON",130 King St West,Toronto Stock Exchange Colo,M5X 1K6,CNDC,COGENT,-79.3833,43.6485,2020/5/24 11:50
131,Canada,"Toronto, ON",130 King St West,The Exchange Tower,M5X 1K6,COB,COGENT,-79.3833,43.6485,2020/5/24 11:50
132,Canada,"Toronto, ON",141 Adelaide St. W.,141 Adelaide St. W.,M5H 3L5,COB,COGENT,-79.3841,43.6487,2020/5/24 11:50
133,Canada,"Toronto, ON",145 King St. W.,145 King Street West,M5H 1J8,COB,COGENT,-79.3839,43.6475,2020/5/24 11:50
134,Canada,"Toronto, ON",15 Richmond East,Cambridge Suites Hotel,M5C 1N2,COB,COGENT,-79.3779,43.6517,2020/5/24 11:50
135,Canada,"Toronto, ON",150 King St West,Sun Life Financial Tower,M5H 1J9,COB,COGENT,-79.3847,43.6481,2020/5/24 11:50
136,Canada,"Toronto, ON",150 York St.,150 York St.,M5H 3S5,COB,COGENT,-79.3847,43.6496,2020/5/24 11:50
137,Canada,"Toronto, ON",151 Front St. W.,Cologix TOR1,M5J 2N1,CNDC,COGENT,-79.3841,43.6447,2020/5/24 11:50
138,Canada,"Toronto, ON",151 Front St. W.,Equinix TR1 (fka. S&D),M5J 2N1,CNDC,COGENT,-79.3841,43.6447,2020/5/24 11:50
139,Canada,"Toronto, ON",151 Front St. W.,151 Front Street West,M5J 2N1,COB,COGENT,-79.3841,43.6447,2020/5/24 11:50
140,Canada,"Toronto, ON",151 Yonge St,Yonge Richmond Centre,M5C 2W7,COB,COGENT,-79.3785,43.6516,2020/5/24 11:50
141,Canada,"Toronto, ON",155 Wellington St W,RBC Centre,M5V3L3,COB,COGENT,-79.3861,43.6456,2020/5/24 11:50
142,Canada,"Toronto, ON",160 Bloor St. E.,JWT Building,M4W 3P7,COB,COGENT,-79.383,43.6716,2020/5/24 11:50
143,Canada,"Toronto, ON",161 Bay St,"Canada Trust Tower, Brookfield Place",M5J 2T3,COB,COGENT,-79.3788,43.6466,2020/5/24 11:50
144,Canada,"Toronto, ON",175 Bloor St. E.,"Phase II, Crown Life Place",M4W 3R8,COB,COGENT,-79.3825,43.6707,2020/5/24 11:50
145,Canada,"Toronto, ON",175 Bloor St. E.,"Phase I, Crown Life Place",M4W 3R8,COB,COGENT,-79.3825,43.6707,2020/5/24 11:50
146,Canada,"Toronto, ON",18 King St. East,18 King St. East,M5C 1C4,COB,COGENT,-79.3771,43.6496,2020/5/24 11:50
147,Canada,"Toronto, ON",180 Dundas St. W,Dundas-Edward Centre,M5G 1Z8,COB,COGENT,-79.3864,43.6554,2020/5/24 11:50
148,Canada,"Toronto, ON",181 Bay Street,"Bay Wellington Tower, Brookfield Place",M5J 2T3,COB,COGENT,-79.3786,43.6473,2020/5/24 11:50
149,Canada,"Toronto, ON",181 University Ave.,DBRS Tower,M5H 3M7,COB,COGENT,-79.3852,43.6492,2020/5/24 11:50
150,Canada,"Toronto, ON",199 Bay St.,"Commerce Court West, Commerce Court",M5J 2J2,COB,COGENT,-79.3795,43.6481,2020/5/24 11:50
151,Canada,"Toronto, ON",2 Bloor St. E.,Hudson Bay Centre,M4W 1A8,COB,COGENT,-79.3861,43.6711,2020/5/24 11:50
152,Canada,"Toronto, ON",2 Bloor St. W.,Cumberland Terrace,M4W 3E2,COB,COGENT,-79.3872,43.6703,2020/5/24 11:50
153,Canada,"Toronto, ON",2 Queen St East,Maritime Life Tower,M5C 3G7,COB,COGENT,-79.3788,43.6528,2020/5/24 11:50
154,Canada,"Toronto, ON",2 St. Clair Ave West,2 St. Clair West,M4V 1L5,COB,COGENT,-79.3946,43.6883,2020/5/24 11:50
155,Canada,"Toronto, ON",20 Bay St.,"Phase I, WaterPark Place",M5H 2N8,COB,COGENT,-79.378,43.6416,2020/5/24 11:50
156,Canada,"Toronto, ON",20 Dundas St W,Atrium on Bay,M5G 2C2,COB,COGENT,-79.6671,43.5357,2020/5/24 11:50
157,Canada,"Toronto, ON",20 Eglinton Ave W,Yonge-Eglinton Centre,M4R 1K8,COB,COGENT,-79.7023,43.5623,2020/5/24 11:50
158,Canada,"Toronto, ON",20 Queen Street West,Toronto Eaton Centre,M5H 2H1,COB,COGENT,-79.3806,43.6523,2020/5/24 11:50
159,Canada,"Toronto, ON",20 Richmond East,20 Richmond East,M5C 2C5,COB,COGENT,-79.3783,43.6521,2020/5/24 11:50
160,Canada,"Toronto, ON",20 York Mills Rd.,"Building 2, York Mills Centre",M2P 2C2,COB,COGENT,-79.4055,43.7448,2020/5/24 11:50
161,Canada,"Toronto, ON",200 Bay St. (N. Tower),Royal Bank Plaza-North Tower,M5J 2J8,COB,COGENT,-79.3801,43.6465,2020/5/24 11:50
162,Canada,"Toronto, ON",200 Bay St. (S.Tower),Royal Bank Plaza-South Tower,M5J 2J8,COB,COGENT,-79.8765,43.2524,2020/5/24 11:50
163,Canada,"Toronto, ON",200 Front St. W.,Simcoe Place,M5V 3K2,COB,COGENT,-79.3862,43.6449,2020/5/24 11:50
164,Canada,"Toronto, ON",200 King St. W.,"Merrill Lynch Tower, Sun Life Centre",M5H 3T4,COB,COGENT,-79.3858,43.6478,2020/5/24 11:50
165,Canada,"Toronto, ON",200 Wellington St. W.,MetroCentre-Wellington Tower (AT&T Canada Building,M5V 3G2,COB,COGENT,-79.3878,43.6459,2020/5/24 11:50
166,Canada,"Toronto, ON",207 Queens Quay W.,Newcourt Centre,M5J 1A7,COB,COGENT,-79.3803,43.6389,2020/5/24 11:50
167,Canada,"Toronto, ON",21 Melinda St.,"Commerce Court East, Commerce Court",M5J 2J2,COB,COGENT,-79.3784,43.6483,2020/5/24 11:50
168,Canada,"Toronto, ON",22 Adelaide Street West,Bay Adelaide East Tower,M5H 1L6,COB,COGENT,-79.3795,43.6505,2020/5/24 11:50
169,Canada,"Toronto, ON",220 Bay St.,220 Bay St.,M5J 2W4,COB,COGENT,-79.3802,43.6477,2020/5/24 11:50
170,Canada,"Toronto, ON",220 Yonge Street,220 Yonge Street,M5B 2H1,COB,COGENT,-79.3807,43.6541,2020/5/24 11:50
171,Canada,"Toronto, ON",222 Bay St.,222 Bay St.,M5K 1B2,COB,COGENT,-79.3799,43.6475,2020/5/24 11:50
172,Canada,"Toronto, ON",2225 Sheppard Ave. E.,Atria III,M2J 5C2,COB,COGENT,-79.3279,43.774,2020/5/24 11:50
173,Canada,"Toronto, ON",2235 Sheppard Ave. E.,Atria II,M2J 5B5,COB,COGENT,-79.3287,43.7746,2020/5/24 11:50
174,Canada,"Toronto, ON",2255 Sheppard Ave. E.,Atria I,M2J 4Y1,COB,COGENT,-79.3276,43.7739,2020/5/24 11:50
175,Canada,"Toronto, ON",2300 Yonge St,Yonge Eglinton Centre,M4R 1K8,COB,COGENT,-79.3996,43.7073,2020/5/24 11:50
176,Canada,"Toronto, ON",234 Bay Street,Toronto Dominion Center,M5K 1B2,COB,COGENT,-79.3804,43.6485,2020/5/24 11:50
177,Canada,"Toronto, ON",243 Consumers Rd.,Parkway Place,M2J 4W8,COB,COGENT,-79.3317,43.7698,2020/5/24 11:50
178,Canada,"Toronto, ON",245 Consumers Rd.,Parkway Place,M2J 1R3,COB,COGENT,-79.3309,43.7696,2020/5/24 11:50
179,Canada,"Toronto, ON",245 Consumers Rd. 300,Cogent Data Center - Toronto,M2J 1R3,CDC,COGENT,-79.3309,43.7696,2020/5/24 11:50
180,Canada,"Toronto, ON",25 Adelaide St East/44 Victoria,Victoria Tower,M5C 3A1,COB,COGENT,-79.3774,43.6505,2020/5/24 11:50
181,Canada,"Toronto, ON",25 King St. W.,"Commerce Court North, Commerce Court",M5J 2J2,COB,COGENT,-79.379,43.6486,2020/5/24 11:50
182,Canada,"Toronto, ON",250 Yonge Street,250 Yonge Street,M5B 2L7,COB,COGENT,-79.381,43.6542,2020/5/24 11:50
183,Canada,"Toronto, ON",251 Consumers Rd.,Parkway Place,M2J 4R3,COB,COGENT,-79.3305,43.77,2020/5/24 11:50
184,Canada,"Toronto, ON",255 Consumers Rd.,Parkway Place,M2J 1R4,COB,COGENT,-79.3301,43.77,2020/5/24 11:50
185,Canada,"Toronto, ON",3 Park Home Avenue,Novotel North York Centre,M2N 6L3,COB,COGENT,-79.4144,43.7684,2020/5/24 11:50
186,Canada,"Toronto, ON",30 Wellington,Commere Court South,M5E 1S3,COB,COGENT,-79.3755,43.6487,2020/5/24 11:50
187,Canada,"Toronto, ON",30 Yonge Street,Hockey Hall of Fame,M5E 1X8,COB,COGENT,-79.3774,43.647,2020/5/24 11:50
188,Canada,"Toronto, ON",320 Bay St.,320 Bay Street,M5H 4A6,COB,COGENT,-79.381,43.6494,2020/5/24 11:50
189,Canada,"Toronto, ON",33 Bloor St. E.,Xerox Centre,M4W 3H1,COB,COGENT,-79.3857,43.6698,2020/5/24 11:50
190,Canada,"Toronto, ON",33 Yonge Street,33 Yonge Street,M5E 1G4,COB,COGENT,-79.3766,43.6476,2020/5/24 11:50
191,Canada,"Toronto, ON",333 Bay Street,Bay Adelaide Centre,M5H 1L6,COB,COGENT,-79.3805,43.6503,2020/5/24 11:50
192,Canada,"Toronto, ON",36 York Mills Road,York Mills Centre,M2P 2G9,COB,COGENT,-79.4047,43.7446,2020/5/24 11:50
193,Canada,"Toronto, ON",390 Bay Street,Munich RE Centre,M5H 2Y2,COB,COGENT,-79.3816,43.6511,2020/5/24 11:50
194,Canada,"Toronto, ON",4 King St West,4 King St West,M5H 1B6,COB,COGENT,-79.3786,43.6493,2020/5/24 11:50
195,Canada,"Toronto, ON",40 Dundas St W,Atrium on Bay,M5G 2C2,COB,COGENT,-79.6166,43.5788,2020/5/24 11:50
196,Canada,"Toronto, ON",40 King St. W,Scotia Plaza,M5H 3Y2,COB,COGENT,-79.3795,43.6495,2020/5/24 11:50
197,Canada,"Toronto, ON",401 Bay St.,The Simpson Tower,M5H 2Y4,COB,COGENT,-78.7572,43.8795,2020/5/24 11:50
198,Canada,"Toronto, ON",4100 Yonge St,Yonge Corporate Center,M2P 2B7,COB,COGENT,-79.4081,43.746,2020/5/24 11:50
199,Canada,"Toronto, ON",4101 Yonge St.,"Building 4, York Mills Centre",M2P 2C9,COB,COGENT,-79.4066,43.7455,2020/5/24 11:50
200,Canada,"Toronto, ON",4110 Yonge,4110 Yonge,M2P 2B7,COB,COGENT,-79.4091,43.7463,2020/5/24 11:50
201,Canada,"Toronto, ON",4120 Yonge,4120 Yonge,M2P 2B7,COB,COGENT,-79.4093,43.7469,2020/5/24 11:50
202,Canada,"Toronto, ON",439 University Avenue,439 University,M5G 1Y8,COB,COGENT,-79.3875,43.6545,2020/5/24 11:50
203,Canada,"Toronto, ON",45 Parliament Street,Equinix TR2,M5A 0G7,CNDC,COGENT,-79.3618,43.6509,2020/5/24 11:50
204,Canada,"Toronto, ON",4711 Yonge St.,Procter & Gamble,M2N 6K8,COB,COGENT,-79.4099,43.7595,2020/5/24 11:50
205,Canada,"Toronto, ON",4950 Yonge St,Madison Centre,M2N 6K1,COB,COGENT,-79.4128,43.7648,2020/5/24 11:50
206,Canada,"Toronto, ON",5000 Yonge St,5000 Yonge St,M2N 7E9,COB,COGENT,-79.4129,43.7662,2020/5/24 11:50
207,Canada,"Toronto, ON",5001 Yonge St.,Royal Bank Building,M2N 6P6,COB,COGENT,-79.4116,43.766,2020/5/24 11:50
208,Canada,"Toronto, ON",5140 Yonge Street,North York City Centre,M2N 6L7,COB,COGENT,-79.4131,43.7684,2020/5/24 11:50
209,Canada,"Toronto, ON",5160 Yonge St.,"5160 Yonge Street, North York City Centre",M2N 6L9,COB,COGENT,-79.4134,43.7689,2020/5/24 11:50
210,Canada,"Toronto, ON",5255 Yonge St,Yonge Norton Center,M2N 6P4,COB,COGENT,-79.4131,43.7721,2020/5/24 11:50
211,Canada,"Toronto, ON",55 University Ave.,55 University Ave.,M5J 2H7,COB,COGENT,-79.3839,43.6467,2020/5/24 11:50
212,Canada,"Toronto, ON",55 York Street,55 York Street,M5J 1R7,COB,COGENT,-79.3825,43.6462,2020/5/24 11:50
214,Canada,"Toronto, ON",595 Bay St,Atrium on Bay,M5G 2C2,COB,COGENT,-79.3833,43.6562,2020/5/24 11:50
215,Canada,"Toronto, ON",60 Bloor Street West,60 Bloor Street West,M4W 3B5,COB,COGENT,-79.3892,43.6701,2020/5/24 11:50
216,Canada,"Toronto, ON",66 Wellington,Toronto Dominion Center,M5X 1E5,COB,COGENT,-79.381,43.6477,2020/5/24 11:50
217,Canada,"Toronto, ON",67 Yonge St.,67 Yonge St.,M5E 1J8,COB,COGENT,-79.3775,43.6487,2020/5/24 11:50
218,Canada,"Toronto, ON",70 York St.,Hongkong Bank of Canada Building,M5J 1S9,COB,COGENT,-79.3834,43.6468,2020/5/24 11:50
219,Canada,"Toronto, ON",77 Bloor St. West,77 Bloor St. West,M5S 1S2,COB,COGENT,-79.3896,43.6692,2020/5/24 11:50
220,Canada,"Toronto, ON",77 King Street West,Toronto Dominion Center North Tower,M5K 1H6,COB,COGENT,-79.382,43.6479,2020/5/24 11:50
221,Canada,"Toronto, ON",777 Bay St.,"Maclean Hunter Building, College Park",M5G 2C8,COB,COGENT,-79.3845,43.6605,2020/5/24 11:50
222,Canada,"Toronto, ON",79 Wellington,79 Wellington,M5J 1H1,COB,COGENT,-79.3813,43.6466,2020/5/24 11:50
223,Canada,"Toronto, ON",80 Richmond Street,Victory Building,M5H 2Y2,COB,COGENT,-79.375,43.6528,2020/5/24 11:50
224,Canada,"Toronto, ON",905 King St. W.,Cologix TOR2,M6K 3G9,CNDC,COGENT,-79.4123,43.6417,2020/5/24 11:50
225,Canada,"Toronto, ON",95 Wellington St.,95 Wellington St. West,M5J 2N7,COB,COGENT,-81.1943,42.7755,2020/5/24 11:50
226,Canada,"Vancouver, BC",1021 W Hastings,MNP Tower,V6E 2E9,COB,COGENT,-123.118,49.2876,2020/5/24 11:50
227,Canada,"Vancouver, BC",1050 West Pender Street,Cologix VAN2,V6E 3S7,CNDC,COGENT,-123.119,49.287,2020/5/24 11:50
228,Canada,"Vancouver, BC",1055 W Hastings,The Guinness Tower,V6E 2E9,COB,COGENT,-123.118,49.2877,2020/5/24 11:50
229,Canada,"Vancouver, BC",1055 West Dunsmuir St,Bentall IV,V6E 4E2,COB,COGENT,-123.121,49.2865,2020/5/24 11:50
230,Canada,"Vancouver, BC",1066 W. Hastings St,Oceanic Plaza,V6E 3X2,COB,COGENT,-123.119,49.2876,2020/5/24 11:50
231,Canada,"Vancouver, BC",1111 W Georgia Street,Terasen Centre,V6E 4G2,COB,COGENT,-123.123,49.2861,2020/5/24 11:50
232,Canada,"Vancouver, BC",200 Granville St,Granville Square,V6C 1S4,COB,COGENT,-123.113,49.2868,2020/5/24 11:50
233,Canada,"Vancouver, BC",355 Burrard St,The Marine Building,V6C 0B2,COB,COGENT,-123.117,49.2875,2020/5/24 11:50
234,Canada,"Vancouver, BC",505 Burrard St,Bentall l,V7X1M5,COB,COGENT,-123.118,49.2866,2020/5/24 11:50
235,Canada,"Vancouver, BC",555 Burrard St,Bentall ll,V7X 1M8,COB,COGENT,-123.119,49.2864,2020/5/24 11:50
236,Canada,"Vancouver, BC",555 Hastings Street West,Cyxtera YVR1,V6B 1M1,CNDC,COGENT,-123.112,49.2847,2020/5/24 11:50
237,Canada,"Vancouver, BC",555 Hastings Street West,Cologix VAN1,V6B 4N4,CNDC,COGENT,-123.112,49.2847,2020/5/24 11:50
238,Canada,"Vancouver, BC",555 Hastings Street West,Spencer Building - OffNet Customer Node,V6B 1M1,COB,COGENT,-123.112,49.2847,2020/5/24 11:50
239,Canada,"Vancouver, BC",555 Hastings Street West 6th Floor,6B MMR at Harbour Centre Tower,V6B 1M1,CNDC,COGENT,-123.112,49.2847,2020/5/24 11:50
240,Canada,"Vancouver, BC",595 Burrard St,Bentall lll,V7X 1G4,COB,COGENT,-123.119,49.2861,2020/5/24 11:50
241,Canada,"Vancouver, BC",609 Granville St,609 Granville,V7Y 1H2,COB,COGENT,-123.117,49.2836,2020/5/24 11:50
242,Canada,"Vancouver, BC",650 W. Georgia St.,Vancouver Center,V6B 4N7,COB,COGENT,-123.118,49.2819,2020/5/24 11:50
243,Canada,"Vancouver, BC",666 Burrard Street,Park Place,V6C 2Z7,COB,COGENT,-123.119,49.285,2020/5/24 11:50
244,Canada,"Vancouver, BC",700 W Pender,Pender Place 700,V6C 1G8,COB,COGENT,-123.115,49.2848,2020/5/24 11:50
245,Canada,"Vancouver, BC",700 West Georgia Street,TD Tower,V7Y 1H4,COB,COGENT,-123.119,49.2827,2020/5/24 11:50
246,Canada,"Vancouver, BC",701 W Georgia Street,IBM Tower,V7Y 1H4,COB,COGENT,-123.118,49.2834,2020/5/24 11:50
247,Canada,"Vancouver, BC",750 West Pender,Pender Place 750,V6C 1G8,COB,COGENT,-123.116,49.2849,2020/5/24 11:50
248,Canada,"Vancouver, BC",777 Dunsmuir,777 Dunsmuir Street Tower,V6C 2T8,COB,COGENT,-123.116,49.2844,2020/5/24 11:50
249,Canada,"Vancouver, BC",789 W Pender Street,789 W Pender,V6C 1H2,COB,COGENT,-123.115,49.2853,2020/5/24 11:50
250,Canada,"Vancouver, BC",885 W Georgia Street,HSBC Building,V6C 3E8,COB,COGENT,-123.12,49.2838,2020/5/24 11:50
251,Canada,"Verdun, QC",20 Place du Commerce,iWeb 3,H3E 1Z6,CNDC,COGENT,-73.5387,45.4681,2020/5/24 11:50
252,China,Chai Wan,399 Chai Wan Road,Evoque HNK1,0,CNDC,COGENT,114.247,22.2661,2020/5/24 11:50
253,China,Chai Wan,399 Chai Wan Road,MEGA-iAdvantage,0,CNDC,COGENT,114.247,22.2661,2020/5/24 11:50
255,China,Hong Kong,"1 Wang Wo Tsai Street, Tsuen Wan, New Territories",Equinix HK3,999077,CNDC,COGENT,114.119,22.3678,2020/5/24 11:50
257,China,Hong Kong,"168 Yeung UK Road, Tsuen Wan NT",Equinix HK1,999077,CNDC,COGENT,114.119,22.3656,2020/5/24 11:50
258,China,Hong Kong,17/F Kerry Warehouse 3 Shing Yiu Street Kwai Chung,Equinix HK2,0,CNDC,COGENT,114.119,22.3623,2020/5/24 11:50
259,China,Hong Kong,"22 Chun Cheong Street, Tseung Kwan O Industrial Es",Town Gas HKDC2,0,CNDC,COGENT,114.275,22.2869,2020/5/24 11:50
261,China,Hong Kong,"299 Wan Po Road, Tower 2, Tseung Kwan O",Equinix HK5,0,CNDC,COGENT,114.275,22.293,2020/5/24 11:50
262,China,Hong Kong,"8/F, Sino Favor Centre, No.1 On Yip Street, Chai W",HKCOLO,0,CNDC,COGENT,114.247,22.2666,2020/5/24 11:50
263,Colombia,Zona Franca de Bogota,Carrera 106 15A- 25 Manzana 6 Lote 27,Equinix BG1,110921,CNDC,COGENT,-74.1619,4.67135,2020/5/24 11:50
264,Croatia,Zagreb,122 Selska cesta,Hrvatski Telekom (Croatian Telecom),10000,CNDC,COGENT,15.9431,45.7971,2020/5/24 11:50
265,Croatia,Zagreb,Josipa Marohnica 5,CIX (Univ Computing Ctr),10000,CNDC,COGENT,15.9694,45.7924,2020/5/24 11:50
266,Croatia,Zagreb,Selska cesta 93,Altus IT,HR-10000,CNDC,COGENT,15.9439,45.8033,2020/5/24 11:50
267,Croatia,Zagreb,Vrtni put 1,VIPNET,10002,CNDC,COGENT,16.0315,45.7855,2020/5/24 11:50
268,Croatia,Zagreb Buzin,"Banl, 75",Megatrend,10100,CNDC,COGENT,15.9846,45.8119,2020/5/24 11:50
269,Czech Republic,Brno,Cejl 20,MasterDC Brno,602 00,CNDC,COGENT,16.6177,49.1972,2020/5/24 11:50
270,Czech Republic,Brno,Cejl 20,Dial Telecom,602 00,CNDC,COGENT,16.6177,49.1972,2020/5/24 11:50
272,Czech Republic,Prague,Nad Elektrarnou 1428/47,CE-COLO (fka. Sitel),10600,CNDC,COGENT,14.483,50.0602,2020/5/24 11:50
273,Czech Republic,Prague,Sazecska 10,TTC-Teleport DC2,108 00,CNDC,COGENT,14.5189,50.0818,2020/5/24 11:50
274,Czech Republic,Prague,Tiskarska 10,TTC Teleport DC1,11000,CNDC,COGENT,14.5348,50.0808,2020/5/24 11:50
275,Czech Republic,Prague,Tiskarska 10,DataCamp CZ,11000,CNDC,COGENT,14.5348,50.0808,2020/5/24 11:50
276,Denmark,Ballerup,Industriparken 20,Interxion CPH1,2750,CNDC,COGENT,12.3757,55.728,2020/5/24 11:50
277,Denmark,Glostrup,Sydvestvej 100,Telia,2600,CNDC,COGENT,12.3773,55.6619,2020/5/24 11:50
278,Denmark,Kolding,Kokbjerg 8,Global Connect,6000,CNDC,COGENT,9.47491,55.5345,2020/5/24 11:50
279,Denmark,Kongens Lyngby,"Matematiktorvet, Bldg 304",DTU / DIX,2800,CNDC,COGENT,12.5195,55.7844,2020/5/24 11:50
281,Estonia,Tallinn,Adala 29,Elisa Datacenter,10614,CNDC,COGENT,24.7217,59.4312,2020/5/24 11:50
282,Estonia,Tallinn,Laevastiku 3R,Infonet DC,10313,CNDC,COGENT,24.6932,59.4572,2020/5/24 11:50
283,Finland,Espoo,Keilaranta 14 (Keilaniemi),RUNNet,2150,CNDC,COGENT,24.8327,60.1781,2020/5/24 11:50
284,Finland,Espoo,Sinimaentie 12,Equinix HE6,2630,CNDC,COGENT,24.7785,60.2032,2020/5/24 11:50
285,Finland,Espoo,Tekniikantie 14,IP-Only (Innopoli 2),2150,CNDC,COGENT,24.812,60.1857,2020/5/24 11:50
286,Finland,Helsinki,Hiomotie 32,Equinix HE1,380,CNDC,COGENT,24.8702,60.2201,2020/5/24 11:50
287,Finland,Helsinki,"Kanavaranta 5, FI-00160",Equinix HE2 - Uspenski,160,CNDC,COGENT,24.9594,60.169,2020/5/24 11:50
289,Finland,Helsinki,"Parrukatu 2, FI-00570 Helsinki",Equinix HE3 - Suvilahti,570,CNDC,COGENT,24.9698,60.1852,2020/5/24 11:50
290,Finland,Helsinki,Radiokatu 5,Digita Oy,240,CNDC,COGENT,24.9238,60.2028,2020/5/24 11:50
292,France,Aubervilliers,10 rue Waldeck Rochet- Bldg. 520,Equinix PA6 (fka. Telecity Condorcet),93300,CNDC,COGENT,2.36825,48.9103,2020/5/24 11:50
293,France,Aubervilliers,20 rue des Gardinoux,Interxion PAR2,93300,CNDC,COGENT,2.37042,48.9066,2020/5/24 11:50
294,France,Aubervilliers,20 rue des Gardinoux,Evoque CDG1,93300,CNDC,COGENT,2.37042,48.9066,2020/5/24 11:50
295,France,Aubervilliers,45 Avenue Victor Hugo,Equinix PA5 (fka. Telecity Victor Hugo),93300,CNDC,COGENT,4.00596,49.2354,2020/5/24 11:50
296,France,Aubervilliers,45 Avenue Victor Hugo,Interxion PAR1,93300,CNDC,COGENT,4.00596,49.2354,2020/5/24 11:50
297,France,Bordeaux,BÃ¢t. G2 - Bassin Ã  flots,Cogent Data Center - Bordeaux,33300,CDC,COGENT,-0.555449,44.8633,2020/5/24 11:50
300,France,Caen,Rue de la Cotonniere,SFR Netcenter,14000,CNDC,COGENT,-0.402217,49.1975,2020/5/24 11:50
301,France,Cesson SevignÃ©,Route de Bray,SFR Netcenter,35 510,CNDC,COGENT,2.67706,49.9805,2020/5/24 11:50
302,France,Clichy,7/9 Rue Petit,Global Switch - Paris,92110,CNDC,COGENT,2.29609,48.8998,2020/5/24 11:50
303,France,Courbevoie,124 boulevard de Verdun,Telia c/o SFR,92400,CNDC,COGENT,2.25955,48.9042,2020/5/24 11:50
304,France,Courbevoie,"130-136 bd de Verdun, Energy Park Bat 8",Equinix PA7 (fka. Telecity Energy Park),92400,CNDC,COGENT,2.25933,48.9049,2020/5/24 11:50
305,France,Dijon,41 Quai Gauthey,Cogent Data Center - Dijon,21000,CDC,COGENT,5.03019,47.3089,2020/5/24 11:50
306,France,Grenoble,33 rue Joseph Chanrion,Cogent Data Center - Grenoble,38000,CDC,COGENT,5.73665,45.1899,2020/5/24 11:50
307,France,Ivry-sur-Seine Paris,11-15 rue Galilee,Interxion PAR6,94200,CNDC,COGENT,-2.79127,48.4859,2020/5/24 11:50
308,France,La Courneuve,1 Rue Rateau,Interxion PAR7,93120,CNDC,COGENT,2.401,48.9253,2020/5/24 11:50
309,France,La Garenne Colombes,77 boulevard de la RÃ©publique,Cogent Data Center - Paris 1,92250,CDC,COGENT,2.24156,48.9071,2020/5/24 11:50
310,France,La Garenne Colombes,77 boulevard de la RÃ©publique,Utility Computing,92250,COB,COGENT,2.24156,48.9071,2020/5/24 11:50
311,France,Les Milles,6505 Route de la Tour dArbois,TDF ProxiCenter,13290,CNDC,COGENT,5.32546,43.4618,2020/5/24 11:50
312,France,Lille,72 rue Jenner,Cogent Data Center - Lille,59000,CDC,COGENT,3.09055,50.6377,2020/5/24 11:50
313,France,Lyon,81 Boulevard du Parc d Artillerie,HOSTELYON,69007,CNDC,COGENT,4.84361,45.7292,2020/5/24 11:50
314,France,Magny les hameaux,1 rue Pablo Picasso,Telehouse 3 (Magny-les-Hameaux),78114,CNDC,COGENT,4.04675,44.1995,2020/5/24 11:50
315,France,Marseille,40 Avenue Roger Salengro,Interxion MRS1,13003,CNDC,COGENT,5.37334,43.3104,2020/5/24 11:50
316,France,Marseille,71 av andre roussin,Jaguar,13016,CNDC,COGENT,5.33824,43.3605,2020/5/24 11:50
319,France,Nanterre,3 Boulevard des Bouvets,Defense Data Center,92 000,CNDC,COGENT,2.22152,48.8979,2020/5/24 11:50
320,France,Nantes,32 boulevard Victor Hugo,Cogent Data Center - Nantes,44000,CDC,COGENT,2.33116,48.9106,2020/5/24 11:50
321,France,Nice,16 Avenue Thiers,Cogent Data Center - Nice,6000,CDC,COGENT,-0.558596,44.8413,2020/5/24 11:50
322,France,Nozay,3 Route de Marcoussis,Data4 - Paris DC07,91620,CNDC,COGENT,2.23646,48.6478,2020/5/24 11:50
323,France,Nozay,3 Route de Marcoussis,Data4 - Paris DC03,91620,CNDC,COGENT,2.23646,48.6478,2020/5/24 11:50
324,France,Nozay,3 Route de Marcoussis,Data4 - Paris DC08,91620,CNDC,COGENT,2.23646,48.6478,2020/5/24 11:50
325,France,Nozay,3 Route de Marcoussis,Data4 - Paris DC02,91620,CNDC,COGENT,2.23646,48.6478,2020/5/24 11:50
326,France,Nozay,3 Route de Marcoussis,Data4 - Paris DC04,91620,CNDC,COGENT,2.23646,48.6478,2020/5/24 11:50
327,France,Nozay,3 Route de Marcoussis,Data4 - Paris DC3,91620,CNDC,COGENT,2.23646,48.6478,2020/5/24 11:50
328,France,Nozay,3 Route de Marcoussis,Data4 - Paris DC05,91620,CNDC,COGENT,2.23646,48.6478,2020/5/24 11:50
329,France,Nozay,3 Route de Marcoussis,Data4 - Paris DC1,91620,CNDC,COGENT,2.23646,48.6478,2020/5/24 11:50
330,France,Pantin,110 bis avenue du General Leclerc,Equinix PA4,93500,CNDC,COGENT,2.41758,48.9016,2020/5/24 11:50
331,France,Paris,114 Rue Ambroise Croizat,Digital Realty (fka. Equinix PA2),93200,CNDC,COGENT,2.35205,48.928,2020/5/24 11:50
332,France,Paris,137 Boulevard Voltaire,Telehouse 2 (Voltaire),75011,CNDC,COGENT,2.38352,48.8563,2020/5/24 11:50
333,France,Paris,"15, Rue du Cap Horn, Les Ulis",COLT - Paris South West,91940,CNDC,COGENT,2.19724,48.6755,2020/5/24 11:50
334,France,Paris,38 Rue des Jeuneurs,Telehouse 1 (JeÃ»neurs),75002,CNDC,COGENT,2.34439,48.87,2020/5/24 11:50
336,France,Reims,3 allÃ©e Thierry Sabine,Hexanet,51686,CNDC,COGENT,4.06374,49.2352,2020/5/24 11:50
338,France,Rennes,Av. Chardonnet - Lorans 30F,Cogent Data Center - Rennes,35000,CDC,COGENT,-1.63623,48.109,2020/5/24 11:50
339,France,Rouen,20 rue Alexandre BarrabÃ©,Cogent Data Center - Rouen,76000,CDC,COGENT,1.07598,49.4291,2020/5/24 11:50
340,France,Sainghin en MÃ©lantois,P.A. du MÃ©lantois - 1681 Rue des Saules,CIV,59262,CNDC,COGENT,3.13651,50.5841,2020/5/24 11:50
341,France,Saint Denis,11-13 Avenue des Arts et Metiers,Interxion PAR5,93200,CNDC,COGENT,2.36246,48.9143,2020/5/24 11:50
342,France,Saint Denis,114 rue Ambroise Croizat,Equinix PA3,93200,CNDC,COGENT,2.35205,48.928,2020/5/24 11:50
343,France,Saint Denis,7-9 avenue des Arts et MÃ©tiers,Interxion PAR3,93200,CNDC,COGENT,2.36241,48.9147,2020/5/24 11:50
344,France,Saint Etienne du Rouvray,Rond-point des Vaches,SFR Netcenter,76 800,CNDC,COGENT,1.10595,49.3679,2020/5/24 11:50
345,France,Saint-Denis,114 rue Ambroise Croizat,Equinix PA2 (fka. IX Europe),93200,CNDC,COGENT,2.35205,48.928,2020/5/24 11:50
346,France,Schiltigheim,46 route de Bischwiller,Cogent Data Center - Strasbourg,67300,CDC,COGENT,7.75226,48.6137,2020/5/24 11:50
347,France,Sophia Antipolis,49 rue Emile Hugues,Euclyde DC1,6560,CNDC,COGENT,7.07559,43.6048,2020/5/24 11:50
348,France,Strasbourg,1 rue du Havre,SFR Netcenter,67000,CNDC,COGENT,7.7824,48.569,2020/5/24 11:50
349,France,Strasbourg,1 rue du Havre,PlusServer DataDock,67000,CNDC,COGENT,7.7824,48.569,2020/5/24 11:50
350,France,Toulouse,125bis ch du Sang de Serp,Cogent Data Center - Toulouse,31000,CDC,COGENT,1.42254,43.6202,2020/5/24 11:50
351,France,Tours,171 rue des Douets,Cogent Data Center - Tours,37100,CDC,COGENT,0.686388,47.4324,2020/5/24 11:50
352,France,Velizy Villacoublay,16 rue Grange Damerose,Cogent Data Center - Paris 2,78140,CDC,COGENT,2.20806,48.7837,2020/5/24 11:50
353,France,Velizy Villacoublay,8 Rue Grange Dame Rose,MGN-Prosodie,78142,CNDC,COGENT,2.20493,48.781,2020/5/24 11:50
354,France,Vennissieux,6-8 rue George Marannes,Lyonix 2E,69200,CNDC,COGENT,4.86185,45.723,2020/5/24 11:50
355,France,Vennissieux,6-8 rue Georges Marannes,SFR Netcenter,69200,CNDC,COGENT,4.86143,45.7229,2020/5/24 11:50
356,France,Vennissieux,6-8 rue Georges Marannes,Lyonix 2A,69200,CNDC,COGENT,4.86143,45.7229,2020/5/24 11:50
357,France,Vennissieux,6-8 rue Georges Marannes,Lyonix 2B,69200,CNDC,COGENT,4.86143,45.7229,2020/5/24 11:50
358,France,Vitry Sur Seine,29 rue Edith Cavell,Iliad Datacenter - DC2,94400,CNDC,COGENT,2.41173,48.7947,2020/5/24 11:50
359,France,Vitry sur Seine,61 rue Julian Grimau,Iliad Datacenter - DC3,94400,CNDC,COGENT,2.37997,48.7753,2020/5/24 11:50
360,Germany,Berlin,Alboinstrasse 36-42,Speedbone,12103,CNDC,COGENT,13.3697,52.4658,2020/5/24 11:50
361,Germany,Berlin,Gradestrasse 40,1&1 Versatel Data Center,12347,CNDC,COGENT,13.4299,52.4526,2020/5/24 11:50
362,Germany,Berlin,Gradestrasse 60,CenturyLink (fka. Level 3),12347,CNDC,COGENT,13.4286,52.453,2020/5/24 11:50
363,Germany,Berlin,LÃ¼tzowstr. 105/106,IPB / Carrier Colo - B,10785,CNDC,COGENT,13.3693,52.502,2020/5/24 11:50
364,Germany,Berlin,Nonnendammallee 15,e-Shelter,13599,CNDC,COGENT,13.238,52.5385,2020/5/24 11:50
365,Germany,Berlin,Wiebestrasse 46,COLT - Berlin 1,D-10553,CNDC,COGENT,13.3189,52.5275,2020/5/24 11:50
366,Germany,Berlin,Wilhelm-von-Siemensstr. 23,MTI/GLH,12227,CNDC,COGENT,13.3813,52.4308,2020/5/24 11:50
367,Germany,Bremen,Hermann-Ritter-Str. 104,MTI/GLH,28359,CNDC,COGENT,8.77376,53.0769,2020/5/24 11:50
368,Germany,Cologne,Hansestrasse 109,Nimblu / PlusServer (fka. Host Europe),51149,CNDC,COGENT,7.05608,50.9127,2020/5/24 11:50
369,Germany,Dortmund,Revierstrasse 3,MTI/GLH,44379,CNDC,COGENT,7.37417,51.5189,2020/5/24 11:50
370,Germany,Dresden,Marie Curie Str. 7,MTI/GLH,1139,CNDC,COGENT,11.4745,48.7816,2020/5/24 11:50
371,Germany,Dusseldorf,Albertstrasse 27,Equinix DU1,40233,CNDC,COGENT,13.3545,52.4856,2020/5/24 11:50
372,Germany,Dusseldorf,Am Gatherhof 44,Myloc,40599,CNDC,COGENT,6.81752,51.2673,2020/5/24 11:50
373,Germany,Dusseldorf,In der Steele 29,Interxion DUS1,40599,CNDC,COGENT,6.86722,51.1883,2020/5/24 11:50
374,Germany,Dusseldorf,In der Steele 37,MTI/GLH,40599,CNDC,COGENT,6.86728,51.1882,2020/5/24 11:50
375,Germany,Dusseldorf,In der Steele 39,Interxion DUS2,40599,CNDC,COGENT,6.8668,51.1877,2020/5/24 11:50
376,Germany,Dusseldorf,Vogelsanger Weg 91,KPN,40470,CNDC,COGENT,6.7987,51.2586,2020/5/24 11:50
377,Germany,Essen,Teilungsweg 28,MTI/GLH,45329,CNDC,COGENT,7.00055,51.4959,2020/5/24 11:50
378,Germany,Frankfurt,Eschborner LandstraÃŸe 100 â€?Building A,e-shelter Frankfurt 1 - Bldg. A,60489,CNDC,COGENT,8.60144,50.1282,2020/5/24 11:50
379,Germany,Frankfurt,Eschborner LandstraÃŸe 100 â€?Building B,e-shelter Frankfurt 1 - Bldg. B,60489,CNDC,COGENT,8.60144,50.1282,2020/5/24 11:50
380,Germany,Frankfurt,Eschborner LandstraÃŸe 100 â€?Building C,e-shelter Frankfurt 1 - Bldg. C,60489,CNDC,COGENT,8.60144,50.1282,2020/5/24 11:50
381,Germany,Frankfurt,Eschborner LandstraÃŸe 100 â€?Building D,e-shelter Frankfurt 1 - Bldg. D,60489,CNDC,COGENT,8.60144,50.1282,2020/5/24 11:50
382,Germany,Frankfurt,Eschborner LandstraÃŸe 100 â€?Building E,e-shelter Frankfurt 1 - Bldg. E,60489,CNDC,COGENT,8.60144,50.1282,2020/5/24 11:50
383,Germany,Frankfurt,Eschborner LandstraÃŸe 100 â€?Building F & G,e-shelter Frankfurt 1 - Bldg. F & G,60489,CNDC,COGENT,8.60144,50.1282,2020/5/24 11:50
384,Germany,Frankfurt,Eschborner LandstraÃŸe 100 â€?Building H,e-shelter FRA 1,60489,CNDC,COGENT,8.60144,50.1282,2020/5/24 11:50
385,Germany,Frankfurt,Eschborner LandstraÃŸe 100 â€?Building T,e-shelter Frankfurt 1 - Bldg. T,60489,CNDC,COGENT,8.60144,50.1282,2020/5/24 11:50
386,Germany,Frankfurt,Eschborner Landstrasse 110,Global Switch,60489,CNDC,COGENT,8.59844,50.1299,2020/5/24 11:50
387,Germany,Frankfurt,Eschborner Landstrasse 110,Evoque FRA1,0,CNDC,COGENT,8.59844,50.1299,2020/5/24 11:50
388,Germany,Frankfurt,Gutleutstrasse 310,Equinix FR7 (fka. Telecity),60327,CNDC,COGENT,8.64417,50.0969,2020/5/24 11:50
389,Germany,Frankfurt,Hanauer Landstr. 302,Interxion FRA1,60314,CNDC,COGENT,8.73545,50.1194,2020/5/24 11:50
390,Germany,Frankfurt,Hanauer Landstr. 320,e-shelter Frankfurt 2,60314,CNDC,COGENT,8.73853,50.1208,2020/5/24 11:50
391,Germany,Frankfurt,Hanauer Landstrasse 264,Interxion - FRA15,60314,CNDC,COGENT,8.73306,50.1188,2020/5/24 11:50
392,Germany,Frankfurt,Hanauer Landstrasse 296a,Interxion FRA7,60314,CNDC,COGENT,8.73469,50.1196,2020/5/24 11:50
393,Germany,Frankfurt,Hanauer Landstrasse 300a,Interxion FRA6,60314,CNDC,COGENT,8.73527,50.12,2020/5/24 11:50
394,Germany,Frankfurt,Hanauer Landstrasse 304,Interxion FRA2,60314,CNDC,COGENT,8.73589,50.1199,2020/5/24 11:50
395,Germany,Frankfurt,Hanauer Landstrasse 308a,Interxion FRA5,60314,CNDC,COGENT,8.73699,50.1205,2020/5/24 11:50
396,Germany,Frankfurt,Hanauer Landstrasse 326,COLT - Frankfurt City,D-60314,CNDC,COGENT,8.74001,50.1218,2020/5/24 11:50
397,Germany,Frankfurt,KleyerstraÃŸe 75 -87,Telehouse Frankfurt,60326,CNDC,COGENT,8.63006,50.097,2020/5/24 11:50
398,Germany,Frankfurt,KleyerstraÃŸe 90,Equinix FR5 (fka. Ancotel),60326,CNDC,COGENT,8.63195,50.0988,2020/5/24 11:50
399,Germany,Frankfurt,Kruppstrasse 121-127,Equinix FR2 (fka. IX Europe),60388,CNDC,COGENT,8.74176,50.1419,2020/5/24 11:50
400,Germany,Frankfurt,Larchenstrasse 110,Equinix FR4,65933,CNDC,COGENT,8.5873,50.0977,2020/5/24 11:50
401,Germany,Frankfurt,Larchenstrasse 110,Equinix FR6,65993,CNDC,COGENT,8.5873,50.0977,2020/5/24 11:50
402,Germany,Frankfurt,Leonhard-Heisswolf-Strasse 4,CyrusOne (fka. Zenium Frankfurt 1),65936,CNDC,COGENT,8.583,50.1281,2020/5/24 11:50
403,Germany,Frankfurt,Lyoner Strasse 15,Cogent Data Center - Frankfurt,60528,CDC,COGENT,8.62711,50.0813,2020/5/24 11:50
404,Germany,Frankfurt,Lyoner Strasse 15,Atricom,60528,COB,COGENT,8.62711,50.0813,2020/5/24 11:50
405,Germany,Frankfurt,Lyoner Strasse 28,Digital Realty (fka. Telecity),60528,CNDC,COGENT,8.62316,50.0817,2020/5/24 11:50
406,Germany,Frankfurt,RebstÃ¶ckerstr. 25-31,ITENOS,60326,CNDC,COGENT,8.63176,50.0997,2020/5/24 11:50
407,Germany,Frankfurt,RebstÃ¶ckerstrasse 25-31,NewTelco,60326,CNDC,COGENT,8.63176,50.0997,2020/5/24 11:50
408,Germany,Frankfurt,Taubenstrasse 7-9,Equinix FR1 (fka. IX Europe),60313,CNDC,COGENT,13.3884,52.5128,2020/5/24 11:50
409,Germany,Frankfurt,Weismueller Strasse 36,Interxion FRA8,60314,CNDC,COGENT,8.7373,50.1193,2020/5/24 11:50
410,Germany,Frankfurt,WeismÃ¼llerstraÃŸe 34,Interxion - FRA14,60314,CNDC,COGENT,8.73808,50.1195,2020/5/24 11:50
411,Germany,Frankfurt,Weismullerstrasse 19,Interxion FRA4,60314,CNDC,COGENT,8.73518,50.1189,2020/5/24 11:50
412,Germany,Frankfurt,Weismullerstrasse 23,Interxion FRA3,60314,CNDC,COGENT,8.73617,50.1195,2020/5/24 11:50
413,Germany,Frankfurt,Weismullerstrasse 29-31,Interxion FRA12,60314,CNDC,COGENT,8.73775,50.1196,2020/5/24 11:50
414,Germany,Frankfurt,Weismullerstrasse 40,Interxion FRA11,60314,CNDC,COGENT,8.73847,50.1198,2020/5/24 11:50
415,Germany,Frankfurt,Weismullerstrasse 42,Interxion FRA13,60314,CNDC,COGENT,8.73873,50.1199,2020/5/24 11:50
416,Germany,Frankfurt,Weissmuellerstrasse 25-27,Interxion FRA9,60314,CNDC,COGENT,8.73696,50.1193,2020/5/24 11:50
418,Germany,Frankfurt,Wilhelm-Fay-StraÃŸe 24,Digital Realty FRA10,65936,CNDC,COGENT,8.58557,50.1272,2020/5/24 11:50
419,Germany,Frankfurt,Wilhelm-Fay-StraÃŸe 24,Cyxtera FRA2,65936,CNDC,COGENT,8.58557,50.1272,2020/5/24 11:50
420,Germany,Hamburg,Flughafen Strasse 54a,Global Connect,22335,CNDC,COGENT,10.0157,53.6387,2020/5/24 11:50
421,Germany,Hamburg,Langenhorner Chaussee 44,e-shelter (fka. Telehouse),22335,CNDC,COGENT,10.0158,53.6372,2020/5/24 11:50
423,Germany,Hamburg,"Wendenstrasse 379, 3rd floor",n@work (fka. Incolocate / Work-IX),20537,CNDC,COGENT,10.0471,53.5509,2020/5/24 11:50
424,Germany,Hessen,Goethering 29 Offenbach Main,Maincubes FRA01,63067,CNDC,COGENT,8.74229,50.1093,2020/5/24 11:50
425,Germany,Mainz,Robert-Bosch-StraÃŸe 48,ITecon,55129,CNDC,COGENT,8.7566,50.2312,2020/5/24 11:50
426,Germany,Mannheim,Hans Thoma Str. 15-17,MTI/GLH,68163,CNDC,COGENT,8.52779,49.4745,2020/5/24 11:50
427,Germany,Munich,Arnulfstrasse 32,KPN,80335,CNDC,COGENT,11.5552,48.1425,2020/5/24 11:50
428,Germany,Munich,Landsberger Strasse 155,Vodafone (fka. C&W/INXS),80687,CNDC,COGENT,11.5266,48.14,2020/5/24 11:50
429,Germany,Munich,Seidlstrasse 3,Equinix MU1 & MU3,D-80335,CNDC,COGENT,11.5558,48.1431,2020/5/24 11:50
430,Germany,Munich,Wamslerstrasse 8,CenturyLink (fka. Level 3),81829,CNDC,COGENT,11.6627,48.1361,2020/5/24 11:50
431,Germany,Nuremberg,DeutschherrnstraÃŸe 15-19,KPN,90429,CNDC,COGENT,11.0636,49.4534,2020/5/24 11:50
432,Germany,Nuremberg,Lina Ammon Str. 22,MTI/GLH,90459,CNDC,COGENT,11.1306,49.4083,2020/5/24 11:50
433,Germany,Nuremberg,Sandreuthstrasse 31,M-Net (fka. Netkom),90441,CNDC,COGENT,11.0612,49.4372,2020/5/24 11:50
434,Germany,Nuremberg,Sigmundstr. 135,Hetzner Data Center,90431,CNDC,COGENT,11.0141,49.4494,2020/5/24 11:50
435,Germany,RÃ¼sselsheim am Main,Karl Landsteiner - Ring 4,e-shelter Frankfurt 3,65428,CNDC,COGENT,8.45027,49.9747,2020/5/24 11:50
436,Germany,Stuttgart,Ruppmannstr. 33,MTI/GLH,70765,CNDC,COGENT,9.11716,48.7257,2020/5/24 11:50
437,Germany,UnterfÃ¶hring,Beta-Strass 1,MTI/GLH,85774,CNDC,COGENT,11.6569,48.1899,2020/5/24 11:50
439,Greece,Athens,"76 Ifestou str, Koropi",Lamda Hellix,19400,CNDC,COGENT,23.8703,37.8714,2020/5/24 11:50
440,Greece,Metamorfosi,37 Ermou str,TI Sparkle (fka. MedNautilus),14452,CNDC,COGENT,23.7304,37.9761,2020/5/24 11:50
442,Greece,Thessaloniki,8 Iatrou Gogousi Str.,Interworks,56429,CNDC,COGENT,22.6563,38.0673,2020/5/24 11:50
443,Hungary,Budapest,Asztalos Sandor u 13,T-Systems (fka. Dataplex),H-1087,CNDC,COGENT,19.1023,47.4932,2020/5/24 11:50
445,Hungary,Budapest,Victor Hugo u 18-22,T-Systems (fka. BIX/Interware),H-1132,CNDC,COGENT,19.0557,47.518,2020/5/24 11:50
446,Ireland,Cherry Orchard,Unit 35 Lavery Avenue (Parkwest Business Park),Interxion DUB1,Dublin 12,CNDC,COGENT,-6.36556,53.3334,2020/5/24 11:50
448,Ireland,Dublin,4033 Citywest Avenue,Keppel DC (fka. Citadell 100),Dublin 24,CNDC,COGENT,-6.42358,53.2878,2020/5/24 11:50
449,Ireland,Dublin,Blanchardstown Corporate Park,Servecentric,Dublin 15,CNDC,COGENT,-6.37321,53.413,2020/5/24 11:50
450,Ireland,Dublin,"Castle Business Park, Nangor Rd.",Interxion DUB3,Dublin 22,CNDC,COGENT,-6.43954,53.3286,2020/5/24 11:50
451,Ireland,Dublin,Clonshaugh Industrial Estate,GTT (fka. Hibernia Atlantic),Dublin 17,CNDC,COGENT,-6.21091,53.4127,2020/5/24 11:50
452,Ireland,Dublin,"Unit 14 Northwest Business Park , Ballycoolin",Equinix DB4 (fka. Telecity),Dublin 15,CNDC,COGENT,-6.35666,53.4112,2020/5/24 11:50
453,Ireland,Dublin,Unit 2 Northwest Business Park Ballycoolin,Equinix DB3 (fka. Telecity),Dublin 15,CNDC,COGENT,-6.34874,53.4109,2020/5/24 11:50
454,Ireland,Dublin,Unit 24 Hume Avenue Park West Business Park,Interxion DUB2,Dublin 12,CNDC,COGENT,-6.36222,53.3309,2020/5/24 11:50
455,Ireland,Dublin,Unit 4027 Kingswood Road City West Business Park,Equinix DB1 (fka. Telecity),Dublin 24,CNDC,COGENT,-6.41732,53.2952,2020/5/24 11:50
456,Italy,Arezzo,Via Gobetti 96,Aruba IT1,52100,CNDC,COGENT,11.838,43.4599,2020/5/24 11:50
457,Italy,Arezzo,via PO #19,Retelit,52100,CNDC,COGENT,11.0816,43.8949,2020/5/24 11:50
459,Italy,Cornaredo,Via Monzoro 105,Data4 - Milan DC01,20010,CNDC,COGENT,9.03162,45.4807,2020/5/24 11:50
460,Italy,Cornaredo,Via Monzoro 105,Data4 - Milan DC03,20010,CNDC,COGENT,9.03162,45.4807,2020/5/24 11:50
461,Italy,Cornaredo,Via Monzoro 105,Data4 DC04,20010,CNDC,COGENT,9.03162,45.4807,2020/5/24 11:50
462,Italy,Cornaredo,Via Monzoro 105,Data4 - Milan DC02,20010,CNDC,COGENT,9.03162,45.4807,2020/5/24 11:50
463,Italy,Florence,Campi Bisenzio VS Quirico 264,Retelit,50013,CNDC,COGENT,11.1441,43.8451,2020/5/24 11:50
464,Italy,Florence,"Via Lucchese 141, Sesto Fiorentino",Playnet / Brain Technology Datacenter,50019,CNDC,COGENT,11.1724,43.8081,2020/5/24 11:50
465,Italy,Genoa,Genova Localita Borzoli CAP,GTT (fka. Interoute),16153,CNDC,COGENT,8.86957,44.4345,2020/5/24 11:50
466,Italy,Milan,Via Caldera 21 -Building E,Infracom Ginevra - Via Caldera E,20153,CNDC,COGENT,8.80267,44.4253,2020/5/24 11:50
467,Italy,Milan,Via Caldera 21 -Building E,Infracom Artu - Via Caldera E,20153,CNDC,COGENT,8.80267,44.4253,2020/5/24 11:50
468,Italy,Milan,Via Caldera 21 -Building E,Infracom Morgana - Via Caldera E,20153,CNDC,COGENT,8.80267,44.4253,2020/5/24 11:50
469,Italy,Milan,Via Caldera 21 -Building E,Infracom Camelot - Via Caldera E,20153,CNDC,COGENT,8.80267,44.4253,2020/5/24 11:50
470,Italy,Milan,Via Caldera 21 -Building E,Infracom Merlino - Via Caldera E,20153,CNDC,COGENT,8.80267,44.4253,2020/5/24 11:50
471,Italy,Milan,Via Caldera 21 Building D,Infracom Romeo e Giulietta - Via Caldera D,20153,CNDC,COGENT,9.10182,45.478,2020/5/24 11:50
472,Italy,Milan,Via Caldera 21 Building D,Infracom Bunker - Via Caldera D,20153,CNDC,COGENT,9.10182,45.478,2020/5/24 11:50
473,Italy,Milan,Via Caldera 21 - Bldg D,MIX,20153,CNDC,COGENT,9.10182,45.478,2020/5/24 11:50
474,Italy,Milan,Via Caldera 21- Bldg. B,seeweb,20153,CNDC,COGENT,9.10182,45.478,2020/5/24 11:50
475,Italy,Milan,Via Caldera 21- Bldg. B,Irideos (fka. KPNQwest Italia),20153,CNDC,COGENT,9.10182,45.478,2020/5/24 11:50
476,Italy,Milan,Via Caldera 21-Bldg. E,Infracom,20153,CNDC,COGENT,9.10182,45.478,2020/5/24 11:50
477,Italy,Milan,Via Savona 125,Equinix ML2,20144,CNDC,COGENT,9.14744,45.4487,2020/5/24 11:50
478,Italy,Milano 3 city Basiglio Mi,Via Franscesco Sforza 13,Equinix ML3,20080,CNDC,COGENT,9.19732,45.4617,2020/5/24 11:50
479,Italy,Padova,Galleria Spagna 28,VSIX,35127,CNDC,COGENT,11.9285,45.3886,2020/5/24 11:50
480,Italy,Pasian di Prato,"Via Spilimbergo 70, 33, 37",InAsset Udine,33037,CNDC,COGENT,13.1104,46.0982,2020/5/24 11:50
481,Italy,Ponte San Pietro,53 Via San Clemente,Aruba IT3 - Global Cloud Data Center (Bergamo),24036,CNDC,COGENT,11.3467,43.8256,2020/5/24 11:50
482,Italy,Rome,Via Dei Tizii 6b,Caspur Namex,185,CNDC,COGENT,12.5119,41.899,2020/5/24 11:50
483,Italy,Rome,Via della Civilta del Lavoro 52,Hitachi Systems CBT / EUR TEL (fka. EURFacility),144,CNDC,COGENT,12.4702,41.8351,2020/5/24 11:50
484,Italy,Siziano,"Via Lombardia, Zona Industriale",Switch - SUPERNAP Italia,27010,CNDC,COGENT,9.25233,45.3997,2020/5/24 11:50
485,Italy,Torino,"Corso Svizzera, 185",IT Gate,10149,CNDC,COGENT,7.65956,45.091,2020/5/24 11:50
486,Japan,Bukyo-ku Tokyo,1-12-3 Suido TY10 Bldg.,Equinix TY10,112-0005,CNDC,COGENT,139.744,35.7086,2020/5/24 11:50
487,Japan,Bunkyo-ku Tokyo,1-12-3 Suido,Equinix TY9,140-0005,CNDC,COGENT,139.744,35.7087,2020/5/24 11:50
489,Japan,Koto-ku Tokyo,1-10-19 Edagawa,Equinix TY5,135-0051,CNDC,COGENT,139.804,35.6556,2020/5/24 11:50
490,Japan,Ota-Ku Tokyo,4F 5-1 Heiwajima 6-Chome,Equinix TY1,1430006,CNDC,COGENT,139.749,35.5764,2020/5/24 11:50
491,Japan,Shinagawa-Ku,2-2-43 Higashi-Shinagawa,Equinix TY7,1400002,CNDC,COGENT,139.748,35.623,2020/5/24 11:50
492,Japan,Shinagawa-Ku,2-2-43 Higashi-Shinagawa,Equinix TY6,1400002,CNDC,COGENT,139.748,35.623,2020/5/24 11:50
493,Japan,Shinigawa-Ku,2-1-7 Higashi-Shinagawa,Equinix TY8,140-0002,CNDC,COGENT,139.748,35.6222,2020/5/24 11:50
494,Japan,Tokyo,"1-9-20, Edagawa,",Equinix TY3,135-0051,CNDC,COGENT,139.803,35.6567,2020/5/24 11:50
495,Japan,Tokyo,3-8-21 Higashi Shinagawa Ku,Equinix TY2,104-0002,CNDC,COGENT,139.748,35.6175,2020/5/24 11:50
496,Japan,Tokyo,Otemachi 1-9-5,Equinix TY4,100 0004,CNDC,COGENT,139.766,35.6884,2020/5/24 11:50
497,Latvia,Riga,Perses iela 2,Perses 2 -Versija,LV-1011,CNDC,COGENT,24.1243,56.952,2020/5/24 11:50
498,Latvia,Riga,Zakusalas Krastmala 1,LVRTC TV Tower,LV 1050,CNDC,COGENT,24.137,56.924,2020/5/24 11:50
500,Luxembourg,Bettembourg,"4, rue A. Graham Bell",LuxConnect DC1,L-3235,CNDC,COGENT,6.11348,49.5052,2020/5/24 11:50
501,Luxembourg,Kayl,210 Rue de Noertzange,EBRC-Kayl Data Center,3670,CNDC,COGENT,6.0522,49.4976,2020/5/24 11:50
502,Luxembourg,Luxembourg,"12 D, Rue Guillaume J. Kroll",European Data Hub,L-1882,CNDC,COGENT,6.10996,49.5795,2020/5/24 11:50
503,Luxembourg,Luxembourg,"43, Boulevard Pierre Frieden",BCE Kirchberg,L-1543,CNDC,COGENT,6.16342,49.6375,2020/5/24 11:50
504,Luxembourg,Rouscht,3 op der poukewiss,LuxConnect DC2 (Roost),L 7795,CNDC,COGENT,6.09007,49.787,2020/5/24 11:50
505,Macedonia,Skopje,23-to Oktomvri 11 A,Akton,1000,CNDC,COGENT,21.4348,41.9912,2020/5/24 11:50
506,Macedonia,Skopje,Nikola Parapunov NN,Telesmart,1000,CNDC,COGENT,21.3833,42.0064,2020/5/24 11:50
507,Mexico,"Colonia San Mateo Cuautepec, TultitlÃ¡n","Boulevard Benito JuÃ¡rez Nr. 20,",KIO Networks MEX5,54948,CNDC,COGENT,-99.1526,19.6281,2020/5/24 11:50
508,Mexico,Guadalajara,"Av. Chapultepec No. 236 Col. Americana, Sector Ju",Marcatel Guadalajara,44160,CNDC,COGENT,-103.369,20.6726,2020/5/24 11:50
509,Mexico,Mexico City,Blv. Magnocentro #6,KIO MEX4 (Interlomas),CP 52760,CNDC,COGENT,-99.2721,19.4021,2020/5/24 11:50
510,Mexico,Mexico City,Poniente 122 No. 551-A Col. Industiral Vallejo,Marcatel Mexico City,2300,CNDC,COGENT,-99.1574,19.4867,2020/5/24 11:50
511,Mexico,Mexico City,Prologacion Paseo de la Reforma No. 5396,KIO ll (Fase 5 & 6) Mexico,5000,CNDC,COGENT,-99.2805,19.3637,2020/5/24 11:50
512,Mexico,Mexico D.F.,Prologacion Paseo de la Reforma No. 5287,"KIO l (Fase 1, 2, 3 & 4) Mexico",5000,CNDC,COGENT,-99.1457,19.4382,2020/5/24 11:50
513,Mexico,Miguel Hidalgo,Sierra Candela #111 Col Lomas de Chapultepec,MCM Data Center,11000,CNDC,COGENT,-99.2116,19.4321,2020/5/24 11:50
514,Mexico,Monterrey,Av. San Jeronimo 210 Pte. Col. San Jeronimo,Marcatel Monterey,64640,CNDC,COGENT,-100.355,25.6738,2020/5/24 11:50
515,Mexico,Queretaro,"Av de la Loma 125, Col. San Pablo",Marcatel Queretaro,76130,CNDC,COGENT,-100.414,20.6192,2020/5/24 11:50
518,Moldova,Chisinau,Decebal Boulevard 2,Decebal Boulevard 2,MD-2001,CNDC,COGENT,28.8587,47.0086,2020/5/24 11:50
520,Netherlands,Alblasserdam,Van Coulsterweg 6,Dataplace,2952 CB,CNDC,COGENT,4.67915,51.8459,2020/5/24 11:50
521,Netherlands,Amsterdam,Cessnalaan 50,Interxion AMS7,1119,CNDC,COGENT,4.76504,52.2812,2020/5/24 11:50
522,Netherlands,Amsterdam,Duivendrechtsekade 80A,Equinix AM6 (fka. Telecity 6),1096 AH,CNDC,COGENT,4.93365,52.3373,2020/5/24 11:50
523,Netherlands,Amsterdam,Gyroscoopweg 2E & 2F,Equinix AM8 (fka. Telecity 3),1042 AB,CNDC,COGENT,4.83874,52.3968,2020/5/24 11:50
524,Netherlands,Amsterdam,Gyroscoopweg 58-62,Interxion AMS1/AMS4,1042 AC,CNDC,COGENT,4.84235,52.3999,2020/5/24 11:50
525,Netherlands,Amsterdam,Johan Huizingalaan 759,Global Switch,1066 VH,CNDC,COGENT,4.82871,52.344,2020/5/24 11:50
526,Netherlands,Amsterdam,Johan Huizingalaan 759,Evoque AMS1,1066,CNDC,COGENT,4.82871,52.344,2020/5/24 11:50
527,Netherlands,Amsterdam,Kabelweg 48a,The Datacenter Group,1014 BB,CNDC,COGENT,4.84921,52.3929,2020/5/24 11:50
528,Netherlands,Amsterdam,Keienbergweg 22,Datacenter.com,1101,CNDC,COGENT,4.93557,52.3095,2020/5/24 11:50
529,Netherlands,Amsterdam,Kuiperbergweg 13,Equinix AM7 (fka Telecity 2),1101 AE,CNDC,COGENT,4.93801,52.3031,2020/5/24 11:50
530,Netherlands,Amsterdam,Kuiperbergweg 13- Building 2,Equinix AM 7.2,1101 AE,CNDC,COGENT,4.93801,52.3031,2020/5/24 11:50
531,Netherlands,Amsterdam,Lemelerbergweg 28,Switch Datacenters AMS1,1101 AH,CNDC,COGENT,4.93948,52.3039,2020/5/24 11:50
532,Netherlands,Amsterdam,Luttenbergweg 4,Equinix AM1 & AM2,1101 EC,CNDC,COGENT,4.94312,52.2999,2020/5/24 11:50
533,Netherlands,Amsterdam,Paul van Vlissingenstraat 16,Eunetworks,1096 BK,CNDC,COGENT,4.91944,52.3334,2020/5/24 11:50
534,Netherlands,Amsterdam,Schepenbergweg 42,Equinix AM5 (fka. Telecity Southeast AMS 5),1105 AT,CNDC,COGENT,4.94505,52.2934,2020/5/24 11:50
535,Netherlands,Amsterdam,Science Park 105,NIKHEF,1098 XG,CNDC,COGENT,4.95084,52.3564,2020/5/24 11:50
536,Netherlands,Amsterdam,Science Park 120,DRT Data Tower (fka. Telecity AMS1),1098 XG,CNDC,COGENT,4.95088,52.3572,2020/5/24 11:50
538,Netherlands,Amsterdam,Science Park 610,Equinix AM4,1098 XH,CNDC,COGENT,4.96096,52.3549,2020/5/24 11:50
539,Netherlands,Amsterdam,Sciencepark 610,Equinix AM3,1098XH,CNDC,COGENT,4.96096,52.3549,2020/5/24 11:50
540,Netherlands,Amsterdam,Wenckebachweg 127,DRT Amstel Business Park (fka. Telecity AMS4),1096 AM,CNDC,COGENT,4.93189,52.3365,2020/5/24 11:50
541,Netherlands,Delft,Heertjeslaan 1,The Datacenter Group,2629 JG,CNDC,COGENT,4.37752,51.9871,2020/5/24 11:50
542,Netherlands,Enschede,Auke Vleerstraat 1,Equinix EN1,7521 PE,CNDC,COGENT,6.84966,52.237,2020/5/24 11:50
543,Netherlands,Haarlem,J.W. Lucasweg 35,Iron Mountain (fka. EvoSwitch AMS1),2031 BE,CNDC,COGENT,4.66375,52.3925,2020/5/24 11:50
544,Netherlands,Halfweg,Haarlemmerstraatweg 133-135,Cogent Data Center - Amsterdam,1165 MK,CDC,COGENT,4.7167,52.3845,2020/5/24 11:50
545,Netherlands,Hengelo,Barnsteenstraat 15,Previder 2,7554 TC,CNDC,COGENT,6.5291,53.2277,2020/5/24 11:50
546,Netherlands,Hengelo,Expolaan 50,Previder 1,7556 BE,CNDC,COGENT,6.76738,52.2849,2020/5/24 11:50
547,Netherlands,Hilversum,Koos Postemalaan 2 (Mediapark),Red Bee (fka. Ericsson) - Media Gateway Bldg. C,1217 ZC,CNDC,COGENT,5.16937,52.2412,2020/5/24 11:50
548,Netherlands,Naaldwijk,Industriestraat 24,Greenhouse Datacenter,2671CT,CNDC,COGENT,4.21083,52.0012,2020/5/24 11:50
549,Netherlands,Roosendaal,9 Argonweg,Colt Rotterdam Data Center,4706 NR,CNDC,COGENT,5.35708,52.1703,2020/5/24 11:50
550,Netherlands,Rotterdam,Schuttevaerweg 48,bytesnet (fka. RIX),CK 3044,CNDC,COGENT,4.41405,51.9301,2020/5/24 11:50
551,Netherlands,Rotterdam,"Van Nelleweg 1, 3044 BC Rotterdam",smartdc,3044 BC,CNDC,COGENT,4.43169,51.9235,2020/5/24 11:50
552,Netherlands,Rozenburg,Pudongweg 37,Interxion AMS8,1437EM,CNDC,COGENT,4.74314,52.2758,2020/5/24 11:50
553,Netherlands,Schiphol,Cateringweg 5,Verizon NAP of Amsterdam (fka. Terremark),1118 AM,CNDC,COGENT,4.80215,52.3226,2020/5/24 11:50
554,Netherlands,Schiphol-Rijk,Capronilaan 2,maincubes (fka. Easynet),1119 NR,CNDC,COGENT,4.77286,52.2822,2020/5/24 11:50
555,Netherlands,Schiphol-Rijk,Cessnalaan 1-33,Interxion AMS3,1119 NJ,CNDC,COGENT,4.77086,52.2841,2020/5/24 11:50
556,Netherlands,Schiphol-Rijk,Tupolevlaan 101,Interxion AMS5,1119 PA,CNDC,COGENT,4.75385,52.2803,2020/5/24 11:50
557,Netherlands,The Hague,Moezel 5,Data Facilities The Hague,2491 CV,CNDC,COGENT,4.38667,52.0647,2020/5/24 11:50
558,Netherlands,Zwolle,"Telfordstraat 3, 8013 RL Zwolle",Equinix ZW1,8013 RL,CNDC,COGENT,6.1426,52.4884,2020/5/24 11:50
559,Norway,Oslo,Gaustadalleen 21,Oslo Innovation Center / Forskningsparken AS,349,CNDC,COGENT,10.7164,59.9423,2020/5/24 11:50
560,Norway,Oslo,Hans Moller Gasmanns Vei 9,Oslo Internet Exchange,598,CNDC,COGENT,10.835,59.9384,2020/5/24 11:50
561,Norway,Oslo,Hans Moller Gasmanns Vei 9,Verizon,598,CNDC,COGENT,10.835,59.9384,2020/5/24 11:50
562,Norway,Oslo,Hans Moller Gasmanns Vei 9,Oslo Internet Exchange,598,CNDC,COGENT,10.835,59.9384,2020/5/24 11:50
563,Norway,Oslo,Selma Ellefsens Vei,Digiplex,581,CNDC,COGENT,10.8059,59.9258,2020/5/24 11:50
564,Poland,Katowice,Adamskiego 7,Quicktel / 4 DataCenter 4DC1,40-069,CNDC,COGENT,19.0031,50.2536,2020/5/24 11:50
565,Poland,Katowice,ul Jordana 25,Quicktel / 4 DataCenter 4DC2,40-056,CNDC,COGENT,19.0159,50.2513,2020/5/24 11:50
566,Poland,Krakow,Biskupinska 19,Netia,30732,CNDC,COGENT,20.0058,50.0304,2020/5/24 11:50
567,Poland,Krakow,Krolewska 57,3S Fibertech,33-332,CNDC,COGENT,19.9156,50.0741,2020/5/24 11:50
568,Poland,Pozan,Adama Kreglewskiego 11th Street,Beyond Data Center 2,61-248,CNDC,COGENT,16.9786,52.377,2020/5/24 11:50
569,Poland,Poznan,Polwiejska 42,Beyond Data Center 1,61-888,CNDC,COGENT,16.9281,52.4011,2020/5/24 11:50
570,Poland,Poznan,Wieniawskiego 17/19,PCSS Data Center,61-713,CNDC,COGENT,16.9178,52.4114,2020/5/24 11:50
571,Poland,Warsaw,Al. Jerozolimskie 65/79,Equinix WA1 (fka. PLIX),00-697 Warsaw,CNDC,COGENT,21.0042,52.2277,2020/5/24 11:50
572,Poland,Warsaw,"Al. Jerozolimskie 65/79 Floor ""+3""","LIM Floor ""+3""",00-697 Warsaw,CNDC,COGENT,21.0042,52.2277,2020/5/24 11:50
573,Poland,Warsaw,Al. Jerozolimskie 65/79 Floor -2,"LIM DC Floor ""-2""",00-697 Warsaw,CNDC,COGENT,21.005,52.2277,2020/5/24 11:50
574,Poland,Warsaw,Konstruktorska 5,Atman WAW-2 (fka. Thinx Poland),02-673,CNDC,COGENT,20.9986,52.1853,2020/5/24 11:50
575,Poland,Warsaw,Piekna 11,Netia,00-549,CNDC,COGENT,21.0198,52.2234,2020/5/24 11:50
576,Poland,Warsaw,U.l. Piekna 11 A,Lukman Multitech,lok12,CNDC,COGENT,21.0191,52.2235,2020/5/24 11:50
577,Poland,Wysogotowo,Wierzbowa 84,DC INEA,62-081,CNDC,COGENT,16.7854,52.4075,2020/5/24 11:50
578,Portugal,Campanha (Porto),Largo EstaÃ§Ã£o de CampanhÃ£,Infraestruturas de Portugal (fka. REFER),4300-173 Port,CNDC,COGENT,-8.58586,41.1488,2020/5/24 11:50
579,Portugal,Lisbon,"AV. SEVERIANO FALCÃƒO NÂº 14, PRIOR VELHO",Equinix LS1 (fka. Telvent),2685-338 PRIO,CNDC,COGENT,-9.12345,38.7879,2020/5/24 11:50
580,Romania,Bucharest,6A Dimitrie Pompeiu,NXDATA-2 (IEMI Bldg.),20337,CNDC,COGENT,26.1148,44.4798,2020/5/24 11:50
581,Romania,Bucharest,8 Dimitrie Pompeiu,NXDATA-1 (Feper Bldg.),20337,CNDC,COGENT,26.1177,44.4805,2020/5/24 11:50
582,Romania,Bucharest,Sos. Fabrica de Glucoza 11 B,M247 Europe,20331,CNDC,COGENT,26.1158,44.4778,2020/5/24 11:50
584,Romania,Drobeta Turnu Severin,Str. IC Bratianu 35A,Data Zyx,220226,CNDC,COGENT,23.5946,46.7691,2020/5/24 11:50
585,Romania,Timisoara,"Grigore Alexandrescu, no 123",Orange MSC,10624,CNDC,COGENT,26.0927,44.4521,2020/5/24 11:50
586,Serbia,Belgrade,Boulevar Vojvode Misica 37,BeotelNet,11100,CNDC,COGENT,20.4388,44.7955,2020/5/24 11:50
587,Serbia,Belgrade,Bulevar Mihaila Pupina 6 (Usce Tower I building at,Akton,11070,CNDC,COGENT,20.4372,44.8162,2020/5/24 11:50
588,Serbia,Belgrade,Omladinskih brigada 92,Telenor Serbia,11070,CNDC,COGENT,20.4115,44.8222,2020/5/24 11:50
589,Singapore,Jurong,29A International Business Park,Cyxtera SIN2 (fka. CenturyLink SG2),609934,CNDC,COGENT,103.746,1.32776,2020/5/24 11:50
590,Singapore,Jurong,29A International Business Park,Digital Realty Jurong,609934,CNDC,COGENT,103.746,1.32776,2020/5/24 11:50
591,Singapore,Jurong,29A International Business Park 520,CyrusOne - International Business Park,609934,CNDC,COGENT,103.746,1.3277,2020/5/24 11:50
592,Singapore,Singapore,15 Pioneer Walk,Equinix SG2,627753,CNDC,COGENT,103.695,1.32182,2020/5/24 11:50
593,Singapore,Singapore,2 Tai Seng Avenue,Global Switch,534408,CNDC,COGENT,103.894,1.33672,2020/5/24 11:50
594,Singapore,Singapore,"20 Ayer Rajah Crescent, Ayer Rajah Industrial Park",Equinix SG1,139964,CNDC,COGENT,103.79,1.29546,2020/5/24 11:50
595,Singapore,Singapore,23 Tai Seng Drive,Racks Central,535224,CNDC,COGENT,103.893,1.33702,2020/5/24 11:50
596,Singapore,Singapore,26A Ayer Rajah Crescent,Equinix SG3,139963,CNDC,COGENT,103.791,1.29618,2020/5/24 11:50
597,Singapore,Singapore,"750D Viva Business Park, #07-04 Chai Chee Road",1-Net - 750D,469005,CNDC,COGENT,103.921,1.32259,2020/5/24 11:50
598,Singapore,Singapore,"750E Viva Business Park, #06-00 Chai Chee Road",1-Net - 750E,469005,CNDC,COGENT,103.922,1.32349,2020/5/24 11:50
599,Singapore,Singapore,"750E Viva Business Park, #06-00 Chai Chee Road",Evoque SNG1,469005,CNDC,COGENT,103.922,1.32349,2020/5/24 11:50
600,Slovakia,Bratislava,Istrijska ulica 26,SWAN / Perpetuus DC,841 07,CNDC,COGENT,16.9742,48.2104,2020/5/24 11:50
601,Slovakia,Bratislava,Kopcianska 18,SITELPOP2,85101,CNDC,COGENT,17.0951,48.1181,2020/5/24 11:50
602,Slovakia,Bratislava,Kopcianska 20 c,SITELPOP1,85101,CNDC,COGENT,17.0956,48.1187,2020/5/24 11:50
603,Slovakia,Bratislava,Kopcianska 92/d,Datacube,85101,CNDC,COGENT,17.0871,48.1135,2020/5/24 11:50
604,Slovakia,Bratislava,Namestie Slobody 17,SIX - Computing Centre STU,81243,CNDC,COGENT,17.1126,48.1512,2020/5/24 11:50
607,Slovenia,Ljubljana,Dunajska Cesta 5,Datacenter DC01,SI-1000,CNDC,COGENT,14.5061,46.0594,2020/5/24 11:50
608,Slovenia,Ljubljana,Å martinska 106,Telemach (fka. TuÅ¡Mobil),SI-1000,CNDC,COGENT,14.5375,46.0678,2020/5/24 11:50
609,Slovenia,Ljubljana,TehnoloÅ¡ki park 19,Datacenter DC02,SI-1000,CNDC,COGENT,14.4606,46.0495,2020/5/24 11:50
610,Slovenia,Ljubljana,Tivolska 50,Telemach (fka. LIX),SI-1000,CNDC,COGENT,14.5052,46.0574,2020/5/24 11:50
611,South Korea,Seoul,"13 Eonju-ro 30-gil, Gangnam-gu",COI Hub - SEL01 - KINX Dogok IDC,0,CNDC,COGENT,127.051,37.4881,2020/5/24 11:50
612,Spain,Alcobendas,Valgrande 6,Equinix MD2 (fka. Telvent Carrier House 2),28108,CNDC,COGENT,-3.6492,40.5367,2020/5/24 11:50
613,Spain,Barcelona,Acero 30-32,Equinix BA1 (fka. Itonic/Telvent Carrier House),8014,CNDC,COGENT,2.13746,41.3516,2020/5/24 11:50
614,Spain,Bilbao,"Poligono Artunduaga Parcela 21 / 4897 Basauri, Biz",GTT (fka. Interoute),48970,CNDC,COGENT,-2.87895,43.2193,2020/5/24 11:50
615,Spain,Gijon,Avda de Oviedo 26,Avda de Oviedo 26,33211,CNDC,COGENT,-5.68999,43.5185,2020/5/24 11:50
619,Spain,L Hospitalet de Llobregat,C/Pablo Iglesias 56,BitNap,8908,CNDC,COGENT,2.12978,41.3486,2020/5/24 11:50
620,Spain,Madrid,Albasanz 71,Interxion MAD1,28037,CNDC,COGENT,-3.62121,40.4393,2020/5/24 11:50
621,Spain,Madrid,"Av. Manoteras 42, Calle 3",COLT - Madrid 1,28050,CNDC,COGENT,-3.66493,40.4862,2020/5/24 11:50
622,Spain,Madrid,"C/ Jose Bardasano Baos, 9 Edificio Gorbea 3, Plant","Jose Bardasano Baos, 9 Edificio Gorbea 3",28016,COB,COGENT,-3.67222,40.4676,2020/5/24 11:50
624,Spain,Madrid,"Calle Maria Tubau, 8 4a floor",Ibercom Datacenter,28050,CNDC,COGENT,-3.6735,40.5139,2020/5/24 11:50
625,Spain,Madrid,Calle Marzo 18,IPCore,28022,CNDC,COGENT,-3.56992,40.4459,2020/5/24 11:50
626,Spain,Madrid,Emilio MuÃ±oz 49,Interxion MAD3,28037,CNDC,COGENT,-3.62408,40.4339,2020/5/24 11:50
627,Spain,Madrid,Isabel Colbrand 6-8,Equinix MD1 (fka. Telvent Carrier House 1),28050,CNDC,COGENT,-3.67189,40.5135,2020/5/24 11:50
628,Spain,Madrid,Julian Camarillo 29A,Protec,0,CNDC,COGENT,-3.62904,40.4346,2020/5/24 11:50
629,Spain,Madrid,Mesena 80,Espanix,28230,CNDC,COGENT,-3.66282,40.4658,2020/5/24 11:50
630,Spain,Madrid,Paseo de la Castellana 95,Europa Tower,28020,COB,COGENT,-3.69167,40.4517,2020/5/24 11:50
631,Spain,Madrid,Plaza Pablo Ruiz Picasso s/n,Picasso Tower,28020,COB,COGENT,-3.69315,40.451,2020/5/24 11:50
634,Spain,Malaga,Eduardo Queipo de Llano Caballero 59,Malaga Data Center / Quenture,29004,CNDC,COGENT,-4.47458,36.6948,2020/5/24 11:50
636,Spain,San Sebastian,Parque Empresarial Zuatzu Edificio Easo 2,Ibercom Datacenter,20018,CNDC,COGENT,-2.00512,43.2968,2020/5/24 11:50
639,Spain,"Sevilla , PCT Cartuja","Calle Isaac Newton, 3 Edificio Bluenet",Nethits (fka Neutral Route),41092,CNDC,COGENT,-6.00508,37.4085,2020/5/24 11:50
640,Spain,Valencia,Louis Pasteur 6,PRORED Datacenter,46980,CNDC,COGENT,-0.467256,39.554,2020/5/24 11:50
641,Spain,Valencina de la ConcepciÃ³,"Avda. de AndalucÃ­a, s/n. Dolmen de la Pastora",Axion,41907,CNDC,COGENT,-6.0697,37.4055,2020/5/24 11:50
642,Spain,Vigo,Avda Madrid s/n,Ufinet (fka. GNFT),36214,CNDC,COGENT,-3.37492,40.4526,2020/5/24 11:50
643,Sweden,Bromma,MariehÃ¤llsvÃ¤gen 36,Equinix SK1,16865,CNDC,COGENT,17.9566,59.3625,2020/5/24 11:50
644,Sweden,Kista,Esbogatan 11,Interxion STO1,164 74,CNDC,COGENT,17.9162,59.4227,2020/5/24 11:50
645,Sweden,Limhamn,Krossverksgatan 15,IP-Only Malmo,216 16,CNDC,COGENT,12.9278,55.5745,2020/5/24 11:50
646,Sweden,Skondal,Kvastvagen 25-29,Equinix SK2,12862,CNDC,COGENT,18.1054,59.2637,2020/5/24 11:50
647,Sweden,SpÃ¥nga,FinspÃ¥ngsgatan 48,Equinix SK3,163 53,CNDC,COGENT,17.8752,59.3909,2020/5/24 11:50
648,Sweden,Stockholm,Artillerigatan 60,Stokab KN7,11445,CNDC,COGENT,18.0851,59.3391,2020/5/24 11:50
649,Sweden,Upplands Vasby,Smedbyvagen 6,Digiplex,19430,CNDC,COGENT,17.8974,59.5099,2020/5/24 11:50
650,Switzerland,Basel,Margarethenstrasse 40,IWB Data Center,4008,CNDC,COGENT,7.58371,47.5466,2020/5/24 11:50
651,Switzerland,Bern,Hochfeldstrasse 114,NTS Workspace - Colobern North,3012,CNDC,COGENT,7.43457,46.9595,2020/5/24 11:50
652,Switzerland,Crissier,17 rue de la Vernie,Brainserve,1023,CNDC,COGENT,6.57338,46.5466,2020/5/24 11:50
653,Switzerland,Geneva,6 rue de la confederation,Equinix GV1,1204,CNDC,COGENT,6.14425,46.2036,2020/5/24 11:50
654,Switzerland,Le Lignon,48 Route du Bois-des-Freres,Equinix GV2,1219,CNDC,COGENT,6.09813,46.2079,2020/5/24 11:50
655,Switzerland,Lupfig,Industriestrasse 33,greenDatacenter - Zurich-West 1+2,5242,CNDC,COGENT,8.21073,47.4464,2020/5/24 11:50
656,Switzerland,Oberengstringen,Almendstasse 9-13,Equinix ZH 5,8102,CNDC,COGENT,8.27231,47.3564,2020/5/24 11:50
657,Switzerland,Prattln,Guterstrasse 72,ColoBÃ¢le,4133,CNDC,COGENT,7.58417,47.5468,2020/5/24 11:50
658,Switzerland,Rumlang,Hofwisenstrasse 56,e-shelter Zurich,CH-8153,CNDC,COGENT,8.54043,47.4495,2020/5/24 11:50
659,Switzerland,Schlieren,Unterrohrstrasse 4,greenDatacenter - Zurich-City 2,CH-8952,CNDC,COGENT,8.43776,47.4036,2020/5/24 11:50
660,Switzerland,Zurich,Badenerstrasse 569,Colozueri.ch,8048,CNDC,COGENT,8.4956,47.3834,2020/5/24 11:50
661,Switzerland,Zurich,Badenerstrasse 820,COLT,8048,CNDC,COGENT,8.4729,47.3938,2020/5/24 11:50
662,Switzerland,Zurich,Josefstrasse 225,Equinix ZH2,8005,CNDC,COGENT,8.5202,47.3877,2020/5/24 11:50
663,Switzerland,Zurich,Josefstrasse 225,Equinix ZH (fka. ICT Center),8005,CNDC,COGENT,8.5202,47.3877,2020/5/24 11:50
664,Switzerland,Zurich,Josefstrasse 225,Equinix ZH4,8005,CNDC,COGENT,8.5202,47.3877,2020/5/24 11:50
665,Switzerland,Zurich,Sagereistrasse 35 (FKA 29),Interxion ZUR1,CH-8152,CNDC,COGENT,8.55731,47.4327,2020/5/24 11:50
666,Taiwan,Taipei,"248/250 Yang-Guang Street, Nei-hu District",Chief Telecom LY Building,0,CNDC,COGENT,121.578,25.0735,2020/5/24 11:50
667,Taiwan,Taipei,"NO. 37, Lane 188, Ruiguang Road",Chief Telecom HD Building,114,CNDC,COGENT,121.578,25.0746,2020/5/24 11:50
668,Turkey,Istanbul,"Buyukdere Cad. Street No. 121, Ercan Han Kat",Turknet Data Center,34394,CNDC,COGENT,29.005,41.0678,2020/5/24 11:50
669,Turkey,Istanbul,Cobancesme Mah. Kimiz Sok. No: 40 Bahcelievler,Comnet Bilgi Iletisim,34196,CNDC,COGENT,28.8259,40.9935,2020/5/24 11:50
670,Turkey,Istanbul,Kimiz SOK No: 30 Cobancesme Bahcellievler,Superonline Yenibosna,34194,CNDC,COGENT,28.8271,40.9935,2020/5/24 11:50
671,Turkey,Istanbul,Kimiz SOK No: 30 Cobancesme Bahcellievler,TI Sparkle (fka. MedNautilus),34194,CNDC,COGENT,28.8271,40.9935,2020/5/24 11:50
676,United Kingdom,Belfast,Titanic Quarter Queens Road,GTT (fka. Hibernia Atlantic),BT3 9DT,CNDC,COGENT,-5.90924,54.6075,2020/5/24 11:50
677,United Kingdom,Bershire,"14 Liverpool Road, Slough Trading Estate",VIRTUS DC LON4 (Slough - fka. Infinity SDC),SL1 4QZ,CNDC,COGENT,-0.622406,51.5231,2020/5/24 11:50
678,United Kingdom,Cambridge,"Newton House, Cambridge Business Park, Cowley Road",Redcentric,CB4 0WZ,CNDC,COGENT,0.150961,52.23,2020/5/24 11:50
679,United Kingdom,Edinburgh,Clocktower Estate South Gyle Crescent,Pulsant (fka. Scolocate Phase II),EH12 9LB,CNDC,COGENT,-3.29451,55.9293,2020/5/24 11:50
680,United Kingdom,Enfield,"Unit 3, 13 Crown Road, Enfield EN1 1TX",VIRTUS DC LON1 (Enfield),EN1 1TX,CNDC,COGENT,-0.0536188,51.651,2020/5/24 11:50
681,United Kingdom,Farnborough,"Cody Technology Park, Ively Road",Datum DataCentres,GU14 0LX,CNDC,COGENT,-0.792156,51.2795,2020/5/24 11:50
682,United Kingdom,Glasgow,88 Middlesex Street,Iomart DC1,G41 1EE,CNDC,COGENT,-0.0768963,51.5169,2020/5/24 11:50
683,United Kingdom,Leeds,11-15 Hunslet Road,AQL / Salem Chapel,LS10 1JQ,CNDC,COGENT,-1.54087,53.7931,2020/5/24 11:50
684,United Kingdom,London,1 Muirfield Crescent,Cloud House London,E14 9SZ,CNDC,COGENT,-0.0185106,51.4961,2020/5/24 11:50
685,United Kingdom,London,"1 Olivers Yard, 55-71 City Road 3 Floors",Digital Realty (fka. Telecity),EC1Y 1HQ,CNDC,COGENT,-0.0873443,51.5242,2020/5/24 11:50
686,United Kingdom,London,11 Hanbury Street,Interxion LON2,E1 6QR,CNDC,COGENT,-0.0725998,51.5205,2020/5/24 11:50
687,United Kingdom,London,11 Hanbury Street,Interxion LON1 (Brick Lane),E1 6QR,CNDC,COGENT,-0.0725998,51.5205,2020/5/24 11:50
688,United Kingdom,London,14 Coriander Avenue,Telehouse East,E14 2AA,CNDC,COGENT,-0.0028934,51.5116,2020/5/24 11:50
689,United Kingdom,London,14 Coriander Avenue,Telehouse North 2,E14 2AA,CNDC,COGENT,-0.0028934,51.5116,2020/5/24 11:50
690,United Kingdom,London,14 Coriander Avenue,Telehouse North,E14 2AA,CNDC,COGENT,-0.0028934,51.5116,2020/5/24 11:50
691,United Kingdom,London,14 Coriander Avenue,Telehouse West,E14 2AA,CNDC,COGENT,-0.0028934,51.5116,2020/5/24 11:50
692,United Kingdom,London,20 Mastmaker Court,Cogent Data Center - London,E149UB,CDC,COGENT,-0.0213869,51.4997,2020/5/24 11:50
693,United Kingdom,London,215 Marsh Wall,Digital Realty Meridian Gate (fka. Telecity 2),E14 9FJ,CNDC,COGENT,-0.0117521,51.5006,2020/5/24 11:50
694,United Kingdom,London,227 Marsh Wall,Digital Realty Sovereign House (fka. Telecity 3),E14 9SD,CNDC,COGENT,-0.010743,51.4997,2020/5/24 11:50
695,United Kingdom,London,227 Marsh Wall,Evoque LHR1,E14 9SD,CNDC,COGENT,-0.010743,51.4997,2020/5/24 11:50
696,United Kingdom,London,240 East India Dock Road,Global Switch - London North,E14 9YY,CNDC,COGENT,-0.0060816,51.5111,2020/5/24 11:50
697,United Kingdom,London,260 Goswell Road,CenturyLink (fka. Level 3),EC1V7EB,CNDC,COGENT,-0.101035,51.5287,2020/5/24 11:50
698,United Kingdom,London,3 Nutmeg Lane,Global Switch - London East,E14 2AX,CNDC,COGENT,-0.0034943,51.511,2020/5/24 11:50
699,United Kingdom,London,36-43 Great Sutton Street,Volta Data Centres,EC1V 0AB,CNDC,COGENT,-0.101148,51.5232,2020/5/24 11:50
700,United Kingdom,London,47 Millharbour,Digital Realty (fka. Telecity Millharbour),E14 9TR,CNDC,COGENT,-0.018944,51.4969,2020/5/24 11:50
701,United Kingdom,London,5-6 Greenwich View Place,Telstra LHC (London Hosting Centre),E14 9NN,CNDC,COGENT,-0.0182725,51.4946,2020/5/24 11:50
702,United Kingdom,London,5-6 Greenwich View Place,Cyxtera LHR2-B (fka LO4),E14 9NN,CNDC,COGENT,-0.0182725,51.4946,2020/5/24 11:50
703,United Kingdom,London,6 Braham St,CenturyLink (fka. Level 3),E1 8EE,CNDC,COGENT,-0.0725399,51.5139,2020/5/24 11:50
704,United Kingdom,London,6-7 Harbour Exchange Sq,Equinix LD8 (fka. Telecity 1),E14 9GE,CNDC,COGENT,-0.0143605,51.4981,2020/5/24 11:50
705,United Kingdom,London,65 Clifton St,Telehouse Metro,EC2V 7QP,CNDC,COGENT,-0.0834298,51.5223,2020/5/24 11:50
706,United Kingdom,London,69-77 Paul Street,Epsilon,EC2A 4NW,CNDC,COGENT,-0.0845184,51.5249,2020/5/24 11:50
707,United Kingdom,London,7 Greenwich View Place,Cyxtera LHR2-A (fka LO3),E149NN,CNDC,COGENT,-0.0187046,51.4947,2020/5/24 11:50
708,United Kingdom,London,8-9 Harbour Exchange,Equinix LD8 (fka. Telecity HEX),E14 9GE,CNDC,COGENT,-0.0152639,51.4982,2020/5/24 11:50
709,United Kingdom,London,80 Clifton Street,Redcentric (fka. City Lifeline),EC2A4HB,CNDC,COGENT,-0.0829156,51.523,2020/5/24 11:50
710,United Kingdom,London,Dray Walk,Interxion LON03,E1 6QL,CNDC,COGENT,-0.0725798,51.521,2020/5/24 11:50
711,United Kingdom,London,"Unit 1- 20 Powergate Business Park, Volt Avenue",Equinix LD9,NW10 6PW,CNDC,COGENT,-0.25736,51.5318,2020/5/24 11:50
712,United Kingdom,London,"Unit 11 Matrix Coronation Road, Park Royal",Equinix LD3,NW10 7PH,CNDC,COGENT,-0.27746,51.5288,2020/5/24 11:50
714,United Kingdom,London,"Units 9-13 Powergate, Volt Avenue",COLT - London IDC3 Power Gate,NW10 6PN,CNDC,COGENT,-0.25736,51.5318,2020/5/24 11:50
716,United Kingdom,Manchester,"Delta House, Wavell Road",Teledata Delta House,M22 5QZ,CNDC,COGENT,-2.25854,53.3791,2020/5/24 11:50
717,United Kingdom,Manchester,"Joule House, 76 Trafford Wharf Road, Trafford Park",Equinix MA3 (fka. Telecity Manchester 3),M17 1ES,CNDC,COGENT,-2.29355,53.468,2020/5/24 11:50
718,United Kingdom,Manchester,Kilburn House Manchester Science Park,Equinix MA1 (fka. Telecity Kilburn House),M15 6SE,CNDC,COGENT,-2.23586,53.463,2020/5/24 11:50
719,United Kingdom,Manchester,"Reynolds House, Manchester Technopark, 4 Archway",Equinix MA2,M15 5RN,CNDC,COGENT,-2.24761,53.4643,2020/5/24 11:50
720,United Kingdom,Manchester,"Synergy House (Unit 1), Manchester Science Park",Equinix MA4 (fka. UK Grid),M15 6SY,CNDC,COGENT,-2.23703,53.462,2020/5/24 11:50
721,United Kingdom,Manchester,"Synergy House (Unit 6), Manchester Science Park",Equinix MA4 - Unit 6 (fka. UK Grid),M15 6SY,CNDC,COGENT,-2.23847,53.4616,2020/5/24 11:50
722,United Kingdom,Manchester,Unit 3 Williams House Manchester Science Park Lloy,Equinix MA1,M15 65E,CNDC,COGENT,-2.23661,53.4641,2020/5/24 11:50
723,United Kingdom,Manchester,"Unit 6 Waterside, Trafford Wharf Road",LDeX2,M17 1WD,CNDC,COGENT,-2.2935,53.4669,2020/5/24 11:50
724,United Kingdom,Manchester,"Unit 7, Waterside, Trafford Park",LDEX3,M17 1WD,CNDC,COGENT,-2.29424,53.466,2020/5/24 11:50
725,United Kingdom,Middlesex,"Western International Park, Hayes Road",VIRTUS DC LON2 (Hayes),UB2 5XJ,CNDC,COGENT,-0.402797,51.4963,2020/5/24 11:50
726,United Kingdom,Slough,1 Banbury Avenue,Equinix LD7,SL1 4LH,CNDC,COGENT,-0.633919,51.5259,2020/5/24 11:50
727,United Kingdom,Slough,"12 Liverpool Road, Slough Trading Estate",CyrusOne London I,SL1 4SS,CNDC,COGENT,-0.621629,51.5227,2020/5/24 11:50
728,United Kingdom,Slough,13 Liverpool Rd,Equinix LD10,SL1 4Qz,CNDC,COGENT,-0.623736,51.522,2020/5/24 11:50
729,United Kingdom,Slough,"2 Buckingham Avenue, Slough Trading Estate",Equinix LD4,SL1 4NB,CNDC,COGENT,-0.636091,51.5235,2020/5/24 11:50
730,United Kingdom,Slough,352 Buckingham Avenue,Equinix LD6,SL1 4PF,CNDC,COGENT,-0.63495,51.5245,2020/5/24 11:50
731,United Kingdom,Slough,630 Ajax Avenue,Cyxtera LHR1-A (fka LO1),SL14DG,CNDC,COGENT,-0.617882,51.5165,2020/5/24 11:50
732,United Kingdom,Slough,631 Ajax Avenue,Cyxtera LHR1-B (fka LO5),SL1 4DG,CNDC,COGENT,-0.618234,51.5172,2020/5/24 11:50
733,United Kingdom,Slough,724-729 Dundee Road,Iron Mountain LON-1,SL1 4JU,CNDC,COGENT,-0.631104,51.5273,2020/5/24 11:50
734,United Kingdom,Slough,772 Buckingham Avenue,Equinix LD4.2,SL1 4NL,CNDC,COGENT,-0.639013,51.5241,2020/5/24 11:50
735,United Kingdom,Slough,8 Buckingham Avenue,Equinix LD5,SL1 4AX,CNDC,COGENT,-0.62905,51.5225,2020/5/24 11:50
737,United Kingdom,Trafford Park Manchester,2 Adbaston Road Trafford Road,UKFast MaNOC 4,M32 0TP,CNDC,COGENT,-2.32678,53.4616,2020/5/24 11:50
738,United Kingdom,Trafford Park-Manchester,4 Blackmore Road,UKFast MaNOC 5,M32 OQY,CNDC,COGENT,0.290298,51.672,2020/5/24 11:50
739,United Kingdom,Trafford Park-Manchester,6 Blackmore Road,UKFast MaNOC 6,M32 0QY,CNDC,COGENT,0.27488,51.6626,2020/5/24 11:50
740,United Kingdom,Trafford Park-Manchester,8 Blackmore road,UKFast MaNOC 7,M32 0QY,CNDC,COGENT,0.290636,51.6721,2020/5/24 11:50
741,United Kingdom,Welwyn Garden City,20 Blackfan Road,COLT - London-North,AL7 1QA,CNDC,COGENT,-0.183762,51.7999,2020/5/24 11:50
742,United Kingdom,West Drayton,"Bath Road, Unit 1, Airport Gate, Harmondsworth, We",Digital Realty (fka. Equinix LD2),UB7 ONA,CNDC,COGENT,-0.460216,51.4815,2020/5/24 11:50
743,United Kingdom,Woking,"Unit 21 Goldsworth Trading Estate, Kestrel Way",Evoque LHR2,GU21 3BA,CNDC,COGENT,-0.587003,51.3235,2020/5/24 11:50
744,United Kingdom,Woking,"Unit 21 Goldsworth Trading Estate, Kestrel Way",Digital Realty Trust - Woking DC,GU21 3BA,CNDC,COGENT,-0.587003,51.3235,2020/5/24 11:50
745,United Kingdom,Wokingham,Eskdale Road Winnersh Triangle,Cyxtera LHR3-A (fka LO6),RG41 5TS,CNDC,COGENT,-0.881352,51.439,2020/5/24 11:50
746,United States,"Addison, TX",15301 Dallas Parkway N,Colonnade I,75001,COB,COGENT,-96.8226,32.9581,2020/5/24 11:50
747,United States,"Addison, TX",15303 Dallas Parkway,Colonnade II,75001,COB,COGENT,-96.8221,32.9588,2020/5/24 11:50
748,United States,"Addison, TX",15305 Dallas Pky N,Colonnade III,75001,COB,COGENT,-96.8232,32.9589,2020/5/24 11:50
749,United States,"Addison, TX",15455 North Dallas Parkway,Millennium Tower,75001,COB,COGENT,-96.822,32.96,2020/5/24 11:50
750,United States,"Addison, TX",5080 Spectrum Dr,Spectrum Center,75001,COB,COGENT,-96.8242,32.9551,2020/5/24 11:50
751,United States,"Albany, NY",11 N Pearl St,FirstLight (fka. Fibertech),12207,CNDC,COGENT,-73.752,42.65,2020/5/24 11:50
752,United States,"Albany, NY",80 State Street,FirstLight (fka. Fibertech),12207,CNDC,COGENT,-73.7526,42.6492,2020/5/24 11:50
753,United States,"Albany, NY",80 State Street,80 State Street,12207,COB,COGENT,-73.7526,42.6492,2020/5/24 11:50
754,United States,"Albuquerque, NM",123 Central Avenue NW,Big Byte Data Center,87102,CNDC,COGENT,-106.649,35.0845,2020/5/24 11:50
755,United States,"Albuquerque, NM",505 Marquette Ave NW,Compass Bank Building,87102,CNDC,COGENT,-106.652,35.0886,2020/5/24 11:50
756,United States,"Albuquerque, NM",505 Marquette Ave NW,Compass Bank Building,87102,CNDC,COGENT,-106.652,35.0886,2020/5/24 11:50
757,United States,"Albuquerque, NM",505 Marquette Ave NW,H5,87102,CNDC,COGENT,-106.652,35.0886,2020/5/24 11:50
758,United States,"Albuquerque, NM",725 6th St NW,OSO Grande Technologies,87102,CNDC,COGENT,-106.653,35.0925,2020/5/24 11:50
759,United States,"Allen, TX",2300 Chelsea Boulevard,Cyrus One Allen TX Data Center,75013,CNDC,COGENT,-96.6634,33.1372,2020/5/24 11:50
760,United States,"Allen, TX",820 Allen Commerce Parkway,Tierpoint DA2,75013,CNDC,COGENT,-96.6574,33.1386,2020/5/24 11:50
761,United States,"Alpharetta, GA",3200 Webb Bridge Rd,T5 Data Centers,30005,CNDC,COGENT,-84.2642,34.0758,2020/5/24 11:50
762,United States,"Anaheim, CA",2421 W. La Palma Avenue,Anaheim Palms Telecom Center - Bldg. 2,92801,CNDC,COGENT,-117.882,33.8475,2020/5/24 11:50
764,United States,"Anaheim, CA",300 S Harbor Blvd,Cogent Data Center - Anaheim,92805,CDC,COGENT,-117.918,33.8316,2020/5/24 11:50
765,United States,"Anaheim, CA",300 S Harbor Blvd,Bank of America Tower,92805,COB,COGENT,-117.918,33.8316,2020/5/24 11:50
766,United States,"Arlington, VA",1100 N Glebe Rd,Three Ballston Plaza,22201,COB,COGENT,-77.118,38.8834,2020/5/24 11:50
767,United States,"Arlington, VA",1110 N Glebe Rd,Two Ballston Plaza,22201,COB,COGENT,-77.1173,38.8838,2020/5/24 11:50
768,United States,"Arlington, VA",1300 N 17th St,Arlington Tower,22209,COB,COGENT,-77.0725,38.8936,2020/5/24 11:50
769,United States,"Arlington, VA",1550 Wilson Blvd,1550 Wilson Blvd,22209,COB,COGENT,-77.0759,38.8943,2020/5/24 11:50
770,United States,"Arlington, VA",1560 Wilson Blvd,1560 Wilson Blvd,22209,COB,COGENT,-77.0765,38.8939,2020/5/24 11:50
771,United States,"Arlington, VA",1600 Wilson Blvd,1600 Wilson Blvd,22209,COB,COGENT,-77.0772,38.894,2020/5/24 11:50
772,United States,"Arlington, VA",1616 N Fort Myer Dr,1616 N Fort Myer Dr,22209,COB,COGENT,-77.0726,38.893,2020/5/24 11:50
773,United States,"Arlington, VA",1700 N Moore St,Rosslyn Metro Center,22209,COB,COGENT,-77.0719,38.8957,2020/5/24 11:50
774,United States,"Arlington, VA",1901 N. Fort Myer Dr.,Rosslyn Gateway South,22209,COB,COGENT,-77.072,38.8974,2020/5/24 11:50
775,United States,"Arlington, VA",1911 N Fort Myer Dr,Rosslyn Gateway North,22209,COB,COGENT,-77.0722,38.8979,2020/5/24 11:50
776,United States,"Arlington, VA",4301 N Fairfax Dr,Ballston Station,22203,COB,COGENT,-77.1134,38.8828,2020/5/24 11:50
777,United States,"Ashburn, VA",21551 Beaumeade Circle,Equinix DC10,20147,CNDC,COGENT,-77.4513,39.0209,2020/5/24 11:50
778,United States,"Ashburn, VA",21635 Red Rum Drive,zColo (fka. Latisys),20147,CNDC,COGENT,-77.4816,39.016,2020/5/24 11:50
779,United States,"Ashburn, VA",21691 Filigree Court,Equinix DC4,20147,CNDC,COGENT,-77.4613,39.0167,2020/5/24 11:50
780,United States,"Ashburn, VA","21701 Filigree Court, Bldg D",Equinix DC5,20147,CNDC,COGENT,-77.4613,39.016,2020/5/24 11:50
781,United States,"Ashburn, VA",21711 Filigree Ct,Equinix DC1,20147,CNDC,COGENT,-77.459,39.0164,2020/5/24 11:50
782,United States,"Ashburn, VA",21715 Filigree Court Bldg. F,Equinix DC2,20147,CNDC,COGENT,-77.459,39.0164,2020/5/24 11:50
783,United States,"Ashburn, VA",21721 Filigree Court,Equinix DC6,20147,CNDC,COGENT,-77.4586,39.015,2020/5/24 11:50
784,United States,"Ashburn, VA",21721 Filigree Court B,Equinix DC11,20147,CNDC,COGENT,-77.4586,39.015,2020/5/24 11:50
785,United States,"Ashburn, VA",21741 Red Rum Drive,Sabey - Intergate.Ashburn Bldg. C,20147,CNDC,COGENT,-77.4783,39.0154,2020/5/24 11:50
786,United States,"Ashburn, VA",21800 Beaumeade Cir,H5 Ashburn,20147,CNDC,COGENT,-77.4517,39.0137,2020/5/24 11:50
787,United States,"Ashburn, VA",21890 Uunet Drive,Aligned Datacenter Ashburn,20147,CNDC,COGENT,-77.4718,39.0129,2020/5/24 11:50
788,United States,"Ashburn, VA",22271 Broderick Drive,QTS Ashburn,20166,CNDC,COGENT,-77.4497,38.9998,2020/5/24 11:50
789,United States,"Ashburn, VA",43790 Devin Shafron Drive,Digital Realty - Bldg. E,20147,CNDC,COGENT,-77.4858,39.0049,2020/5/24 11:50
790,United States,"Ashburn, VA",43791 Devin Shafron Drive,Digital Realty - Bldg. D,20147,CNDC,COGENT,-77.487,39.0033,2020/5/24 11:50
791,United States,"Ashburn, VA",43830 Devin Shafron,Digital Realty - Bldg. F,20147,CNDC,COGENT,-77.4844,39.0041,2020/5/24 11:50
792,United States,"Ashburn, VA",43831 Devin Shafron Drive,Digital Realty - Bldg. C,20147,CNDC,COGENT,-77.4854,39.0027,2020/5/24 11:50
793,United States,"Ashburn, VA",43881 Devin Shafron Dr,Digital Realty - Bldg. B,20147,CNDC,COGENT,-77.4833,39.0021,2020/5/24 11:50
794,United States,"Ashburn, VA",43915 Devin Shafron Drive,Digital Realty - Bldg. A,20147,CNDC,COGENT,-77.482,39.0028,2020/5/24 11:50
795,United States,"Ashburn, VA",43940 Digital Loudoun Plaza,Digital Realty - Bldg. G,20147,CNDC,COGENT,-77.4792,39.0008,2020/5/24 11:50
796,United States,"Ashburn, VA",44060 Digital Loudoun Plaza,Digital Realty,20147,CNDC,COGENT,-77.4794,39.0015,2020/5/24 11:50
797,United States,"Ashburn, VA",44100 Digital Loudoun Plaza,Digital Realty,20147,CNDC,COGENT,-77.4785,39.0012,2020/5/24 11:50
798,United States,"Ashburn, VA",44245 Gigabit Plaza,RagingWire VA3,20147,CNDC,COGENT,-77.4707,39.0193,2020/5/24 11:50
800,United States,"Ashburn, VA",44461 Chilum Place,Digital Realty (fka. DuPont Fabros ACC6),20147,CNDC,COGENT,-77.4636,39.0211,2020/5/24 11:50
801,United States,"Ashburn, VA","44470 Chilum Place, Bldg 1",Equinix DC3,20147,CNDC,COGENT,-77.4609,39.022,2020/5/24 11:50
802,United States,"Ashburn, VA",44480 Hastings Drive,Digital Realty (fka. DuPont Fabros ACC4),20147,CNDC,COGENT,-77.4625,39.0178,2020/5/24 11:50
803,United States,"Ashburn, VA",44521 Hastings Drive,Digital Realty (fka. DuPont Fabros ACC5),20147,CNDC,COGENT,-77.4632,39.0201,2020/5/24 11:50
804,United States,"Ashburn, VA",44610 Guilford Dr,RagingWire VA2,20147,CNDC,COGENT,-77.4573,39.0235,2020/5/24 11:50
805,United States,"Ashburn, VA",44664 Guilford Dr.,RagingWire VA1,20147,CNDC,COGENT,-77.455,39.0226,2020/5/24 11:50
806,United States,"Ashburn, VA",44790 Performance Circle,Equinix DC12,20147,CNDC,COGENT,-77.4505,39.0222,2020/5/24 11:50
807,United States,"Atlanta, GA",1 Concourse Pky NE,Concourse Corporate Center One,30328,COB,COGENT,-84.3536,33.9162,2020/5/24 11:50
808,United States,"Atlanta, GA",1 Ravinia Dr,One Ravinia Drive,30346,COB,COGENT,-84.3372,33.9218,2020/5/24 11:50
809,United States,"Atlanta, GA",100 Galleria Pky SE,100 Galleria Pky SE,30339,COB,COGENT,-84.4627,33.8849,2020/5/24 11:50
810,United States,"Atlanta, GA",100 Peachtree St NW,Equitable Bldg,30303,COB,COGENT,-84.3884,33.7568,2020/5/24 11:50
811,United States,"Atlanta, GA",1000 Abernathy Rd NE,Bldg 400,30328,COB,COGENT,-84.3545,33.9344,2020/5/24 11:50
812,United States,"Atlanta, GA",1003 Donnelly Avenue SW,Edgeconnex ADCATL01,30310,CNDC,COGENT,-84.4201,33.7294,2020/5/24 11:50
813,United States,"Atlanta, GA",101 Marietta St NW,Centennial Tower,30303,COB,COGENT,-84.3921,33.7568,2020/5/24 11:50
814,United States,"Atlanta, GA",1033 Jefferson St,QTS Datacenter,30319,CNDC,COGENT,-84.4209,33.7777,2020/5/24 11:50
815,United States,"Atlanta, GA",1040 Crown Pointe Pky,Tower 1040,30338,COB,COGENT,-84.3466,33.9314,2020/5/24 11:50
816,United States,"Atlanta, GA",1050 Crown Pointe Pky,Crown Pointe,30338,COB,COGENT,-84.3454,33.9317,2020/5/24 11:50
817,United States,"Atlanta, GA",1075 Peachtree St. NE,12th & Midtown,30309,COB,COGENT,-84.3826,33.7838,2020/5/24 11:50
818,United States,"Atlanta, GA",1100 Abernathy Rd NE,Bldg 500,30328,COB,COGENT,-81.367,31.2031,2020/5/24 11:50
819,United States,"Atlanta, GA",1100 Peachtree St NE,Eleven Hundred Peachtree Street,30309,COB,COGENT,-84.3839,33.785,2020/5/24 11:50
821,United States,"Atlanta, GA",1117 Perimeter Ctr W,1117 Perimeter Ctr W,30338,COB,COGENT,-84.3526,33.9307,2020/5/24 11:50
822,United States,"Atlanta, GA",115 Perimeter Center Place (aka 115 Perimeter Cent,115 Perimeter Center,30346,COB,COGENT,-84.3404,33.9281,2020/5/24 11:50
823,United States,"Atlanta, GA",1170 Peachtree Street NE,The Proscenium,30309,COB,COGENT,-84.3837,33.7864,2020/5/24 11:50
824,United States,"Atlanta, GA",1175 Peachtree St NE,100 Colony Square,30361,COB,COGENT,-84.3827,33.7869,2020/5/24 11:50
825,United States,"Atlanta, GA",1180 Peachtree NE,1180 Peachtree NE,30309,COB,COGENT,-84.3841,33.7869,2020/5/24 11:50
826,United States,"Atlanta, GA",1180 West Peachtree Street NW,Regions Plaza,30309,COB,COGENT,-84.3882,33.787,2020/5/24 11:50
827,United States,"Atlanta, GA",1200 Abernathy Rd NE,Bldg 600,30328,COB,COGENT,-84.3522,33.9343,2020/5/24 11:50
828,United States,"Atlanta, GA",1200 Ashwood Pky,1200 Ashwood Pky,30338,COB,COGENT,-84.3406,33.9336,2020/5/24 11:50
829,United States,"Atlanta, GA",1201 Peachtree St NE,400 Colony Square,30361,COB,COGENT,-84.3831,33.7876,2020/5/24 11:50
830,United States,"Atlanta, GA",1201 W Peachtree St NW,One Atlantic Center,30309,COB,COGENT,-84.9149,33.3853,2020/5/24 11:50
831,United States,"Atlanta, GA",1230 Peachtree St NE,Promenade II,30309,COB,COGENT,-84.3852,33.7879,2020/5/24 11:50
832,United States,"Atlanta, GA",133 Peachtree St NE,Georgia Pacific Center,30303,COB,COGENT,-84.387,33.7572,2020/5/24 11:50
833,United States,"Atlanta, GA",1349 W Peachtree Street NE,Two Midtown Plaza,30309,COB,COGENT,-84.3865,33.7915,2020/5/24 11:50
834,United States,"Atlanta, GA",1355 Peachtree St NE,The Peachtree Bldg,30309,COB,COGENT,-84.385,33.7918,2020/5/24 11:50
835,United States,"Atlanta, GA",1360 Peachtree St NE,One Midtown Plaza,30309,COB,COGENT,-84.3882,33.792,2020/5/24 11:50
836,United States,"Atlanta, GA",1375 Peachtree Street NE,Pershing Pointe,30309,COB,COGENT,-84.3856,33.7924,2020/5/24 11:50
837,United States,"Atlanta, GA",180 Peachtree St NW,Equinix AT1 (fka. S&D),30303,CNDC,COGENT,-84.3878,33.7586,2020/5/24 11:50
838,United States,"Atlanta, GA",191 Peachtree St NE,One Ninety One Peachtree Tower,30303,COB,COGENT,-84.3867,33.759,2020/5/24 11:50
839,United States,"Atlanta, GA",2 Concourse Pky NE,Concourse Corporate Center Two,30328,COB,COGENT,-84.353,33.9156,2020/5/24 11:50
840,United States,"Atlanta, GA",2 Ravinia Dr,Ravinia Two,30346,COB,COGENT,-81.54,31.1714,2020/5/24 11:50
841,United States,"Atlanta, GA",200 Galleria Pky SE,200 Galleria Pky SE,30339,COB,COGENT,-84.4636,33.8859,2020/5/24 11:50
842,United States,"Atlanta, GA",201 17th Street NW,Atlantic Station,30363,COB,COGENT,-77.0405,38.8929,2020/5/24 11:50
843,United States,"Atlanta, GA",225 Peachtree St NE,South Tower,30303,COB,COGENT,-84.4536,33.2823,2020/5/24 11:50
844,United States,"Atlanta, GA",229 Peachtree St NE,International Tower,30303,COB,COGENT,-84.3864,33.7599,2020/5/24 11:50
845,United States,"Atlanta, GA",230 Peachtree St NW,230 Peachtree St NW,30303,COB,COGENT,-84.3879,33.7601,2020/5/24 11:50
846,United States,"Atlanta, GA",2300 Windy Ridge Pky SE,Wildwood Center,30339,COB,COGENT,-84.4602,33.906,2020/5/24 11:50
847,United States,"Atlanta, GA",233 Peachtree St NE,Harris Tower,30303,COB,COGENT,-84.3872,33.7612,2020/5/24 11:50
848,United States,"Atlanta, GA",235 Peachtree St NE,North Tower,30303,COB,COGENT,-84.3872,33.7605,2020/5/24 11:50
849,United States,"Atlanta, GA",245 Peachtree Center Ave NE,Marquis One,30303,COB,COGENT,-84.3854,33.7613,2020/5/24 11:50
850,United States,"Atlanta, GA",250 Williams St NW,Inforum,30303,COB,COGENT,-84.3914,33.7615,2020/5/24 11:50
851,United States,"Atlanta, GA",250 Williams St NW E100,Internap,30303,CNDC,COGENT,-84.3914,33.7615,2020/5/24 11:50
852,United States,"Atlanta, GA",2839 Paces Ferry Rd SE,Overlook II,30339,COB,COGENT,-84.472,33.8675,2020/5/24 11:50
853,United States,"Atlanta, GA",285 Peachtree Center Ave NE,Marquis Two,30303,COB,COGENT,-84.3857,33.762,2020/5/24 11:50
854,United States,"Atlanta, GA",2859 Paces Ferry Rd SE,Overlook III,30339,COB,COGENT,-84.4694,33.8689,2020/5/24 11:50
855,United States,"Atlanta, GA",3 Ravinia Dr,Three Ravinia Drive,30346,COB,COGENT,-84.3347,33.9217,2020/5/24 11:50
856,United States,"Atlanta, GA",300 Galleria Pky SE,300 Galleria Pky SE,30339,COB,COGENT,-84.4633,33.8832,2020/5/24 11:50
857,United States,"Atlanta, GA",303 Peachtree Center Ave NE,Suntrust Plaza Garden Offices,30308,COB,COGENT,-84.3865,33.7628,2020/5/24 11:50
858,United States,"Atlanta, GA",303 Peachtree St NE,Suntrust Plaza,30308,COB,COGENT,-84.3865,33.7628,2020/5/24 11:50
859,United States,"Atlanta, GA",3060 Peachtree Rd NW,One Buckhead Plaza,30305,COB,COGENT,-84.3817,33.8395,2020/5/24 11:50
860,United States,"Atlanta, GA",3280 Peachtree Road NE,Terminus (100) the Place,30305,COB,COGENT,-84.3713,33.8445,2020/5/24 11:50
861,United States,"Atlanta, GA",3333 Peachtree Rd NE,Atlanta Financial Center,30326,COB,COGENT,-84.3676,33.8463,2020/5/24 11:50
862,United States,"Atlanta, GA",3333 Piedmont Rd NE,Terminus 200,30305,COB,COGENT,-84.372,33.846,2020/5/24 11:50
863,United States,"Atlanta, GA",3340 Peachtree Rd NE,Tower Place 100,30326,COB,COGENT,-84.371,33.8474,2020/5/24 11:50
864,United States,"Atlanta, GA",3343 Peachtree Rd NE,Atlanta Finanical Center - East Tower,30326,COB,COGENT,-84.3667,33.847,2020/5/24 11:50
865,United States,"Atlanta, GA",3350 Peachtree Rd NE,One Capital City Plaza,30326,COB,COGENT,-84.3685,33.8476,2020/5/24 11:50
866,United States,"Atlanta, GA",3350 Riverwood Parkway,Riverwood 100,30339,COB,COGENT,-81.3724,31.2653,2020/5/24 11:50
867,United States,"Atlanta, GA",3353 Peachtree Rd NE,Atlanta Financial Center - North Tower,30326,COB,COGENT,-84.3669,33.8474,2020/5/24 11:50
868,United States,"Atlanta, GA",3399 Peachtree Rd NE,Buckhead Tower,30326,COB,COGENT,-84.3596,33.8461,2020/5/24 11:50
869,United States,"Atlanta, GA",34 Peachtree St NW,Total Server,30303,CNDC,COGENT,-84.3896,33.755,2020/5/24 11:50
870,United States,"Atlanta, GA",34 Peachtree St NW,34 Peachtree St NW,30303,COB,COGENT,-84.3896,33.755,2020/5/24 11:50
871,United States,"Atlanta, GA",3414 Peachtree Rd NE,Monarch Plaza at Monarch Centre,30326,COB,COGENT,-84.3641,33.8504,2020/5/24 11:50
872,United States,"Atlanta, GA",3424 Peachtree Rd NE,Monarch Tower at Monarch Centre,30326,COB,COGENT,-84.3645,33.8517,2020/5/24 11:50
873,United States,"Atlanta, GA",3438 Peachtree Rd,Phipps Tower,30326,COB,COGENT,-84.364,33.8528,2020/5/24 11:50
874,United States,"Atlanta, GA",3445 Peachtree Rd NE,Two Live Oak Center,30326,COB,COGENT,-84.3613,33.85,2020/5/24 11:50
875,United States,"Atlanta, GA",345 Courtland St NE - 1st Floor,CenturyLink (fka. Level 3),30308,CNDC,COGENT,-84.3837,33.7652,2020/5/24 11:50
876,United States,"Atlanta, GA",3455 Peachtree Rd NE,The Pinnacle,30326,COB,COGENT,-84.3602,33.8502,2020/5/24 11:50
877,United States,"Atlanta, GA",3475 Piedmont Rd,Prominence Bldg,30305,COB,COGENT,-84.3751,33.8506,2020/5/24 11:50
878,United States,"Atlanta, GA",3490 Piedmont Rd NE,One Securities Centre,30305,COB,COGENT,-84.3761,33.8488,2020/5/24 11:50
879,United States,"Atlanta, GA",3500 Lenox Road NE,One Alliance Center,30326,COB,COGENT,-84.3663,33.852,2020/5/24 11:50
880,United States,"Atlanta, GA",3500 Piedmont Road,3500 Securities Centre,30305,COB,COGENT,-84.3771,33.849,2020/5/24 11:50
881,United States,"Atlanta, GA","3525 Piedmont Road, NE",Piedmont Center Building 5,30305,COB,COGENT,-84.3767,33.8517,2020/5/24 11:50
882,United States,"Atlanta, GA","3525 Piedmont Road, NE",Piedmont Center Building 8,30305,COB,COGENT,-84.3767,33.8517,2020/5/24 11:50
883,United States,"Atlanta, GA","3525 Piedmont Road, NE",Piedmont Center Building 6,30305,COB,COGENT,-84.3767,33.8517,2020/5/24 11:50
884,United States,"Atlanta, GA","3525 Piedmont Road, NE",Piedmont Cetner Building 7,30305,COB,COGENT,-84.3767,33.8517,2020/5/24 11:50
885,United States,"Atlanta, GA",3535 Piedmont Rd,Piedmont 14,30305,COB,COGENT,-84.3771,33.8529,2020/5/24 11:50
886,United States,"Atlanta, GA",3550 Lenox Road,Three Alliance Center,30326,COB,COGENT,-84.3676,33.8517,2020/5/24 11:50
887,United States,"Atlanta, GA",3565 Piedmont Road (Bldg 1),Building 1,30305,COB,COGENT,-84.3786,33.8522,2020/5/24 11:50
888,United States,"Atlanta, GA",3565 Piedmont Road (Bldg 2),Building 2,30305,COB,COGENT,-84.3786,33.8522,2020/5/24 11:50
889,United States,"Atlanta, GA",3565 Piedmont Road (Bldg 3),Building 3,30305,COB,COGENT,-84.3786,33.8522,2020/5/24 11:50
890,United States,"Atlanta, GA",3565 Piedmont Road (Bldg 4),Building 4,30305,COB,COGENT,-84.3786,33.8522,2020/5/24 11:50
891,United States,"Atlanta, GA",3575 Piedmont Rd NE,Piedmont 15,30305,COB,COGENT,-84.3791,33.8541,2020/5/24 11:50
892,United States,"Atlanta, GA",400 Galleria Pky SE,400 Galleria Pky SE,30339,COB,COGENT,-84.4613,33.8866,2020/5/24 11:50
893,United States,"Atlanta, GA",400 Perimeter Center Terrace,400 Perimeter Center Ter NE,30346,COB,COGENT,-81.3694,31.1543,2020/5/24 11:50
894,United States,"Atlanta, GA",450 Interstate Parkway SE,Equinix AT4 (Formerly Verizon),30339,CNDC,COGENT,-84.4536,33.8954,2020/5/24 11:50
895,United States,"Atlanta, GA",470 E Paces Ferry Rd NE,470 Exchange,30305,CNDC,COGENT,-84.3724,33.8389,2020/5/24 11:50
896,United States,"Atlanta, GA",4905 North Point Parkway,Ascent ATL1,30009,CNDC,COGENT,-84.2723,34.0567,2020/5/24 11:50
897,United States,"Atlanta, GA",5 Concourse Pky NE,Concourse Corporate Center V (Queen Building),30328,COB,COGENT,-84.3545,33.9171,2020/5/24 11:50
898,United States,"Atlanta, GA",55 Marietta St NW,Utility Computing,30303,,COGENT,-84.3911,33.7559,2020/5/24 11:50
899,United States,"Atlanta, GA",55 Marietta St NW,Cogent Data Center - Atlanta,30303,CDC,COGENT,-84.3911,33.7559,2020/5/24 11:50
900,United States,"Atlanta, GA",55 Marietta St NW,Building MMR,30303,CNDC,COGENT,-84.3911,33.7559,2020/5/24 11:50
901,United States,"Atlanta, GA",55 Marietta St NW,55 Marietta St NW,30303,COB,COGENT,-84.3911,33.7559,2020/5/24 11:50
902,United States,"Atlanta, GA",55 Marietta St NW 8th Floor,Colo Atl (JT Communications),30303,CNDC,COGENT,-84.3912,33.7559,2020/5/24 11:50
903,United States,"Atlanta, GA",55 Park Pl NE,55 Park Place Phase I,30303,COB,COGENT,-78.9395,43.0374,2020/5/24 11:50
904,United States,"Atlanta, GA",56 Marietta St NW,Equinix AT2 (fka. PAIX),30303,CNDC,COGENT,-84.3914,33.7555,2020/5/24 11:50
905,United States,"Atlanta, GA",56 Marietta St NW,Equinix AT3 (fka. S&D),30303,CNDC,COGENT,-84.3914,33.7555,2020/5/24 11:50
906,United States,"Atlanta, GA",56 Marietta St NW (2nd Floor),Digital Realty ATL1 (fka. Telx MMR),30303,CNDC,COGENT,-84.3914,33.7555,2020/5/24 11:50
907,United States,"Atlanta, GA",5901A Peachtree Dunwoody Rd NE,Palisades Building A,30328,COB,COGENT,-84.3496,33.9165,2020/5/24 11:50
908,United States,"Atlanta, GA",5901B Peachtree Dunwoody Rd NE,Palisades Building B,30328,COB,COGENT,-84.3514,33.9151,2020/5/24 11:50
909,United States,"Atlanta, GA",5901C Peachtree Dunwoody Rd NE,Palisades Building C,30328,COB,COGENT,-84.3495,33.9165,2020/5/24 11:50
910,United States,"Atlanta, GA",5909 Peachtree Dunwoody Rd NE,Palisades Building D,30328,COB,COGENT,-84.3495,33.9171,2020/5/24 11:50
911,United States,"Atlanta, GA",6 Concourse Pky NE,Concourse Corporate Center VI (King Building),30328,COB,COGENT,-84.3559,33.9171,2020/5/24 11:50
912,United States,"Atlanta, GA",6 West Druid Hills Drive,DC BLOX,30329,CNDC,COGENT,-84.3392,33.8317,2020/5/24 11:50
913,United States,"Atlanta, GA",600 Peachtree St NE,Bank of America Plaza,30308,COB,COGENT,-84.3861,33.771,2020/5/24 11:50
914,United States,"Atlanta, GA",7000 Central Pky NE,7000 Central Pky NE,30328,COB,COGENT,-84.3484,33.9276,2020/5/24 11:50
915,United States,"Atlanta, GA",730 Peachtree St NE,730 Midtown,30308,COB,COGENT,-84.386,33.7745,2020/5/24 11:50
916,United States,"Atlanta, GA",756 W Peachtree Street NW,CODA,30308,COB,COGENT,-84.3876,33.7753,2020/5/24 11:50
917,United States,"Atlanta, GA",760 West Peachtree Street NW,Data Bank ATL1,20208,CNDC,COGENT,-84.3877,33.7755,2020/5/24 11:50
918,United States,"Atlanta, GA",900 Circle 75 Pky SE,Bldg 900,30339,COB,COGENT,-84.4678,33.8875,2020/5/24 11:50
919,United States,"Atlanta, GA",945 E Paces Ferry Rd NE,Resurgens Plaza,30326,COB,COGENT,-81.484,31.1246,2020/5/24 11:50
920,United States,"Atlanta, GA",950 E Paces Ferry Rd NE,Atlanta Plaza One,30326,COB,COGENT,-81.3866,31.14,2020/5/24 11:50
921,United States,"Atlanta, GA",999 Peachtree St NE,First Union Plaza,30309,COB,COGENT,-84.3837,33.7811,2020/5/24 11:50
922,United States,"AURORA, CO",11900 E CORNELL AVE,Flexential Aurora (fka. Peak 10 / ViaWest),80014,CNDC,COGENT,-104.85,39.6611,2020/5/24 11:50
923,United States,"Aurora, IL",2905 Diehl Road,CyrusOne Aurora I (fka. CME),60502,CNDC,COGENT,-88.2437,41.7976,2020/5/24 11:50
924,United States,"Aurora, CO",3431 N Windsor Dr,SunGard den-3431,80011,CNDC,COGENT,-104.762,39.7627,2020/5/24 11:50
925,United States,"Austin, TX",111 Congress Ave,111 Congress Ave,78701,COB,COGENT,-88.4514,41.6655,2020/5/24 11:50
926,United States,"Austin, TX","1122 Colorado St, Suite 110",Bestline Austin Data Center,78701,CNDC,COGENT,-97.7429,30.2735,2020/5/24 11:50
927,United States,"Austin, TX",205 W. 9th Street,Flexential (fka. Peak 10 / ViaWest),78701,CNDC,COGENT,-97.7436,30.2713,2020/5/24 11:50
928,United States,"Austin, TX",300 W 6th Street,300 West 6th,78701,COB,COGENT,-97.7458,30.2692,2020/5/24 11:50
929,United States,"Austin, TX",301 Congress,301 Congress,78701,COB,COGENT,-97.7435,30.2653,2020/5/24 11:50
930,United States,"Austin, TX",303 Colorado St,Colorado Tower,78701,COB,COGENT,-97.7449,30.2657,2020/5/24 11:50
931,United States,"Austin, TX",401 Congress Avenue,Frost Bank Tower,78701,COB,COGENT,-97.7428,30.2667,2020/5/24 11:50
932,United States,"Austin, TX",4100 Smith School Rd,Data Foundry Texas 1,78744,CNDC,COGENT,-97.7134,30.1992,2020/5/24 11:50
933,United States,"Austin, TX",4100 Smith School Rd,Data Foundry Texas 2,78744,CNDC,COGENT,-97.7134,30.1992,2020/5/24 11:50
934,United States,"Austin, TX",600 Congress Avenue,One American Center,78701,COB,COGENT,-97.7433,30.2689,2020/5/24 11:50
935,United States,"Austin, TX",701 Brazos Street,Austin Centre,78701,COB,COGENT,-97.7406,30.2689,2020/5/24 11:50
936,United States,"Austin, TX",7100 Metropolis Dr,CyrusOne Austin III (fka. Met Center),78744,CNDC,COGENT,-97.7011,30.2068,2020/5/24 11:50
937,United States,"Austin, TX",7301 Metropolis Dr,CyrusOne Austin II (fka. Met Center),78744,CNDC,COGENT,-97.6979,30.2059,2020/5/24 11:50
938,United States,"Austin, TX","7401 Ben White Blvd East, Bldg 1",Data Foundry Austin 1,78741,CNDC,COGENT,-97.695,30.2151,2020/5/24 11:50
939,United States,"Austin, TX",8025 I.H. 35 North,vXchnge,78753,CNDC,COGENT,-97.6966,30.3428,2020/5/24 11:50
940,United States,"Austin, TX",816 Congress Ave,816 Congress,78701,COB,COGENT,-97.742,30.2708,2020/5/24 11:50
941,United States,"Austin, TX",98 San Jacinto Boulevard,San Jacinto Center,78701,COB,COGENT,-97.7418,30.2615,2020/5/24 11:50
942,United States,"Austin,, TX",100 Congress Ave,100 Congress Ave,78701,COB,COGENT,-97.7449,30.264,2020/5/24 11:50
943,United States,"Baltimore, MD",1 North Charles Street,1 North Charles Street,21201,COB,COGENT,-76.6148,39.2901,2020/5/24 11:50
944,United States,"Baltimore, MD",100 East Pratt St,100 East Pratt St,21202,COB,COGENT,-76.6128,39.2872,2020/5/24 11:50
945,United States,"Baltimore, MD",100 Light St,Transamerica Tower,21202,COB,COGENT,-76.6145,39.2874,2020/5/24 11:50
946,United States,"Baltimore, MD",100 S Charles Street,Bank of America Tower I,21201,COB,COGENT,-76.6154,39.287,2020/5/24 11:50
947,United States,"Baltimore, MD",100 S Charles Street,Tower ll,21201,COB,COGENT,-76.6154,39.287,2020/5/24 11:50
948,United States,"Baltimore, MD",111 Market Place,111 Market Place,21202,COB,COGENT,-76.6062,39.287,2020/5/24 11:50
949,United States,"Baltimore, MD",111 S Calvert St,Harborplace Tower,21202,COB,COGENT,-76.6115,39.2874,2020/5/24 11:50
950,United States,"Baltimore, MD",1401 Russell Street,Tierpoint BAL (Baltimore Technology Park),21230,CNDC,COGENT,-76.6255,39.2754,2020/5/24 11:50
951,United States,"Baltimore, MD",201 N Charles St,201 N Charles St,21201,COB,COGENT,-76.6145,39.2913,2020/5/24 11:50
952,United States,"Baltimore, MD",250 W Pratt St,250 W Pratt St,21201,COB,COGENT,-76.6188,39.2867,2020/5/24 11:50
953,United States,"Baltimore, MD",300 East Lombard St,300 East Lombard St,21202,COB,COGENT,-76.6105,39.2883,2020/5/24 11:50
955,United States,"Baltimore, MD",36 S Charles St,Charles Center South,21201,COB,COGENT,-76.6155,39.2879,2020/5/24 11:50
956,United States,"Baltimore, MD",401 East Pratt St,401 East Pratt St,21202,COB,COGENT,-76.6215,39.2857,2020/5/24 11:50
957,United States,"Baltimore, MD",500 East Pratt St,500 East Pratt St,21202,COB,COGENT,-76.6221,39.2863,2020/5/24 11:50
958,United States,"Baltimore, MD",7 St Paul St,Wachovia Tower,21202,COB,COGENT,-76.6137,39.2901,2020/5/24 11:50
959,United States,"Bellevue, NE",1001 North Fort Crook Road,Tierpoint - BEL,68005,CNDC,COGENT,-95.9248,41.1782,2020/5/24 11:50
960,United States,"Bellevue, WA",10885 NE 4th St,Summit ll,98004,COB,COGENT,-122.195,47.6135,2020/5/24 11:50
961,United States,"Bellevue, WA",10900 NE 4th St,Skyline Tower,98004,COB,COGENT,-122.195,47.614,2020/5/24 11:50
962,United States,"Bellevue, WA",10900 NE 8th St,Plaza Center,98004,COB,COGENT,-122.194,47.6179,2020/5/24 11:50
963,United States,"Bellevue, WA",355 110th Avenue NE,Summit l,98004,COB,COGENT,-98.1987,47.465,2020/5/24 11:50
964,United States,"Bellevue, WA",411 108th Ave NE,One Bellevue Center,98004,COB,COGENT,-122.197,47.6141,2020/5/24 11:50
965,United States,"Bellevue, WA",500 108th Ave NE,City Center Bellevue,98004,COB,COGENT,-122.196,47.615,2020/5/24 11:50
966,United States,"Bellevue, WA",600 108th Ave NE,Bellevue Corporate Plaza,98004,COB,COGENT,-122.196,47.616,2020/5/24 11:50
967,United States,"Bellevue, WA",601 108th Ave NE,Key Center,98004,COB,COGENT,-122.197,47.6161,2020/5/24 11:50
968,United States,"Bellevue, WA",777 108th Ave NE,Symetra Financial Center,98004,COB,COGENT,-122.197,47.6166,2020/5/24 11:50
969,United States,"Beltsville, MD",11700 Montgomery Rd,AiNET,20705,CNDC,COGENT,-76.9259,39.0528,2020/5/24 11:50
970,United States,"Bethesda, MD",3 Bethesda Metro Center,3 Bethesda Metro Center,20814,COB,COGENT,-77.0962,38.9843,2020/5/24 11:50
973,United States,"Bethesda, MD",6905 Rockledge Dr,Three Democracy,20817,COB,COGENT,-77.1399,39.0273,2020/5/24 11:50
974,United States,"Bethesda, MD",7501 Wisconsin Ave.,7501 Wisconsin Avenue,20814,COB,COGENT,-77.0939,38.9853,2020/5/24 11:50
976,United States,"Bethlehem, PA",3949 Schelden Circle,Tierpoint-LVQ (Lehigh Valley),18017,CNDC,COGENT,-75.3781,40.6749,2020/5/24 11:50
977,United States,"Billerica, MA",41 Alexander Road,Equinix BO2 (Formerly Verizon),1822,CNDC,COGENT,-71.2163,42.5524,2020/5/24 11:50
978,United States,"Birmingham, AL",1901 6th Avenue N,1901 6th Avenue North,35203,COB,COGENT,-81.4129,31.1438,2020/5/24 11:50
979,United States,"Birmingham, AL",2 20th Street N,Two North Twentieth,35203,COB,COGENT,-86.8058,33.5134,2020/5/24 11:50
980,United States,"Birmingham, AL",2001 Park Place North,Park Place Tower,35203,COB,COGENT,-86.8097,33.5192,2020/5/24 11:50
981,United States,"Birmingham, AL",420 20th Street N,Wachovia Tower,35203,COB,COGENT,-86.8084,33.5176,2020/5/24 11:50
982,United States,"Birmingham, AL",505 20th Street N,Financial Center,35203,COB,COGENT,-86.8078,33.5186,2020/5/24 11:50
983,United States,"Birmingham, AL",600 18th Street N,APCO/Southern Telecom Data Center,35203,CNDC,COGENT,-86.8129,33.518,2020/5/24 11:50
984,United States,"Bloomington, MN",5600 West 83rd Street,8200 Tower,55437,COB,COGENT,-93.3553,44.8545,2020/5/24 11:50
985,United States,"Bloomington, MN",5600 West 84th Street,Normandale Lake Office Park,55437,COB,COGENT,-73.976,40.7864,2020/5/24 11:50
986,United States,"Bloomington, MN",8300 Norman Center Dr,Normandale Lake Office Park,55437,COB,COGENT,-93.353,44.8527,2020/5/24 11:50
987,United States,"Bloomington, MN",8331 Norman Center Drive,8000 Tower,55437,COB,COGENT,-93.3513,44.8527,2020/5/24 11:50
988,United States,"Bloomington, MN",8500 Normandale Lake Blvd,Normandale Lake Office Park,55437,COB,COGENT,-93.3567,44.8535,2020/5/24 11:50
989,United States,"Bluffdale, UT",14926 Pony Express Drive,DataBank (fka. C7 - GP2),84065,CNDC,COGENT,-111.906,40.4799,2020/5/24 11:50
990,United States,"Bluffdale, UT",14944 Pony Express,DataBank (fka. C7 - GP1),84065,CNDC,COGENT,-111.905,40.48,2020/5/24 11:50
992,United States,"Boca Raton, FL",4680 Conference Way South,Equinix MI3,33431,CNDC,COGENT,-80.1093,26.3889,2020/5/24 11:50
993,United States,"Boca Raton, FL",4680 Conference Way South,4680 Conference Way South,33431,COB,COGENT,-80.1093,26.3889,2020/5/24 11:50
994,United States,"Boca Raton, FL",4700 Exchange Court,Boca Innovation Campus-4700 Address,33431-4483,COB,COGENT,-80.1067,26.3892,2020/5/24 11:50
995,United States,"Boca Raton, FL",4800 T Rex Building,Boca Innovation Campus-4800 T Rex,33431-4483,COB,COGENT,-80.1077,26.3899,2020/5/24 11:50
996,United States,"Boca Raton, FL",5000 T Rex Ave,Boca Raton Innovation Campus,33431-4483,COB,COGENT,-80.1082,26.3913,2020/5/24 11:50
997,United States,"Boca Raton, FL",5050 Conference Way North,Cogent Data Center - Boca Raton,33431,CDC,COGENT,-80.1097,26.3928,2020/5/24 11:50
998,United States,"Boca Raton, FL",5050 Conference Way North,T-Rex Corporate Center,33431,COB,COGENT,-80.1097,26.3928,2020/5/24 11:50
999,United States,"Boise, ID",10215 Emerald St,Fiberpipe,83704,CNDC,COGENT,-116.31,43.608,2020/5/24 11:50
1000,United States,"Boise, ID",2653 South Victory View Way,Involta Data Center,83709,CNDC,COGENT,-116.291,43.5797,2020/5/24 11:50
1002,United States,"Boston, MA",1 Beacon St,1 Beacon St,2108,COB,COGENT,-71.0608,42.3585,2020/5/24 11:50
1003,United States,"Boston, MA",1 Boston Pl (aka 201 Washington St),Boston Company Bldg,2108,COB,COGENT,-71.0583,42.3586,2020/5/24 11:50
1004,United States,"Boston, MA",1 Center Plz,One Center Plaza,2108,COB,COGENT,-71.0601,42.359,2020/5/24 11:50
1005,United States,"Boston, MA",1 Federal St,1 Federal St,2110,COB,COGENT,-71.0568,42.356,2020/5/24 11:50
1006,United States,"Boston, MA",1 Financial Ctr,One Financial Center,2111,COB,COGENT,-71.0564,42.3522,2020/5/24 11:50
1007,United States,"Boston, MA",1 Post Office Sq,1 Post Office Sq,2109,COB,COGENT,-71.0543,42.3566,2020/5/24 11:50
1008,United States,"Boston, MA",1 Summer St,5th Floor BMMR,2110,CNDC,COGENT,-71.0603,42.3547,2020/5/24 11:50
1009,United States,"Boston, MA",1 Summer St,1 Summer St,2110,COB,COGENT,-71.0603,42.3547,2020/5/24 11:50
1011,United States,"Boston, MA",1 Washington Mall,One Washington Mall,2108,COB,COGENT,-71.3271,41.9994,2020/5/24 11:50
1012,United States,"Boston, MA",10 High Street,10 High Street,2110,COB,COGENT,-74.4696,40.4893,2020/5/24 11:50
1013,United States,"Boston, MA",10 Post Office Sq,Ten Post Office Sq (North & South),2109,COB,COGENT,-71.0558,42.3573,2020/5/24 11:50
1014,United States,"Boston, MA",10 St James Ave,10 St James Ave,2116,COB,COGENT,-71.0706,42.3511,2020/5/24 11:50
1015,United States,"Boston, MA",100 Federal St,Bank of America Bldg (Formerly BankBoston Bldg),2110,COB,COGENT,-71.056,42.355,2020/5/24 11:50
1016,United States,"Boston, MA",100 Franklin St,100 Franklin St,2110,COB,COGENT,-93.2753,31.1156,2020/5/24 11:50
1017,United States,"Boston, MA",100 High Street,100 High St (aka 150 Federal St),2110,COB,COGENT,-79.098,35.9185,2020/5/24 11:50
1018,United States,"Boston, MA",100 Oliver Street,1 International Pl,2110,COB,COGENT,-95.7651,37.2231,2020/5/24 11:50
1019,United States,"Boston, MA",100 Summer St,100 Summer St,2110,COB,COGENT,-71.0574,42.3538,2020/5/24 11:50
1020,United States,"Boston, MA",101 Arch St,101 Arch St,2110,COB,COGENT,-71.0588,42.355,2020/5/24 11:50
1021,United States,"Boston, MA",101 Federal St,101 Federal St,2110,COB,COGENT,-71.0568,42.355,2020/5/24 11:50
1022,United States,"Boston, MA",101 Huntington Ave,Prudential Center,2199,COB,COGENT,-71.0809,42.3469,2020/5/24 11:50
1023,United States,"Boston, MA",111 Huntington Ave,Prudential Center,2199,COB,COGENT,-71.0819,42.3473,2020/5/24 11:50
1024,United States,"Boston, MA",116 Huntington Ave,116 Huntington Ave,2116,COB,COGENT,-71.0798,42.3466,2020/5/24 11:50
1027,United States,"Boston, MA",125 Summer St,125 Summer St,2110,COB,COGENT,-71.0575,42.3531,2020/5/24 11:50
1028,United States,"Boston, MA",131 Oliver Street,131 Oliver Street,2110,COB,COGENT,-71.0525,42.3552,2020/5/24 11:50
1029,United States,"Boston, MA",150 Oliver,2 International Pl,2110,COB,COGENT,-83.3442,33.9632,2020/5/24 11:50
1030,United States,"Boston, MA",155 Federal St,155 Federal St,2110,COB,COGENT,-71.0564,42.3537,2020/5/24 11:50
1031,United States,"Boston, MA",155 Seaport Blvd,Seaport World Trade Center West,2210,COB,COGENT,-71.0428,42.3498,2020/5/24 11:50
1032,United States,"Boston, MA",160 Federal St,Landmark,2110,COB,COGENT,-71.056,42.3541,2020/5/24 11:50
1033,United States,"Boston, MA",167 Milk St,167 Milk St,2109,COB,COGENT,-71.0535,42.3581,2020/5/24 11:50
1034,United States,"Boston, MA",175 Federal St,Fiduciary Trust Bldg,2110,COB,COGENT,-71.0564,42.3532,2020/5/24 11:50
1035,United States,"Boston, MA",18 Tremont St,18 Tremont St,2120,COB,COGENT,-71.0596,42.3588,2020/5/24 11:50
1036,United States,"Boston, MA",197 Clarendon St,Hancock Extension,2116,COB,COGENT,-71.0744,42.3495,2020/5/24 11:50
1037,United States,"Boston, MA",2 Center Plz,Two Center Plaza,2108,COB,COGENT,-71.0601,42.3596,2020/5/24 11:50
1038,United States,"Boston, MA",2 Oliver St,2 Oliver St,2109,COB,COGENT,-71.0545,42.3572,2020/5/24 11:50
1039,United States,"Boston, MA",20 Custom House Street,20 Custom House Road,2110,COB,COGENT,-71.0534,42.3581,2020/5/24 11:50
1040,United States,"Boston, MA",200 Berkeley St,Old Hancock Tower,2116,COB,COGENT,-71.0728,42.3499,2020/5/24 11:50
1041,United States,"Boston, MA",200 Clarendon St,John Hancock Tower,2116,COB,COGENT,-71.0751,42.3493,2020/5/24 11:50
1042,United States,"Boston, MA",200 State St,200 State St - South Building,2109,COB,COGENT,-71.0532,42.36,2020/5/24 11:50
1043,United States,"Boston, MA",200 State Street North,200 State Street - North Building,2109,COB,COGENT,-71.0532,42.36,2020/5/24 11:50
1044,United States,"Boston, MA",21 Custom House Street,21 Custom House Road,2110,COB,COGENT,-71.0526,42.358,2020/5/24 11:50
1045,United States,"Boston, MA",22 Batterymarch St,22 Batterymarch St,2109,COB,COGENT,-71.0545,42.3579,2020/5/24 11:50
1046,United States,"Boston, MA",222 Berkeley St,Two Twenty Two Berkeley,2116,COB,COGENT,-71.0733,42.3507,2020/5/24 11:50
1047,United States,"Boston, MA",225 Franklin St,225 Franklin St,2110,COB,COGENT,-71.054,42.356,2020/5/24 11:50
1048,United States,"Boston, MA",255 State St,Eaton Vance Bldg,2109,COB,COGENT,-71.0514,42.3594,2020/5/24 11:50
1049,United States,"Boston, MA",260 Franklin St,260 Franklin St,2110,COB,COGENT,-71.054,42.3569,2020/5/24 11:50
1050,United States,"Boston, MA",265 Franklin St,Paine Webber Bldg,2110,COB,COGENT,-71.0534,42.3565,2020/5/24 11:50
1051,United States,"Boston, MA",28 State St,28 State St,2109,COB,COGENT,-71.0575,42.3592,2020/5/24 11:50
1053,United States,"Boston, MA",3 Center Plz,Three Center Plaza,2108,COB,COGENT,-71.0609,42.3605,2020/5/24 11:50
1054,United States,"Boston, MA",31 St James Ave,Park Square Building,2116,COB,COGENT,-71.0716,42.3509,2020/5/24 11:50
1055,United States,"Boston, MA",33 Arch St,33 Arch St,2110,COB,COGENT,-71.0582,42.3561,2020/5/24 11:50
1056,United States,"Boston, MA",380 Stuart St,Stuart Bldg,2116,COB,COGENT,-71.0732,42.3491,2020/5/24 11:50
1057,United States,"Boston, MA",399 Boylston St,399 Boylston St,2116,COB,COGENT,-71.0722,42.3518,2020/5/24 11:50
1058,United States,"Boston, MA",40 Broad St,Insurance Exchange Bldg,2109,COB,COGENT,-74.0118,40.706,2020/5/24 11:50
1059,United States,"Boston, MA",451 D St (a/ka 70 Fargo & 80 Fargo),Seaport Center,2210,COB,COGENT,-71.0423,42.3452,2020/5/24 11:50
1060,United States,"Boston, MA",50 Congress Street,50 Congress Street,2109,COB,COGENT,-79.9906,40.4396,2020/5/24 11:50
1061,United States,"Boston, MA",500 Boylston St,Five Hundred Boylston,2116,COB,COGENT,-71.0744,42.3508,2020/5/24 11:50
1063,United States,"Boston, MA",53 State St,Exchange Place,2109,COB,COGENT,-71.0563,42.3587,2020/5/24 11:50
1064,United States,"Boston, MA",56 Roland St - 2nd Floor,Crown Castle (fka. Sidera/NEON),2129,CNDC,COGENT,-71.0801,42.3813,2020/5/24 11:50
1065,United States,"Boston, MA",60 State St,60 State St,2109,COB,COGENT,-71.0566,42.3594,2020/5/24 11:50
1066,United States,"Boston, MA",600 Atlantic Ave,Federal Reserve Plaza,2210,COB,COGENT,-71.0534,42.3528,2020/5/24 11:50
1067,United States,"Boston, MA",75 Federal St,75 Federal St,2110,COB,COGENT,-71.0568,42.355,2020/5/24 11:50
1068,United States,"Boston, MA",75 State St,75 Street St.,2109,COB,COGENT,-71.055,42.3586,2020/5/24 11:50
1069,United States,"Boston, MA",75-81 Arlington St,75-81 Arlington St,2116,COB,COGENT,-71.0701,42.3506,2020/5/24 11:50
1071,United States,"Boston, MA",8 Ashburton Place,Suffolk University - Sawyer Building,2108,COB,COGENT,-71.0618,42.359,2020/5/24 11:50
1072,United States,"Boston, MA",800 Boylston St,800 Boylston St,2199,COB,COGENT,-71.082,42.3474,2020/5/24 11:50
1073,United States,"Boston, MA",888 Boylston Street,Prudential Center,2199,COB,COGENT,-71.084,42.3476,2020/5/24 11:50
1074,United States,"Boston, MA",99 High St,99 High St,2110,COB,COGENT,-71.0546,42.3541,2020/5/24 11:50
1075,United States,"Boston, MA",99 Summer St,99 Summer St,2110,COB,COGENT,-71.0586,42.3535,2020/5/24 11:50
1076,United States,"Breinigsville, PA",9999 Hamilton Blvd,Tierpoint-TEK (Tek Park),18031,CNDC,COGENT,-75.6557,40.5442,2020/5/24 11:50
1077,United States,"Brentwood, TN",7100 Commerce Way,"Flexential NAS01, NAS04",37027,CNDC,COGENT,-86.8087,35.9725,2020/5/24 11:50
1078,United States,"Brooklyn, NY",1 Metrotech Center,1 Metrotech Center,11201,COB,COGENT,-73.9867,40.6933,2020/5/24 11:50
1081,United States,"Brooklyn, NY",15 Metrotech Center,15 Metrotech Center,11201,COB,COGENT,-73.9839,40.694,2020/5/24 11:50
1082,United States,"Brooklyn, NY",2 Metrotech Center,2 Metrotech Center,11201,COB,COGENT,-73.9857,40.6934,2020/5/24 11:50
1084,United States,"Brooklyn, NY","21, 55, 83 33rd St","Industry City, Building 9",22322,COB,COGENT,-73.9087,40.7759,2020/5/24 11:50
1085,United States,"Brooklyn, NY","219, 253, 36th St","Industry City, Building 3",11232,COB,COGENT,-74.0073,40.6562,2020/5/24 11:50
1087,United States,"Brooklyn, NY","32, 68, 86 33rd St","Industry City, Building 8",11232,COB,COGENT,-73.9254,40.7597,2020/5/24 11:50
1090,United States,"Brooklyn, NY","34, 68, 88 34th St","Industry City, Building 6",11232,COB,COGENT,-73.9876,40.7496,2020/5/24 11:50
1091,United States,"Brooklyn, NY","34, 68, 88 35th St","Industry City, Building 4",11232,COB,COGENT,-74.0072,40.6567,2020/5/24 11:50
1092,United States,"Brooklyn, NY",4014 1st Ave,"Industry City, Building 26",11232,COB,COGENT,-74.0138,40.6562,2020/5/24 11:50
1094,United States,"Brooklyn, NY",80 39th St,"Industry City, Building 22",11232,COB,COGENT,-73.9973,40.6473,2020/5/24 11:50
1095,United States,"Brooklyn, NY",882 3rd Ave,"Industry City, Building 10",11232,COB,COGENT,-74.0049,40.6578,2020/5/24 11:50
1096,United States,"Brooklyn, NY",882 3rd Ave - 9th Floor,ColoGuard CGNY1,11232,CNDC,COGENT,-74.0049,40.6578,2020/5/24 11:50
1098,United States,"Buffalo, NY","350 Main St, Suite 102",365 Data Centers (fka. Equinix - S&D),14202,CNDC,COGENT,-88.1747,43.3312,2020/5/24 11:50
1099,United States,"Buffalo, NY","350 Main St, Suite 550",Shatter IT,14202,CNDC,COGENT,-88.1747,43.3312,2020/5/24 11:50
1100,United States,"Burbank, CA",3015 Winona Ave,Cyxtera BR1,91504,CNDC,COGENT,-118.344,34.1998,2020/5/24 11:50
1101,United States,"Cambridge, MA",1 Kendall Sq,Building 400,2139,COB,COGENT,-71.0901,42.3673,2020/5/24 11:50
1102,United States,"Cambridge, MA",1 Kendall Square (Building 100),Building 100,2139,COB,COGENT,-71.0824,42.3641,2020/5/24 11:50
1103,United States,"Cambridge, MA",1 Kendall Square (Building 600),Building 600,2139,COB,COGENT,-71.0824,42.3641,2020/5/24 11:50
1104,United States,"Cambridge, MA",1 Kendall Square (Building 1400),Building 1400,2139,COB,COGENT,-71.0901,42.3673,2020/5/24 11:50
1105,United States,"Cambridge, MA",1 Kendall Square (Building 200),Building 200,2139,COB,COGENT,-71.0901,42.3673,2020/5/24 11:50
1106,United States,"Cambridge, MA",1 Kendall Square (Building 300),Building 300,2139,COB,COGENT,-71.0824,42.3641,2020/5/24 11:50
1107,United States,"Cambridge, MA",1 Kendall Square (Building 700),Building 700,2139,COB,COGENT,-71.0901,42.3673,2020/5/24 11:50
1108,United States,"Cambridge, MA",1 Main Street,1 Main Street,2142,COB,COGENT,-73.9903,40.7037,2020/5/24 11:50
1109,United States,"Cambridge, MA",1 Memorial Drive,1 Memorial Drive,2142,COB,COGENT,-95.5143,29.7513,2020/5/24 11:50
1110,United States,"Cambridge, MA",100 Technology Square,100 Technology Square,2139,COB,COGENT,-71.0908,42.3632,2020/5/24 11:50
1111,United States,"Cambridge, MA",101 Main Street,101 Main Street,2142,COB,COGENT,-71.0816,42.3622,2020/5/24 11:50
1112,United States,"Cambridge, MA",200 Technology Square,200 Technology Square,2139,COB,COGENT,-71.0907,42.3633,2020/5/24 11:50
1113,United States,"Cambridge, MA",215 First St,Athenaeum Center,2142,COB,COGENT,-71.0792,42.3643,2020/5/24 11:50
1114,United States,"Cambridge, MA",255 Main Street,Kendall Center,2142,COB,COGENT,-70.2798,41.6545,2020/5/24 11:50
1115,United States,"Cambridge, MA",300 Bent St - 1st Floor,CenturyLink (fka. Level 3),2141,CNDC,COGENT,-71.0865,42.3676,2020/5/24 11:50
1116,United States,"Cambridge, MA",300 Technology Square,300 Technology Square,2139,COB,COGENT,-71.0916,42.3634,2020/5/24 11:50
1117,United States,"Cambridge, MA",400 Technology Square,400 Technology Square,2139,COB,COGENT,-71.0918,42.364,2020/5/24 11:50
1118,United States,"Cambridge, MA",500 Technology Square,500 Technology Square,2139,COB,COGENT,-71.0926,42.3634,2020/5/24 11:50
1119,United States,"Cambridge, MA",600 Technology Square,600 Technology Square,2139,COB,COGENT,-71.0931,42.3634,2020/5/24 11:50
1120,United States,"Cambridge, MA",700 Technology Square,700 Technology Square,2139,COB,COGENT,-71.0928,42.364,2020/5/24 11:50
1121,United States,"Carlstadt, NJ",410 Commerce Blvd,SunGard CRL-410,7072,CNDC,COGENT,-74.0498,40.828,2020/5/24 11:50
1122,United States,"Carlstadt, NJ",760 Washington Ave,SunGard crl-760,7072,CNDC,COGENT,-74.0559,40.8307,2020/5/24 11:50
1123,United States,"Carlstadt, NJ",777 Central Blvd,SunGard crl-777,7072,CNDC,COGENT,-74.0443,40.8283,2020/5/24 11:50
1124,United States,"Carrollton, TX",1649 W. Frankford Rd,CyrusOne Dallas - Carrollton,75007,CNDC,COGENT,-96.93,32.9918,2020/5/24 11:50
1125,United States,"Carrollton, TX",2440 Marsh Lane,Digital Realty,75006,CNDC,COGENT,-96.8537,32.9801,2020/5/24 11:50
1126,United States,"Carteret, NJ",1400 Federal Blvd,Equinix NY 11 (fka. Verizon),7008,CNDC,COGENT,-74.2434,40.5843,2020/5/24 11:50
1127,United States,"Centennial, CO",12500 E ARAPAHOE Rd D,Flexential Centennial (fka. Peak 10 / ViaWest),80112,CNDC,COGENT,-104.843,39.5942,2020/5/24 11:50
1128,United States,"Centennial, CO",6900 S Peoria Street,zColo (fka. Latisys),80112,CNDC,COGENT,-104.847,39.5913,2020/5/24 11:50
1129,United States,"Chandler, AZ",2121 S Price Rd,Digital Realty,85286,CNDC,COGENT,-111.887,33.2748,2020/5/24 11:50
1130,United States,"Chandler, AZ",2305 Ellis Street-Bldg. 4,CyrusOne Phoenix - Chandler Bldg. 4 (SalesForce.co,85286,CNDC,COGENT,-111.883,33.272,2020/5/24 11:50
1132,United States,"Chandler, AZ",2335 South Ellis Street,CyrusOne Phoenix - Chandler Bldg. 3,85286,CNDC,COGENT,-111.882,33.2707,2020/5/24 11:50
1133,United States,"Chandler, AZ",2405 South Ellis Street,Chandler V (Bldg 2),85286,CNDC,COGENT,-111.882,33.2698,2020/5/24 11:50
1134,United States,"Chandler, AZ",2600 W Germann Rd,H5 Phoenix (fka. NextFort),85286,CNDC,COGENT,-111.886,33.2771,2020/5/24 11:50
1135,United States,"Chandler, AZ",2605 South Ellis Street,Chandler Vl (Bldg 8),85286,CNDC,COGENT,-111.883,33.2674,2020/5/24 11:50
1136,United States,"Charlestown, MA",500 Rutherford Ave,Hood Plant,2129,COB,COGENT,-71.0727,42.381,2020/5/24 11:50
1137,United States,"Charlotte, NC",101 N Tryon St (Independence Center),Independence Center,28246,COB,COGENT,-80.8422,35.2274,2020/5/24 11:50
1138,United States,"Charlotte, NC",101 S Tryon St (BofA Plaza),Bank of America Plaza,28280,COB,COGENT,-80.8422,35.2274,2020/5/24 11:50
1139,United States,"Charlotte, NC",10105 David Taylor Drive,Flexential CLT04,28262,CNDC,COGENT,-80.7654,35.3299,2020/5/24 11:50
1140,United States,"Charlotte, NC",113 N Myers St,Digital Realty (fka. Telx MMR),28202,CNDC,COGENT,-80.8356,35.2219,2020/5/24 11:50
1141,United States,"Charlotte, NC",121 W Trade St,Interstate Tower,28202,COB,COGENT,-80.8438,35.2275,2020/5/24 11:50
1142,United States,"Charlotte, NC",125 N Myers St,Tierpoint CL2 (fka. DRT),28202,CNDC,COGENT,-80.8353,35.2221,2020/5/24 11:50
1143,United States,"Charlotte, NC",128 S. Tryon Street,First Citizens Plaza,28202,COB,COGENT,-80.8442,35.2269,2020/5/24 11:50
1144,United States,"Charlotte, NC",1612 Cross Beam Drive,Lumos Data Centers CLT4,28217,CNDC,COGENT,-80.9264,35.1801,2020/5/24 11:50
1145,United States,"Charlotte, NC",200 S College St,BB&T Center,28202,COB,COGENT,-80.8439,35.2252,2020/5/24 11:50
1146,United States,"Charlotte, NC",201 N Tryon St,Fifth Third Center,28202,COB,COGENT,-80.8417,35.2287,2020/5/24 11:50
1147,United States,"Charlotte, NC",201 S College St,Charlotte Plaza,28244,COB,COGENT,-80.8433,35.2249,2020/5/24 11:50
1148,United States,"Charlotte, NC",201 S Tryon St,Tryon Square,28202,COB,COGENT,-80.844,35.2261,2020/5/24 11:50
1149,United States,"Charlotte, NC",214 N Tryon St,Hearst Tower,28280,COB,COGENT,-80.8413,35.228,2020/5/24 11:50
1150,United States,"Charlotte, NC",227 West Trade St.,Carillon,28202,COB,COGENT,-80.8453,35.2279,2020/5/24 11:50
1151,United States,"Charlotte, NC",300 South Tryon Street,300 South Tryon,28202,COB,COGENT,-80.8458,35.2256,2020/5/24 11:50
1152,United States,"Charlotte, NC",301 S College St,Wells Fargo Center,28202,COB,COGENT,-80.8442,35.2235,2020/5/24 11:50
1153,United States,"Charlotte, NC",3101 International Airport Dr,Lumos Data Centers (fka. DC74),28208,CNDC,COGENT,-80.9385,35.1938,2020/5/24 11:50
1154,United States,"Charlotte, NC",550 S Tryon St,Duke Energy Center,28202,COB,COGENT,-80.8486,35.2241,2020/5/24 11:50
1155,United States,"Charlotte, NC",701 E Trade Street,Cogent Data Center - Charlotte,28202,CDC,COGENT,-80.8438,35.2275,2020/5/24 11:50
1156,United States,"Charlotte, NC",8910 Lenox Pointe Dr.,Flexential CLT01-03,28273,CNDC,COGENT,-80.9369,35.1354,2020/5/24 11:50
1157,United States,"Chaska, MN",3500 Lyman Blvd,Flexential Minneapolis (fmr Peak 10 ViaWest),55318,CNDC,COGENT,-93.5995,44.8513,2020/5/24 11:50
1158,United States,"Chicago, IL",1 E Wacker Dr,One East Wacker Bldg,60601,COB,COGENT,-87.6275,41.8864,2020/5/24 11:50
1159,United States,"Chicago, IL",1 N Dearborn St,One North Dearborn Bldg,60602,COB,COGENT,-87.6286,41.8823,2020/5/24 11:50
1160,United States,"Chicago, IL",1 N Franklin St,One North Franklin Bldg,60606,COB,COGENT,-87.6352,41.8978,2020/5/24 11:50
1161,United States,"Chicago, IL",1 N State St,One North State Bldg,60602,COB,COGENT,-87.6273,41.8824,2020/5/24 11:50
1162,United States,"Chicago, IL",1 N Wacker Dr.,UBS Building,60606,COB,COGENT,-87.6361,41.8823,2020/5/24 11:50
1163,United States,"Chicago, IL",1 S Dearborn,One South Dearborn,60603,COB,COGENT,-87.6288,41.8816,2020/5/24 11:50
1164,United States,"Chicago, IL",1 S Wacker Dr,One South Wacker Bldg,60606,COB,COGENT,-87.636,41.8815,2020/5/24 11:50
1166,United States,"Chicago, IL",10 S LaSalle St,Chase Plaza,60603,COB,COGENT,-87.6323,41.8819,2020/5/24 11:50
1167,United States,"Chicago, IL",10 S Riverside Plz,10 S Riverside Plz,60606,COB,COGENT,-87.6391,41.8813,2020/5/24 11:50
1168,United States,"Chicago, IL",10 S Wacker Dr,10 S Wacker Dr,60606,COB,COGENT,-87.6371,41.8815,2020/5/24 11:50
1169,United States,"Chicago, IL",100 S Wacker Dr,Hartford Plaza,60606,COB,COGENT,-87.6375,41.8803,2020/5/24 11:50
1170,United States,"Chicago, IL",101 N Wacker Dr,101 N Wacker Dr,60606,COB,COGENT,-87.6365,41.8836,2020/5/24 11:50
1171,United States,"Chicago, IL",105 W Adams St,The Clark-Adams Bldg,60603,COB,COGENT,-99.9262,41.2931,2020/5/24 11:50
1172,United States,"Chicago, IL",11 E Adams St,Lasalle Talman Bank Bldg,60603,COB,COGENT,-87.6271,41.8793,2020/5/24 11:50
1173,United States,"Chicago, IL",111 E Wacker Dr,One Illinois Center,60601,COB,COGENT,-87.6236,41.8876,2020/5/24 11:50
1174,United States,"Chicago, IL",111 N Canal St - Suite 200,CenturyLink (fka. Level 3),60606,CNDC,COGENT,-87.6394,41.8836,2020/5/24 11:50
1175,United States,"Chicago, IL",111 N Wabash Ave,Garland Bldg,60602,COB,COGENT,-87.6259,41.8835,2020/5/24 11:50
1176,United States,"Chicago, IL",111 S Wacker Dr,111 S Wacker,60606,COB,COGENT,-87.6365,41.8803,2020/5/24 11:50
1177,United States,"Chicago, IL",111 W Jackson Blvd,111 W Jackson Blvd,60604,COB,COGENT,-87.6313,41.878,2020/5/24 11:50
1178,United States,"Chicago, IL",111 W Monroe St,Harris Bank Bldg,60603,COB,COGENT,-87.6312,41.8803,2020/5/24 11:50
1179,United States,"Chicago, IL",115 S LaSalle St,Harris Bank Bldg,60603,COB,COGENT,-87.6319,41.8804,2020/5/24 11:50
1180,United States,"Chicago, IL",120 N LaSalle St,120 N LaSalle St,60602,COB,COGENT,-87.6331,41.8837,2020/5/24 11:50
1181,United States,"Chicago, IL",120 S LaSalle St,120 S LaSalle St,60603,COB,COGENT,-87.6326,41.8803,2020/5/24 11:50
1182,United States,"Chicago, IL",120 S Riverside Plz,120 S Riverside Plz,60606,COB,COGENT,-87.6391,41.88,2020/5/24 11:50
1183,United States,"Chicago, IL",121 W Wacker (fka 221 N LaSalle),LaSalle-Wacker Building,60601,COB,COGENT,-87.6319,41.8867,2020/5/24 11:50
1184,United States,"Chicago, IL",123 N Wacker Dr,123 N Wacker Dr,60606,COB,COGENT,-87.6365,41.8841,2020/5/24 11:50
1185,United States,"Chicago, IL",125 S Wacker Dr,125 S Wacker Dr,60606,COB,COGENT,-87.6364,41.8797,2020/5/24 11:50
1186,United States,"Chicago, IL",130 E Randolph St,One Prudential Plaza,60601,COB,COGENT,-87.6231,41.885,2020/5/24 11:50
1187,United States,"Chicago, IL",131 South Dearborn,Citadel Center,60603,COB,COGENT,-87.6288,41.88,2020/5/24 11:50
1188,United States,"Chicago, IL",135 S LaSalle St,Lasalle Bank Bldg,60603,COB,COGENT,-87.6314,41.8798,2020/5/24 11:50
1189,United States,"Chicago, IL",14 E Jackson Blvd/247 S State St,The Hub,60604,COB,COGENT,-87.6275,41.8785,2020/5/24 11:50
1190,United States,"Chicago, IL",140 S Dearborn St,Marquette Bldg,60603,COB,COGENT,-87.6295,41.8798,2020/5/24 11:50
1191,United States,"Chicago, IL",141 W Jackson Blvd,Chicago Board of Trade,60604,COB,COGENT,-87.6324,41.878,2020/5/24 11:50
1192,United States,"Chicago, IL",141 W Jackson Blvd,Special Colo - Dow pit IC 1 cage,60604,COB,COGENT,-87.6324,41.878,2020/5/24 11:50
1193,United States,"Chicago, IL",141 W Jackson Blvd - Annex,Board of Trade Annex,60604,COB,COGENT,-87.6325,41.878,2020/5/24 11:50
1194,United States,"Chicago, IL",141 W Jackson Blvd Trading Floor,Special Colo - Agriculture floor,60604,COB,COGENT,-87.6325,41.878,2020/5/24 11:50
1195,United States,"Chicago, IL",150 N Michigan Ave,Stone Container Bldg,60601,COB,COGENT,-87.6249,41.8848,2020/5/24 11:50
1196,United States,"Chicago, IL",150 N Riverside,150 N Riverside,60606,COB,COGENT,-87.6385,41.8855,2020/5/24 11:50
1197,United States,"Chicago, IL",150 S Wacker Dr,Hartford Plaza,60606,COB,COGENT,-87.6374,41.8796,2020/5/24 11:50
1198,United States,"Chicago, IL",151 N Franklin,CNA Center,60606,COB,COGENT,-87.6346,41.8846,2020/5/24 11:50
1199,United States,"Chicago, IL",155 N Wacker Dr,155 N Wacker Dr,60606,COB,COGENT,-87.6363,41.8847,2020/5/24 11:50
1200,United States,"Chicago, IL",161 N Clark St,Chicago Title & Trust Ctr I,60601,COB,COGENT,-87.6307,41.8849,2020/5/24 11:50
1201,United States,"Chicago, IL",175 W Jackson Blvd,Insurance Exchange Bldg,60604,COB,COGENT,-87.6332,41.8779,2020/5/24 11:50
1202,United States,"Chicago, IL",180 N LaSalle St,180 N LaSalle St,60601,COB,COGENT,-87.633,41.8855,2020/5/24 11:50
1203,United States,"Chicago, IL",180 N Stetson Ave,Two Prudential Plaza,60601,COB,COGENT,-87.6223,41.8862,2020/5/24 11:50
1204,United States,"Chicago, IL",181 W Madison St,181 W Madison St,60602,COB,COGENT,-87.6335,41.8817,2020/5/24 11:50
1205,United States,"Chicago, IL",190 S LaSalle St,190 S LaSalle St,60603,COB,COGENT,-87.6328,41.8798,2020/5/24 11:50
1206,United States,"Chicago, IL",191 N Wacker Dr,191 N Wacker Dr,60606,COB,COGENT,-87.6366,41.8854,2020/5/24 11:50
1207,United States,"Chicago, IL",2 N LaSalle St,Two North LaSalle Bldg,60602,COB,COGENT,-87.6326,41.8824,2020/5/24 11:50
1208,United States,"Chicago, IL",2 North Riverside Plaza,2 North Riverside Plaza,60606,COB,COGENT,-87.636,41.8861,2020/5/24 11:50
1209,United States,"Chicago, IL",20 N Clark St,20 N Clark St,60602,COB,COGENT,-87.6311,41.8823,2020/5/24 11:50
1210,United States,"Chicago, IL",20 N Wacker Dr,Civic Opera Bldg,60606,COB,COGENT,-87.6373,41.8826,2020/5/24 11:50
1211,United States,"Chicago, IL",20 S Clark St,Two First National Plaza,60603,COB,COGENT,-87.6311,41.8823,2020/5/24 11:50
1212,United States,"Chicago, IL",200 E Randolph St,Aon Center,60601,COB,COGENT,-87.6215,41.8853,2020/5/24 11:50
1213,United States,"Chicago, IL",200 N LaSalle St,200 N LaSalle St,60601,COB,COGENT,-87.633,41.8861,2020/5/24 11:50
1214,United States,"Chicago, IL",200 S Wacker Dr,200 S Wacker Dr,60606,COB,COGENT,-87.6372,41.8789,2020/5/24 11:50
1215,United States,"Chicago, IL",200 W Adams St,200 W Adams St,60606,COB,COGENT,-87.6341,41.8798,2020/5/24 11:50
1216,United States,"Chicago, IL",200 W Jackson Blvd,200 W Jackson Blvd,60606,COB,COGENT,-87.6341,41.8784,2020/5/24 11:50
1217,United States,"Chicago, IL",200 W Madison St,Madison Plaza,60606,COB,COGENT,-87.6343,41.8823,2020/5/24 11:50
1218,United States,"Chicago, IL",200 W Monroe St,200 W Monroe St,60606,COB,COGENT,-87.6343,41.8807,2020/5/24 11:50
1219,United States,"Chicago, IL",203 N LaSalle St,203 N LaSalle St,60601,COB,COGENT,-87.6315,41.8862,2020/5/24 11:50
1220,United States,"Chicago, IL",205 N Michigan Ave,Michigan Plaza,60601,COB,COGENT,-87.6236,41.8861,2020/5/24 11:50
1221,United States,"Chicago, IL",208 S LaSalle St,208 S LaSalle St,60604,COB,COGENT,-87.6325,41.8791,2020/5/24 11:50
1222,United States,"Chicago, IL",216 W Jackson Blvd,Cogent Data Center - Chicago,60606,CDC,COGENT,-87.6345,41.8785,2020/5/24 11:50
1223,United States,"Chicago, IL",216 W Jackson Blvd,216 West Jackson Bldg,60606,COB,COGENT,-87.6345,41.8785,2020/5/24 11:50
1224,United States,"Chicago, IL",222 Merchandise Mart Plaza,Merchandise Mart Plaza,60654,COB,COGENT,-87.6355,41.8884,2020/5/24 11:50
1225,United States,"Chicago, IL",222 N LaSalle St,222 N LaSalle St,60601,COB,COGENT,-87.6328,41.8866,2020/5/24 11:50
1226,United States,"Chicago, IL",222 S Riverside Plz,222 S Riverside Plz,60606,COB,COGENT,-87.6393,41.8785,2020/5/24 11:50
1227,United States,"Chicago, IL",222 W Adams St,USG Bldg,60606,COB,COGENT,-87.6347,41.8795,2020/5/24 11:50
1228,United States,"Chicago, IL",223 W Jackson Blvd,Brooks Bldg,60606,COB,COGENT,-87.6345,41.878,2020/5/24 11:50
1229,United States,"Chicago, IL",225 N Michigan Ave,Michigan Plaza,60601,COB,COGENT,-87.6238,41.8864,2020/5/24 11:50
1230,United States,"Chicago, IL",225 W Wacker Dr,225 W Wacker Dr,60606,COB,COGENT,-87.6351,41.8863,2020/5/24 11:50
1231,United States,"Chicago, IL",225 W Washington St,225 W Washington St,60606,COB,COGENT,-89.4586,41.3779,2020/5/24 11:50
1232,United States,"Chicago, IL",227 W Monroe St,AT&T Corporate Center 1,60606,COB,COGENT,-87.6348,41.8796,2020/5/24 11:50
1233,United States,"Chicago, IL",230 W Monroe St,230 W Monroe St,60606,COB,COGENT,-87.6351,41.8811,2020/5/24 11:50
1234,United States,"Chicago, IL",233 N Michigan Ave,Two Illinois Center,60601,COB,COGENT,-87.6239,41.887,2020/5/24 11:50
1235,United States,"Chicago, IL",233 S Wacker Dr,Willis Tower,60606,COB,COGENT,-87.6358,41.8789,2020/5/24 11:50
1236,United States,"Chicago, IL",2800 South Ashland Avenue,QTS Datacenter,60608,CNDC,COGENT,-87.6671,41.8422,2020/5/24 11:50
1237,United States,"Chicago, IL",30 N LaSalle St,30 N LaSalle St,60602,COB,COGENT,-87.6328,41.8829,2020/5/24 11:50
1238,United States,"Chicago, IL",30 S Wacker Dr,30 S Wacker Dr,60606,COB,COGENT,-87.6372,41.881,2020/5/24 11:50
1239,United States,"Chicago, IL",300 North LaSalle,300 North LaSalle,60654,COB,COGENT,-87.6327,41.8881,2020/5/24 11:50
1240,United States,"Chicago, IL",300 S Riverside Plz,Gateway IV,60606,COB,COGENT,-87.6386,41.8775,2020/5/24 11:50
1241,United States,"Chicago, IL",300 S Wacker Dr,300 S Wacker Dr,60606,COB,COGENT,-87.6371,41.8776,2020/5/24 11:50
1242,United States,"Chicago, IL",303 E Wacker Dr,Three Illinois Center,60601,COB,COGENT,-87.6205,41.8872,2020/5/24 11:50
1243,United States,"Chicago, IL",303 W Madison St,303 W Madison St,60606,COB,COGENT,-87.6356,41.8816,2020/5/24 11:50
1244,United States,"Chicago, IL",311 S Wacker Dr,311 S Wacker Dr,60606,COB,COGENT,-87.6357,41.8774,2020/5/24 11:50
1245,United States,"Chicago, IL",318 W Adams St,318 W Adams St,60606,COB,COGENT,-87.636,41.8797,2020/5/24 11:50
1246,United States,"Chicago, IL",321 N Clark St,Riverfront Plaza,60654,COB,COGENT,-87.6306,41.8882,2020/5/24 11:50
1247,United States,"Chicago, IL",33 N Dearborn St,33 N Dearborn St,60602,COB,COGENT,-87.6289,41.8829,2020/5/24 11:50
1248,United States,"Chicago, IL",33 N LaSalle St,33 N LaSalle St,60602,COB,COGENT,-87.6322,41.8828,2020/5/24 11:50
1249,United States,"Chicago, IL",33 W Monroe St,33 W Monroe St,60603,COB,COGENT,-87.629,41.8808,2020/5/24 11:50
1250,United States,"Chicago, IL",330 N Wabash Ave,One IBM Plaza,60611,COB,COGENT,-87.6275,41.8886,2020/5/24 11:50
1252,United States,"Chicago, IL",333 W Wacker Dr,333 W Wacker Dr,60606,COB,COGENT,-87.636,41.8861,2020/5/24 11:50
1253,United States,"Chicago, IL",35 E Wacker Dr,35 E Wacker Dr,60601,COB,COGENT,-87.6268,41.8865,2020/5/24 11:50
1254,United States,"Chicago, IL",350 E Cermak Rd,Advantage Futures,60616,CNDC,COGENT,-87.6185,41.8548,2020/5/24 11:50
1255,United States,"CHICAGO, IL",350 E CERMAK RD,Equinix CH2,60616,CNDC,COGENT,-87.6185,41.8548,2020/5/24 11:50
1256,United States,"CHICAGO, IL",350 E CERMAK RD,Equinix CH4,60616,CNDC,COGENT,-87.6185,41.8548,2020/5/24 11:50
1257,United States,"Chicago, IL",350 E Cermak Rd,Equinix CH1,60616,CNDC,COGENT,-87.6185,41.8548,2020/5/24 11:50
1260,United States,"Chicago, IL",350 E Cermak Rd - 2nd floor,Digital Realty CHI1 (fka. Telx MMR),60616,CNDC,COGENT,-87.6185,41.8548,2020/5/24 11:50
1261,United States,"Chicago, IL",353 N Clark St (aka 65 W Kinzie St),353 North Clark,60654,COB,COGENT,-87.6305,41.889,2020/5/24 11:50
1262,United States,"Chicago, IL",400 N Michigan Ave,Wrigley Bldg-South Tower,60611,COB,COGENT,-87.6246,41.8769,2020/5/24 11:50
1263,United States,"Chicago, IL",400 S. LaSalle Street,400 S. LaSalle,60605,COB,COGENT,-87.6321,41.8765,2020/5/24 11:50
1264,United States,"Chicago, IL",401 N Michigan Ave,Equitable Bldg,60611,COB,COGENT,-87.6232,41.8899,2020/5/24 11:50
1265,United States,"Chicago, IL",401 S LaSalle St,401 S LaSalle St,60605,COB,COGENT,-87.6316,41.8766,2020/5/24 11:50
1266,United States,"Chicago, IL",410 N Michigan Ave,Wrigley Bldg-North Tower,60611,COB,COGENT,-87.6248,41.8899,2020/5/24 11:50
1268,United States,"Chicago, IL",427 S LaSalle,Cogent Data Center - Chicago,60605,CDC,COGENT,-87.6315,41.8762,2020/5/24 11:50
1269,United States,"Chicago, IL",427 S LaSalle St,CoreSite CH1,60605,CNDC,COGENT,-87.6315,41.8762,2020/5/24 11:50
1270,United States,"Chicago, IL",427 S LaSalle St Suite 405,365 Data Centers (fka. Equinix CH6 - S&D),60605,CNDC,COGENT,-87.6315,41.8762,2020/5/24 11:50
1271,United States,"Chicago, IL",427 S LaSalle St Suite 405,365 Data Centers (fka. Equinix CH6 - S&D),60605,CNDC,COGENT,-87.6315,41.8762,2020/5/24 11:50
1272,United States,"Chicago, IL",444 N Michigan Ave,444 N Michigan Ave,60611,COB,COGENT,-87.6248,41.8907,2020/5/24 11:50
1273,United States,"Chicago, IL",444 W Lake Street,River Point,60606,COB,COGENT,-87.6395,41.8862,2020/5/24 11:50
1274,United States,"Chicago, IL",455 Cityfront Plaza Dr,NBC Tower,60611,COB,COGENT,-87.621,41.8902,2020/5/24 11:50
1275,United States,"Chicago, IL",500 W Madison St,Citicorp Center,60661,COB,COGENT,-87.6406,41.8822,2020/5/24 11:50
1276,United States,"Chicago, IL",500 W Monroe St,500 W Monroe St,60661,COB,COGENT,-87.6399,41.8807,2020/5/24 11:50
1277,United States,"Chicago, IL",515 N State St,The AMA Bldg,60654,COB,COGENT,-87.6274,41.8914,2020/5/24 11:50
1278,United States,"Chicago, IL",525 W Monroe St,525 W Monroe St,60661,COB,COGENT,-87.6404,41.8802,2020/5/24 11:50
1279,United States,"Chicago, IL",540 W. Madison,Bank of America Plaza,60661,COB,COGENT,-87.6417,41.8826,2020/5/24 11:50
1280,United States,"Chicago, IL",55 E Monroe St,55 E Monroe St,60603,COB,COGENT,-87.6257,41.8805,2020/5/24 11:50
1281,United States,"Chicago, IL",55 W Monroe St,55 W Monroe St,60603,COB,COGENT,-87.6298,41.8805,2020/5/24 11:50
1282,United States,"Chicago, IL",55 W Wacker Dr,55 W Wacker Dr,60601,COB,COGENT,-87.6297,41.8865,2020/5/24 11:50
1283,United States,"Chicago, IL",550 W. Jackson Blvd,550 W. Jackson,60661,COB,COGENT,-87.6415,41.8783,2020/5/24 11:50
1284,United States,"Chicago, IL",600 S Federal St,Digital Realty CHI2 (fka. Telx MMR Suite 702),60605,CNDC,COGENT,-87.6299,41.8742,2020/5/24 11:50
1285,United States,"Chicago, IL",600 W Chicago Ave,600 W Chicago Ave,60610,COB,COGENT,-87.6435,41.8977,2020/5/24 11:50
1286,United States,"Chicago, IL",601 South LaSalle,601 South LaSalle,60605,COB,COGENT,-87.6318,41.8771,2020/5/24 11:50
1287,United States,"Chicago, IL",601 W Polk St,Tierpoint POL (fka. InnerNet Center),60607,CNDC,COGENT,-95.3882,29.756,2020/5/24 11:50
1288,United States,"Chicago, IL",625 N Michigan Ave,625 N Michigan Ave,60611,COB,COGENT,-87.6237,41.8932,2020/5/24 11:50
1289,United States,"Chicago, IL",70 W Madison St,Three First National Plaza,60602,COB,COGENT,-87.6304,41.8823,2020/5/24 11:50
1290,United States,"Chicago, IL",71 S Wacker Dr,Hyatt Center,60606,COB,COGENT,-87.636,41.8812,2020/5/24 11:50
1291,United States,"Chicago, IL",710 N Lake Shore Drive,Starlight,60611,CNDC,COGENT,-87.6173,41.8731,2020/5/24 11:50
1292,United States,"Chicago, IL",717 S Wells St,Inteliquent (fka. Neutral Tandem/Layer One),60607,CNDC,COGENT,-87.6333,41.8729,2020/5/24 11:50
1294,United States,"Chicago, IL",725 S Wells Street,725 South Wells,60607,COB,COGENT,-87.6333,41.8726,2020/5/24 11:50
1295,United States,"Chicago, IL","725 S Wells Street, Suite 300",Digital Capital Partners,60607,CNDC,COGENT,-87.6333,41.8726,2020/5/24 11:50
1296,United States,"Chicago, IL",737 N Michigan Ave,Olympia Centre,60611,COB,COGENT,-87.6236,41.8961,2020/5/24 11:50
1297,United States,"Chicago, IL",77 W Wacker Dr,R.R. Donnelley Bldg,60601,COB,COGENT,-87.6307,41.8864,2020/5/24 11:50
1298,United States,"Chicago, IL",840 S Canal Street,zcolo - Chicago,60607,CNDC,COGENT,-87.6402,41.8709,2020/5/24 11:50
1299,United States,"Chicago, IL",840 S Canal Street,ServerFarm,60607,CNDC,COGENT,-87.6402,41.8709,2020/5/24 11:50
1300,United States,"Chicago, IL",875 N Michigan Ave,John Hancock Center,60611,COB,COGENT,-87.6229,41.8988,2020/5/24 11:50
1301,United States,"Chicago, IL",900 N Michigan Ave,900 N Michigan Ave,60611,COB,COGENT,-87.6247,41.8995,2020/5/24 11:50
1302,United States,"Cincinnati, OH",201 E 5th St,201 E 5th St,45202,COB,COGENT,-95.3959,29.7798,2020/5/24 11:50
1303,United States,"Cincinnati, OH",221 East Fourth Street,Atrium Two,45202,COB,COGENT,-84.5086,39.1003,2020/5/24 11:50
1304,United States,"Cincinnati, OH",229 West 7th St,CyrusOne Cincinatti - 7th Street,45202,CNDC,COGENT,-84.5173,39.1028,2020/5/24 11:50
1305,United States,"Cincinnati, OH",250 E 5th St,250 East Fifth,45202,COB,COGENT,-84.5086,39.1022,2020/5/24 11:50
1306,United States,"Cincinnati, OH",255 E. 5th St,255 E. 5th St,45202,COB,COGENT,-84.5084,39.1016,2020/5/24 11:50
1307,United States,"Cincinnati, OH",302 West Third Street,Immedion (fka. Profitability.net),45202,CNDC,COGENT,-84.5177,39.0984,2020/5/24 11:50
1308,United States,"Cincinnati, OH",312 Walnut Street,312 Walnut Street,45202,COB,COGENT,-84.5104,39.0993,2020/5/24 11:50
1309,United States,"Cincinnati, OH",425 Walnut St,US Bank Tower,45202,COB,COGENT,-81.6214,31.1413,2020/5/24 11:50
1310,United States,"Cincinnati, OH",600 Vine Street,600 Vine Street,45202,COB,COGENT,-84.513,39.1029,2020/5/24 11:50
1311,United States,"Cleveland, OH",100 Public Square (aka 200 W. Prospect Ave),vXchnge,44113,CNDC,COGENT,-81.6948,41.4978,2020/5/24 11:50
1312,United States,"Cleveland, OH",100 Public Square (aka 200 W. Prospect Ave),Higbee Building,44113,COB,COGENT,-81.6948,41.4978,2020/5/24 11:50
1313,United States,"Cleveland, OH",1001 Lakeside Avenue East,North Point Tower,44114,COB,COGENT,-81.6908,41.5067,2020/5/24 11:50
1314,United States,"Cleveland, OH",1100 Superior Ave East,Diamond Building,44114,COB,COGENT,-81.687,41.5026,2020/5/24 11:50
1315,United States,"Cleveland, OH",1111 Superior Ave East,Eaton Center,44114,COB,COGENT,-81.6869,41.5035,2020/5/24 11:50
1316,United States,"Cleveland, OH",1228 Euclid Avenue,Halle Building,44115,COB,COGENT,-81.6841,41.5002,2020/5/24 11:50
1317,United States,"Cleveland, OH",1255 Euclid Ave,DataBank CLE1 (fka. 365 Data Center),44115,CNDC,COGENT,-81.6837,41.5008,2020/5/24 11:50
1318,United States,"Cleveland, OH",1255 Euclid Ave,BlueBridge,44115,CNDC,COGENT,-81.6837,41.5008,2020/5/24 11:50
1319,United States,"Cleveland, OH",127 Public Square,Key Tower,44114,COB,COGENT,-81.6942,41.5014,2020/5/24 11:50
1320,United States,"Cleveland, OH",1300 E 9th St,Penton Media Building,44114,COB,COGENT,-81.6909,41.5042,2020/5/24 11:50
1321,United States,"Cleveland, OH",1301 E 9th St,Tower at Erieview,44114,COB,COGENT,-115.364,32.8186,2020/5/24 11:50
1322,United States,"Cleveland, OH",1375 E 9th St,One Cleveland Center,44114,COB,COGENT,-81.6885,41.5038,2020/5/24 11:50
1323,United States,"Cleveland, OH",1525 Rockwell Ave,zColo,44114,CNDC,COGENT,-81.6835,41.5064,2020/5/24 11:50
1324,United States,"Cleveland, OH",1525 Rockwell Ave,SecureData 365,44114,CNDC,COGENT,-81.6835,41.5064,2020/5/24 11:50
1325,United States,"Cleveland, OH",1621 Euclid Ave,Cogent Data Center - Cleveland,44115,CDC,COGENT,-81.6805,41.5017,2020/5/24 11:50
1326,United States,"Cleveland, OH",1621 Euclid Avenue,BF Keith Building,44115,COB,COGENT,-81.6805,41.5017,2020/5/24 11:50
1327,United States,"Cleveland, OH",1621 Euclid Avenue Suite 730,Fusion (fka. Fidelity),44115,CNDC,COGENT,-81.6805,41.5017,2020/5/24 11:50
1328,United States,"CLEVELAND, OH",1625 Rockwell AVE,Cleveland Technology Center,44114,CNDC,COGENT,-81.683,41.5067,2020/5/24 11:50
1329,United States,"Cleveland, OH",1660 W 2nd St,Skylight Office Tower,44113,COB,COGENT,-81.6935,41.4972,2020/5/24 11:50
1330,United States,"Cleveland, OH",200 Public Square,200 Public Square,44114,COB,COGENT,-81.6919,41.5,2020/5/24 11:50
1331,United States,"Cleveland, OH",50 Public Square,Terminal Tower,44113,COB,COGENT,-80.5723,41.6078,2020/5/24 11:50
1332,United States,"Cleveland, OH",55 Public Square,55 Public Square,44113,COB,COGENT,-81.696,41.5002,2020/5/24 11:50
1333,United States,"Cleveland, OH",600 Superior Ave East,Fifth Third Center,44114,COB,COGENT,-81.6896,41.5014,2020/5/24 11:50
1334,United States,"Cleveland, OH",777 St. Clair Avenue,The Westin Cleveland Downtown,44114,COB,COGENT,-81.5835,41.5506,2020/5/24 11:50
1335,United States,"Cleveland, OH",925 Euclid Avenue,925 Euclid Avenue,44115,COB,COGENT,-81.6865,41.501,2020/5/24 11:50
1336,United States,"Clifton, NJ",100 Delawanna,Digital Realty NJR2 (fka. Telx MMR),7014,CNDC,COGENT,-74.1267,40.8299,2020/5/24 11:50
1337,United States,"Clifton, NJ",2 Peekay Drive,Digital Realty (fka. Telx MMR),7014,CNDC,COGENT,-74.1252,40.8309,2020/5/24 11:50
1338,United States,"Collegeville, PA",1000 Black Rock Road,Flexential fka GlaxoSmithKline Data Center,19426,CNDC,COGENT,-76.937,39.7755,2020/5/24 11:50
1339,United States,"Colorado Springs, CO",102 S Tejon Street,Data102,80903,CNDC,COGENT,-104.824,38.8321,2020/5/24 11:50
1340,United States,"Columbus, OH",10 W Broad St,One Columbus Center,43215,COB,COGENT,-83.0012,39.9625,2020/5/24 11:50
1341,United States,"Columbus, OH",100 East Broad Street,Chase Bank Building,43215,COB,COGENT,-82.9975,39.9631,2020/5/24 11:50
1342,United States,"Columbus, OH",175 South Third Street,175 on the Park,43215,COB,COGENT,-90.0505,35.1432,2020/5/24 11:50
1343,United States,"Columbus, OH",21 East State Street,Fifth Third Center,43215,COB,COGENT,-82.9995,39.9597,2020/5/24 11:50
1347,United States,"Columbus, OH",251 Neilston St,zColo,43215,CNDC,COGENT,-82.9949,39.9681,2020/5/24 11:50
1348,United States,"Columbus, OH",41 South High Street,Huntington Center,43215,COB,COGENT,-83.001,39.9611,2020/5/24 11:50
1349,United States,"Columbus, OH",535 Scherers Court,Cologix COL2,43085,CNDC,COGENT,-83.0024,40.1158,2020/5/24 11:50
1350,United States,"Columbus, OH",555 Scherers Court,Cologix COL1,43085,CNDC,COGENT,-83.0024,40.1158,2020/5/24 11:50
1351,United States,"Columbus, OH",585 Scherers Court,Colgix- COL3,43085,CNDC,COGENT,-83.0021,40.1158,2020/5/24 11:50
1352,United States,"Columbus, OH",65 East State Street,Capitol Square,43215,COB,COGENT,-82.9982,39.9599,2020/5/24 11:50
1353,United States,"Columbus, OH",88 East Broad Street,Key Bank Building,43215,COB,COGENT,-82.9983,39.9628,2020/5/24 11:50
1354,United States,"Culpeper, VA",18155 Technology Drive - Bldg. A,Equinix CU1 (fka. NAP Capital - A),22701,CNDC,COGENT,-77.9783,38.4549,2020/5/24 11:50
1355,United States,"Culpeper, VA",18155 Technology Drive - Bldg. B,Equinix CU2 (fka. NAP Capital - B),22701,CNDC,COGENT,-77.9783,38.4549,2020/5/24 11:50
1356,United States,"Culpeper, VA",18155 Technology Drive - Bldg. C,Equinix CU3 (fka. NAP Capital - C),22701,CNDC,COGENT,-77.9783,38.4549,2020/5/24 11:50
1357,United States,"Culpeper, VA",18155 Technology Drive- Bldg. D,Equinix CU4,22701,CNDC,COGENT,-77.9783,38.4549,2020/5/24 11:50
1358,United States,"Dallas, TX",100 Crescent Court,The Crescent,75201,COB,COGENT,-96.8033,32.7939,2020/5/24 11:50
1359,United States,"Dallas, TX",11830 Webb Chapel Road,Evoque DLLTTX01,75234,CNDC,COGENT,-96.8706,32.9119,2020/5/24 11:50
1360,United States,"Dallas, TX",1201 Elm St,Renaissance Tower,75270,COB,COGENT,-96.8024,32.7812,2020/5/24 11:50
1361,United States,"Dallas, TX",1201 Main St,1201 Main St,75202,COB,COGENT,-96.8019,32.7805,2020/5/24 11:50
1362,United States,"Dallas, TX",12221 Merit Dr,Three Forest Plaza,75251,COB,COGENT,-96.7724,32.9145,2020/5/24 11:50
1363,United States,"Dallas, TX",12222 Merit Drive,Merit Tower - Four Forest,75251,COB,COGENT,-96.7731,32.9185,2020/5/24 11:50
1364,United States,"Dallas, TX",13155 Noel Rd,Three Galleria Tower,75240,COB,COGENT,-96.818,32.9375,2020/5/24 11:50
1365,United States,"Dallas, TX","1333 North Stemmons Freeway, Suite 110",IBM (fka. Softlayer),75207,CNDC,COGENT,-96.8148,32.7911,2020/5/24 11:50
1366,United States,"Dallas, TX",13355 Noel Rd,One Galleria Tower,75240,COB,COGENT,-96.818,32.9375,2020/5/24 11:50
1367,United States,"Dallas, TX",1341 W Mockingbird Ln,Mockingbird Towers,75247,COB,COGENT,-96.8658,32.8221,2020/5/24 11:50
1368,United States,"Dallas, TX",13455 Noel Rd,13455 Noel Rd,75240,COB,COGENT,-96.818,32.9375,2020/5/24 11:50
1369,United States,"Dallas, TX",14001 N Dallas Pky,Stanford Corporate Center,75240,COB,COGENT,-96.8231,32.9363,2020/5/24 11:50
1370,United States,"Dallas, TX",1420 Mockingbird Ln W,One Mockingbird Plaza,75247,COB,COGENT,-96.8643,32.8206,2020/5/24 11:50
1371,United States,"Dallas, TX",1445 Ross Ave,Fountain Place,75202,COB,COGENT,-96.8027,32.7848,2020/5/24 11:50
1372,United States,"Dallas, TX",1501 LBJ Fwy,Park West I,75234,COB,COGENT,-96.9628,32.9179,2020/5/24 11:50
1373,United States,"Dallas, TX",1505-1507 LBJ Fwy,Park West II,75234,COB,COGENT,-96.9267,32.91,2020/5/24 11:50
1374,United States,"Dallas, TX",1515 Round Table Drive,Carrier-1 Data Center,75247,CNDC,COGENT,-96.8784,32.8379,2020/5/24 11:50
1375,United States,"Dallas, TX",1601 Elm St,Thanksgiving Tower,75201,COB,COGENT,-96.7982,32.7819,2020/5/24 11:50
1376,United States,"Dallas, TX",1700 Pacific Ave,1700 Pacific Ave,75201,COB,COGENT,-96.7962,32.7825,2020/5/24 11:50
1377,United States,"Dallas, TX",1717 Main St,Comerica Bank Tower,75201,COB,COGENT,-96.7968,32.7816,2020/5/24 11:50
1378,United States,"Dallas, TX",1950 N Stemmons Fwy,Equinix DA2,75207,CNDC,COGENT,-96.8195,32.8003,2020/5/24 11:50
1379,United States,"Dallas, TX",1950 N Stemmons Fwy,Equinix DA6,75207,CNDC,COGENT,-96.8195,32.8003,2020/5/24 11:50
1380,United States,"Dallas, TX",1950 N Stemmons Fwy,zColo (fka. CoreXchange),75207,CNDC,COGENT,-96.8195,32.8003,2020/5/24 11:50
1381,United States,"Dallas, TX",1950 N Stemmons Fwy,Equinix AT3 (fka. S&D/PAIX),75207,CNDC,COGENT,-96.8195,32.8003,2020/5/24 11:50
1382,United States,"Dallas, TX",1950 N Stemmons Fwy,Equinix DA2,75207,CNDC,COGENT,-96.8195,32.8003,2020/5/24 11:50
1383,United States,"Dallas, TX",1950 N Stemmons Fwy,Infomart,75207,COB,COGENT,-96.8195,32.8003,2020/5/24 11:50
1385,United States,"Dallas, TX",1950 N Stemmons Fwy Suite 1032,Cologix DAL1,75207,CNDC,COGENT,-96.8195,32.8009,2020/5/24 11:50
1386,United States,"Dallas, TX",1950 N Stemmons Fwy Suite 2033,Flexential Downtown (fka. Peak 10 / ViaWest),75207,CNDC,COGENT,-96.8195,32.8003,2020/5/24 11:50
1387,United States,"Dallas, TX",1999 Bryan St,Harwood Center,75201,COB,COGENT,-96.7973,32.7847,2020/5/24 11:50
1388,United States,"Dallas, TX",200 Crescent Court,Crescent Court,75201,COB,COGENT,-96.8039,32.7937,2020/5/24 11:50
1389,United States,"Dallas, TX",2001 Bryan St,Bryan Tower,75201,COB,COGENT,-96.7963,32.7854,2020/5/24 11:50
1390,United States,"Dallas, TX",2001 Ross Ave,Trammell Crow Center,75201,COB,COGENT,-81.4145,31.0374,2020/5/24 11:50
1391,United States,"Dallas, TX",2020 Live Oak St,Telecom Carrier Hotel Bldg,75201,CNDC,COGENT,-96.7947,32.7837,2020/5/24 11:50
1392,United States,"Dallas, TX",2021 McKinney Ave,McKinney & Olive,75201,COB,COGENT,-96.8036,32.7913,2020/5/24 11:50
1394,United States,"Dallas, TX",2101 Cedar Springs Rd,Rosewood Court,75201,COB,COGENT,-96.8063,32.7941,2020/5/24 11:50
1395,United States,"Dallas, TX",2200 Ross Ave,Chase Tower,75201,COB,COGENT,-96.7965,32.788,2020/5/24 11:50
1396,United States,"Dallas, TX","2323 Bryan St, Suite 2440",Digital Realty (fka. Telx MMR),75201,CNDC,COGENT,-96.7942,32.7872,2020/5/24 11:50
1397,United States,"Dallas, TX",2323 Bryan Street,Equinix DA4,75201,CNDC,COGENT,-96.7942,32.7872,2020/5/24 11:50
1398,United States,"Dallas, TX",2323 Ross Ave,KPMG Plaza,75201,COB,COGENT,-96.7972,32.789,2020/5/24 11:50
1399,United States,"Dallas, TX",2501 N. Harwood St,St Anne Court,75201,COB,COGENT,-96.8058,32.7922,2020/5/24 11:50
1400,United States,"Dallas, TX",2711 Haskell Ave N,Cityplace Building,75204,COB,COGENT,-96.7918,32.8065,2020/5/24 11:50
1401,United States,"Dallas, TX",2711 LBJ Fwy,The Meridian,75234,COB,COGENT,-96.8875,32.9107,2020/5/24 11:50
1402,United States,"Dallas, TX",2777 Stemmons Fwy,Trinity Towers,75207,COB,COGENT,-96.8496,32.8098,2020/5/24 11:50
1403,United States,"Dallas, TX",2914 Taylor St,Vergent (fka. OPUS-3),75226,CNDC,COGENT,-96.7809,32.7814,2020/5/24 11:50
1404,United States,"Dallas, TX",300 Crescent Court,Crescent Court,75201,COB,COGENT,-96.8041,32.7935,2020/5/24 11:50
1406,United States,"Dallas, TX",3010 LBJ Fwy,Graystone Centre,75234,COB,COGENT,-96.8777,32.911,2020/5/24 11:50
1407,United States,"Dallas, TX",3030 LBJ Fwy,Club Corp Building,75234,COB,COGENT,-96.8758,32.9113,2020/5/24 11:50
1408,United States,"Dallas, TX",3102 Oak Lawn Ave,Centrum Plaza,75219,COB,COGENT,-96.8076,32.8087,2020/5/24 11:50
1409,United States,"Dallas, TX",3180 Irving Blvd - 1st Floor,CenturyLink (fka. Level 3),75247,CNDC,COGENT,-96.8684,32.8068,2020/5/24 11:50
1410,United States,"Dallas, TX",3232 McKinney Ave,One McKinney Plaza,75204,COB,COGENT,-96.799,32.8035,2020/5/24 11:50
1411,United States,"Dallas, TX",325 N Saint Paul St,Tower II,75201,COB,COGENT,-96.7973,32.7833,2020/5/24 11:50
1412,United States,"Dallas, TX",3333 Lee Pky,3333 Lee Pkwy,75219,COB,COGENT,-96.8053,32.8094,2020/5/24 11:50
1413,United States,"Dallas, TX",350 N Saint Paul St,One Dallas Centre,75201,COB,COGENT,-96.7969,32.784,2020/5/24 11:50
1414,United States,"Dallas, TX",3500 Maple Ave,Reverchon Plaza,75219,COB,COGENT,-96.8116,32.8024,2020/5/24 11:50
1415,United States,"Dallas, TX",400 Crescent Court,Rosewood Crescent Hotel,75201,COB,COGENT,-96.8045,32.7946,2020/5/24 11:50
1416,United States,"Dallas, TX",400 S Akard St,Cogent Data Center - Dallas,75202,CDC,COGENT,-96.7987,32.7782,2020/5/24 11:50
1417,United States,"Dallas, TX",400 S Akard St,Databank (fka. LMC MMR),75202,CNDC,COGENT,-96.7987,32.7782,2020/5/24 11:50
1418,United States,"Dallas, TX","400 S Akard St, 2nd Floor",DataBank DFW1,75202,CNDC,COGENT,-96.7987,32.7782,2020/5/24 11:50
1419,United States,"Dallas, TX",4144 N Central Expy,Uptown Tower,75204,COB,COGENT,-96.789,32.814,2020/5/24 11:50
1420,United States,"Dallas, TX",500 Crescent Court,Crescent Court - The Atrium,75201,COB,COGENT,-96.8047,32.7951,2020/5/24 11:50
1421,United States,"Dallas, TX",500 N Akard St,Lincoln Plaza,75201,COB,COGENT,-96.8002,32.7845,2020/5/24 11:50
1422,United States,"Dallas, TX",5001 Spring Valley Rd,Providence Towers,75244,COB,COGENT,-96.8236,32.94,2020/5/24 11:50
1423,United States,"Dallas, TX",5050 Quorum Dr,5050 Quorum Dr,75240,COB,COGENT,-96.8241,32.9483,2020/5/24 11:50
1424,United States,"Dallas, TX",5949 Sherry Ln,Sterling Plaza,75225,COB,COGENT,-96.8104,32.8627,2020/5/24 11:50
1425,United States,"Dallas, TX",600 N Pearl St,South Tower,75201,COB,COGENT,-96.7951,32.7866,2020/5/24 11:50
1426,United States,"Dallas, TX",650 N Pearl St,Dallas Marriott City Center,75201,COB,COGENT,-96.7957,32.7869,2020/5/24 11:50
1427,United States,"Dallas, TX",700 N Pearl St,North Tower,75201,COB,COGENT,-96.7953,32.7872,2020/5/24 11:50
1428,United States,"Dallas, TX",717 N Harwood St,KPMG Centre,75201,COB,COGENT,-96.7982,32.7856,2020/5/24 11:50
1429,United States,"Dallas, TX",750 N St Paul,Saint Paul Place,75201,COB,COGENT,-96.7997,32.786,2020/5/24 11:50
1430,United States,"Dallas, TX",7557 Rambler Rd,7557 Rambler Rd,75231,COB,COGENT,-96.7632,32.8852,2020/5/24 11:50
1431,United States,"Dallas, TX",8144 Walnut Hill Lane,Walnut Glen Tower,75231,COB,COGENT,-96.7662,32.8809,2020/5/24 11:50
1432,United States,"Dallas, TX",8435 Stemmons Fwy N,Digital Realty (fka. Telx),75247,CNDC,COGENT,-96.8737,32.8275,2020/5/24 11:50
1433,United States,"Dallas, TX",8600 Harry Hines Blvd,zColo (fka. CoreXchange),75235,CNDC,COGENT,-96.8651,32.8386,2020/5/24 11:50
1434,United States,"Dallas, TX",8750 N Central Expy,Northpark Central I,75231,COB,COGENT,-96.769,32.8671,2020/5/24 11:50
1435,United States,"Dallas, TX",901 Main St,Bank of America Plaza,75202,COB,COGENT,-88.0039,44.5139,2020/5/24 11:50
1436,United States,"Dayton, OH",130 W 2nd St,DataYard (fka. Donet),45402,CNDC,COGENT,-112.863,32.3801,2020/5/24 11:50
1437,United States,"Dayton, OH",33 West First Street,DataYard (fka. Donet),45402,CNDC,COGENT,-73.5752,40.6415,2020/5/24 11:50
1438,United States,"Denver, CO",1001 17th Street,1001 17th Street,80202,COB,COGENT,-104.994,39.7489,2020/5/24 11:50
1439,United States,"Denver, CO",1050 17th St,Independence Plaza,80265,COB,COGENT,-77.0388,38.9034,2020/5/24 11:50
1440,United States,"Denver, CO",1099 18th St,Granite Tower,80202,COB,COGENT,-104.993,39.7497,2020/5/24 11:50
1441,United States,"Denver, CO",1125 17th St,1125 17th St,80202,COB,COGENT,-104.994,39.7499,2020/5/24 11:50
1442,United States,"Denver, CO",1144 15th,1144 Fifteenth,80202,COB,COGENT,-104.997,39.7473,2020/5/24 11:50
1443,United States,"Denver, CO",1200 17th St,Tabor Center One,80202,COB,COGENT,-78.4054,40.5198,2020/5/24 11:50
1444,United States,"Denver, CO",1225 17th St,17th Street Plaza,80202,COB,COGENT,-104.995,39.7503,2020/5/24 11:50
1445,United States,"Denver, CO",1401 17th St,Alamo Plaza,80202,COB,COGENT,-104.763,40.6266,2020/5/24 11:50
1446,United States,"Denver, CO",1500 Champa St,1500 Champa St,80202,CNDC,COGENT,-104.995,39.7458,2020/5/24 11:50
1447,United States,"Denver, CO",1500 Champa St,Flexential Downtown (fka. Peak 10 / ViaWest),80202,CNDC,COGENT,-104.995,39.7458,2020/5/24 11:50
1448,United States,"Denver, CO",1515 Arapahoe St,Park Central,80202,COB,COGENT,-104.996,39.7481,2020/5/24 11:50
1449,United States,"Denver, CO",1560 Broadway,Denver Post Tower,80202,COB,COGENT,-73.9846,40.7587,2020/5/24 11:50
1450,United States,"Denver, CO",1600 Broadway,Colorado State Bank,80202,COB,COGENT,-73.9844,40.7602,2020/5/24 11:50
1451,United States,"Denver, CO",1625 Broadway,Denver Energy Center,80202,COB,COGENT,-104.988,39.7423,2020/5/24 11:50
1452,United States,"Denver, CO",1660 Lincoln St,Lincoln Center,80202,COB,COGENT,-104.986,39.743,2020/5/24 11:50
1453,United States,"Denver, CO",1670 Broadway,1670 Broadway,80202,COB,COGENT,-104.987,39.7429,2020/5/24 11:50
1454,United States,"Denver, CO",1672 Lawrence St,1672 Lawrence St,80202,COB,COGENT,-104.996,39.7489,2020/5/24 11:50
1455,United States,"Denver, CO",1675 Broadway,World Trade Center,80202,COB,COGENT,-73.9835,40.7631,2020/5/24 11:50
1456,United States,"Denver, CO",1700 Broadway,Mile High Center,80290,COB,COGENT,-73.9823,40.7635,2020/5/24 11:50
1457,United States,"Denver, CO",1700 Lincoln St,Wells Fargo Center,80203,COB,COGENT,-104.985,39.7436,2020/5/24 11:50
1458,United States,"Denver, CO",1750 Welton Street,Grand Hyatt Denver,80202,,COGENT,-104.989,39.7456,2020/5/24 11:50
1459,United States,"Denver, CO",1775 Sherman St,Denver Financial Center I,80203,COB,COGENT,-104.985,39.7445,2020/5/24 11:50
1460,United States,"Denver, CO",1776 Lincoln Street,Denver Financial Center II,80203,COB,COGENT,-104.986,39.7445,2020/5/24 11:50
1461,United States,"Denver, CO",1777 S Harrison,Denver Centerpoint II,80210,COB,COGENT,-104.942,39.6841,2020/5/24 11:50
1462,United States,"Denver, CO",1801 California St,1801 California St,80202,COB,COGENT,-104.989,39.748,2020/5/24 11:50
1464,United States,"Denver, CO",1999 Broadway,1999 Broadway,80202,COB,COGENT,-104.988,39.7478,2020/5/24 11:50
1465,United States,"Denver, CO",303 16th St,Republic Plaza Retail,80202,COB,COGENT,-104.989,39.743,2020/5/24 11:50
1466,United States,"Denver, CO",303 E 17th Ave,17th & Grant Building,80203,COB,COGENT,-104.983,39.7435,2020/5/24 11:50
1467,United States,"Denver, CO",370 17th St,Republic Plaza,80202,COB,COGENT,-104.989,39.7434,2020/5/24 11:50
1468,United States,"Denver, CO",3773 Cherry Creek Dr N,Ptarmigan Place,80209,COB,COGENT,-104.942,39.7089,2020/5/24 11:50
1469,United States,"Denver, CO",3900 E Mexico Ave,Denver Centerpoint,80210,COB,COGENT,-104.941,39.6855,2020/5/24 11:50
1470,United States,"Denver, CO",410 17th St,410 Bldg,80202,COB,COGENT,-104.989,39.744,2020/5/24 11:50
1471,United States,"Denver, CO",4300 Brighton Blvd,Iron Mountain (fka. ForTrust Data Center),80216,CNDC,COGENT,-104.968,39.7774,2020/5/24 11:50
1472,United States,"Denver, CO",4600 S Syracuse St,Hines Development Bldg,80237,COB,COGENT,-104.899,39.6292,2020/5/24 11:50
1473,United States,"Denver, CO",4600 S Ulster St,Metropoint 1,80237,COB,COGENT,-104.896,39.6318,2020/5/24 11:50
1474,United States,"Denver, CO",4610 S Ulster St,Metropoint II,80237,COB,COGENT,-104.895,39.6312,2020/5/24 11:50
1475,United States,"Denver, CO",4643 S Ulster St,Cogent Data Center - Denver,80237,CDC,COGENT,-104.897,39.6306,2020/5/24 11:50
1476,United States,"Denver, CO",4643 S Ulster St,Regency Plaza One,80237,COB,COGENT,-104.897,39.6306,2020/5/24 11:50
1477,United States,"Denver, CO",555 17th St,555 17th St,80202,COB,COGENT,-104.989,39.7454,2020/5/24 11:50
1478,United States,"Denver, CO",600 17th St,Dominion Plaza,80202,COB,COGENT,-104.991,39.7454,2020/5/24 11:50
1479,United States,"Denver, CO",621 17th St,Colorado Plaza Tower ll,80293,COB,COGENT,-104.99,39.7464,2020/5/24 11:50
1480,United States,"Denver, CO",633 17th St,Colorado Plaza Tower One,80202,COB,COGENT,-104.99,39.7466,2020/5/24 11:50
1481,United States,"Denver, CO",639 E 18th Ave,CoreSite DE2 (fka. Comfluent),80203,CNDC,COGENT,-104.979,39.7453,2020/5/24 11:50
1482,United States,"Denver, CO",707 17th St,707 17th Street,80202,COB,COGENT,-104.99,39.7471,2020/5/24 11:50
1483,United States,"Denver, CO",717 17th St,Johns Manville Plaza,80202,COB,COGENT,-104.991,39.7469,2020/5/24 11:50
1484,United States,"Denver, CO",7800 E Union Ave,Denver Corporate Center 2,80237,COB,COGENT,-104.898,39.6276,2020/5/24 11:50
1485,United States,"Denver, CO",7900 E Union Ave,Denver Corporate Center 3,80237,COB,COGENT,-104.896,39.6277,2020/5/24 11:50
1486,United States,"Denver, CO",900 South Broadway 400,Hosting.com Denver,80209,CNDC,COGENT,-90.187,38.6327,2020/5/24 11:50
1487,United States,"Denver, CO",910 15th St,910 Telecom,80202,CNDC,COGENT,-104.996,39.7456,2020/5/24 11:50
1488,United States,"Denver, CO",910 15th St,CoreSite DE1,80202,CNDC,COGENT,-104.996,39.7456,2020/5/24 11:50
1490,United States,"Denver, CO",910 15th St - Suite 751,CoreSite DE1 (fka. Comfluent),80202,CNDC,COGENT,-104.996,39.7456,2020/5/24 11:50
1491,United States,"Denver, CO",950 17th St,U.S. Bank,80202,COB,COGENT,-104.994,39.7476,2020/5/24 11:50
1492,United States,"Denver, CO",999 18th St,Denver Place-North & South Towers,80202,COB,COGENT,-104.991,39.7495,2020/5/24 11:50
1493,United States,"Des Moines, IA",317 6th Ave,BofA Building,50309,COB,COGENT,-97.2283,38.4858,2020/5/24 11:50
1494,United States,"Des Moines, IA",400 Locust Street,Capital Square,50309,COB,COGENT,-96.6171,32.6332,2020/5/24 11:50
1495,United States,"Des Moines, IA",666 Grand Ave,Ruan Center,50309,COB,COGENT,-93.6261,41.5872,2020/5/24 11:50
1496,United States,"Des Moines, IA",666 Walnut Street,Connect Des Moines (fka. Fiberlink),50309,CNDC,COGENT,-93.6254,41.5855,2020/5/24 11:50
1497,United States,"Des Moines, IA",666 Walnut Street,Financial Center,50309,COB,COGENT,-93.6254,41.5855,2020/5/24 11:50
1498,United States,"Des Moines, IA",801 Grand Ave,The Principal Building,50309,COB,COGENT,-93.6286,41.5875,2020/5/24 11:50
1499,United States,"Detroit, MI",100 Renaissance Center,Renaissance Center,48243,COB,COGENT,-83.0398,42.3293,2020/5/24 11:50
1500,United States,"Detroit, MI",150 W Jefferson Ave,150 W Jefferson Ave,48226,COB,COGENT,-83.0465,42.3282,2020/5/24 11:50
1501,United States,"Detroit, MI",200 Renaissance Center,200 Renaissance Ctr,482423,COB,COGENT,-83.0398,42.3293,2020/5/24 11:50
1502,United States,"Detroit, MI",300 River Place Dr,Stroh River Place,48207,COB,COGENT,-83.0174,42.3366,2020/5/24 11:50
1503,United States,"Detroit, MI",3011 West Grand Blvd,Fisher Building,48202,COB,COGENT,-83.0776,42.3699,2020/5/24 11:50
1504,United States,"Detroit, MI",3031 W Grand Blvd,DTW05 - New Center One,48202,COB,COGENT,-83.0761,42.37,2020/5/24 11:50
1505,United States,"Detroit, MI",400 Renaissance Center,400 Renaissance Ctr,48243,COB,COGENT,-83.0398,42.3293,2020/5/24 11:50
1506,United States,"DORAL, FL",1525 NW 98TH CT,Equinix MI6 (Fka Verizon),33172,CNDC,COGENT,-80.3546,25.7878,2020/5/24 11:50
1507,United States,"Doral, FL",3701 NW 82nd Avenue,Cogent Data Center - Doral,33166,CDC,COGENT,-80.3287,25.8079,2020/5/24 11:50
1508,United States,"Doral, FL",8180 NW 36th St,8180 NW 36th St,33166,COB,COGENT,-80.3288,25.8091,2020/5/24 11:50
1509,United States,"Eagan, MN",3255 Neil Armstrong Blvd,DataBank MSP2,55121,CNDC,COGENT,-93.1458,44.8385,2020/5/24 11:50
1510,United States,"Edina, MN",7700 France Ave S,DataBank MSP1,55435,CNDC,COGENT,-93.3306,44.8627,2020/5/24 11:50
1511,United States,"Edina, MN",7700 France Ave S,7700 France Ave,55435,COB,COGENT,-93.3306,44.8627,2020/5/24 11:50
1512,United States,"Edison, NJ",3003 Woodbridge Ave,Iron Mountain (fka. IO Data Center),8837,CNDC,COGENT,-74.3474,40.5211,2020/5/24 11:50
1513,United States,"El Paso, TX",500 W Overland Avenue,Transtelco Data Center,79901,CNDC,COGENT,-106.493,31.7553,2020/5/24 11:50
1514,United States,"El Segundo, CA",100 N Pacific Coast Hwy (fka 100 N Sepulveda Blvd),Pacific Corporate Towers Bldg A,90245,COB,COGENT,-118.395,33.917,2020/5/24 11:50
1515,United States,"El Segundo, CA",1920 E Maple Avenue,Equinix LA3,90245,CNDC,COGENT,-118.394,33.9263,2020/5/24 11:50
1516,United States,"El Segundo, CA",200 N Nash,Cyxtera LA1,90245,CNDC,COGENT,-118.386,33.9178,2020/5/24 11:50
1517,United States,"El Segundo, CA",200 N Pacific Coast Hwy (fka 200 N Sepulveda Blvd),Pacific Corporate Towers Bldg B,90245,COB,COGENT,-118.396,33.9236,2020/5/24 11:50
1518,United States,"El Segundo, CA",222 N Pacific Coast Hwy (fka 222 N Sepulveda Blvd),Pacific Corporate Towers Bldg C,90245,COB,COGENT,-118.395,33.9188,2020/5/24 11:50
1519,United States,"El Segundo, CA",2260 E El Segundo Blvd,Digital Realty,90245,CNDC,COGENT,-118.385,33.9157,2020/5/24 11:50
1520,United States,"El Segundo, CA",444 N. Nash,T5 Data Centers,90245,CNDC,COGENT,-118.387,33.9216,2020/5/24 11:50
1521,United States,"El Segundo, CA",445 N Douglas St,Equinix LA4,90245,CNDC,COGENT,-118.384,33.9212,2020/5/24 11:50
1522,United States,"Elk Grove Village, IL",1441 Touhy Ave,T5 Chicago (fka. Forsythe),60007,CNDC,COGENT,-87.9644,42.0067,2020/5/24 11:50
1523,United States,"Elk Grove Village, IL",1905 Lunt Avenue,Equinix CH3,60007,CNDC,COGENT,-87.9549,42.0011,2020/5/24 11:50
1524,United States,"Elk Grove Village, IL",2200 Busse Road,Digital Realty (fka. DuPont Fabros CH1),60007,CNDC,COGENT,-87.962,41.9943,2020/5/24 11:50
1525,United States,"Elk Grove Village, IL",2200 Busse Road,Rackspace,60007,CNDC,COGENT,-87.962,41.9943,2020/5/24 11:50
1526,United States,"Elk Grove Village, IL",2200 Busse Road,Server Central,60007,CNDC,COGENT,-87.962,41.9943,2020/5/24 11:50
1527,United States,"Elk Grove Village, IL",2299 Busse Rd.,Digital Realty (fka. DuPont Fabros),60007,CNDC,COGENT,-87.9577,41.9942,2020/5/24 11:50
1528,United States,"Elk Grove Village, IL",2425 Busse Road,Cyxtera ORD2 (fka. CenturyLink CH3),60007,CNDC,COGENT,-87.9583,41.9923,2020/5/24 11:50
1529,United States,"Elmsford, NY",401 Fieldcrest Drive,Equinix NY13 (Fka Verizon),10523,CNDC,COGENT,-73.8199,41.0708,2020/5/24 11:50
1530,United States,"Emeryville, CA",1400 65th Street,Evocative (fka. 365 Data Centers),94608,CNDC,COGENT,-122.293,37.8467,2020/5/24 11:50
1531,United States,"Emeryville, CA",5000 Hollis St,CenturyLink (fka. Level 3),94608,CNDC,COGENT,-122.287,37.8353,2020/5/24 11:50
1532,United States,"Englewood, CO",335 Inverness Drive South,Equinix DE2 (fka. Verizon),80112,CNDC,COGENT,-97.2991,27.8837,2020/5/24 11:50
1533,United States,"Englewood, CO",383 Inverness Parkway,Plaza at Inverness,80112,COB,COGENT,-104.867,39.5622,2020/5/24 11:50
1534,United States,"Englewood, CO",385 Inverness Parkway,Plaza at Inverness,80112,COB,COGENT,-104.867,39.5631,2020/5/24 11:50
1535,United States,"Englewood, CO",391 Inverness Pkwy,zColo (fka. Latisys),80112,CNDC,COGENT,-104.867,39.564,2020/5/24 11:50
1536,United States,"Englewood, CO",393 Inverness Pkwy,zColo (fka. Latisys),80112,CNDC,COGENT,-104.867,39.5641,2020/5/24 11:50
1537,United States,"Englewood, CO",8534 Concord Center Drive,Cyxtera DN3,80112,CNDC,COGENT,-104.831,39.5609,2020/5/24 11:50
1539,United States,"Englewood, CO",8636 S Peoria St,Flexential Englewood (fka. Peak 10 / ViaWest),80112,CNDC,COGENT,-104.838,39.559,2020/5/24 11:50
1540,United States,"Englewood, CO",9706 E Easter Ave,Equinix DE1 (fka. S&D),80112,CNDC,COGENT,-104.876,39.5873,2020/5/24 11:50
1541,United States,"Farmers Branch, TX",1503 LBJ Fwy,Park West I,75234,COB,COGENT,-96.9279,32.91,2020/5/24 11:50
1542,United States,"Farmers Branch, TX",4055 Valley View Ln,Granite Tower,75244,COB,COGENT,-96.8437,32.9261,2020/5/24 11:50
1543,United States,"Federal Way, WA",32275 32nd Avenue South,Utility Computing,98001,,COGENT,-122.295,47.3118,2020/5/24 11:50
1544,United States,"Federal Way, WA",32275 32nd Avenue South,Cogent Data Center - Seattle,98001,CDC,COGENT,-122.295,47.3118,2020/5/24 11:50
1545,United States,"Fort Lauderdale, FL",1 E Broward Blvd,Well Fargo Tower,33301,COB,COGENT,-80.1431,26.1228,2020/5/24 11:50
1546,United States,"Fort Lauderdale, FL",5301 NW 33rd Avenue,"Flexential (fka. Peak 10 / ViaWest - FLL 1,2)",33309,CNDC,COGENT,-80.1933,26.1903,2020/5/24 11:50
1547,United States,"Fort Worth, TX",100 Throckmorton St,Two City Place,76102,COB,COGENT,-97.3346,32.7557,2020/5/24 11:50
1548,United States,"Fort Worth, TX",14901 FAA Blvd-Bldg. 1,Cyxtera DL1,76155,CNDC,COGENT,-97.0408,32.8328,2020/5/24 11:50
1549,United States,"Fort Worth, TX",14901 FAA Boulevard- Building 2,Cyxtera DL2,76155,CNDC,COGENT,-97.0408,32.8328,2020/5/24 11:50
1550,United States,"Fort Worth, TX",201 Main St,Wells Fargo Tower,76102,COB,COGENT,-75.332,43.4825,2020/5/24 11:50
1551,United States,"Fort Worth, TX",300 Throckmorton St,One City Place,76102,COB,COGENT,-97.3338,32.7545,2020/5/24 11:50
1552,United States,"Fort Worth, TX",301 Commerce St,DR Horton Tower,76102,COB,COGENT,-97.3305,32.756,2020/5/24 11:50
1553,United States,"Fort Worth, TX",307 W 7th Street,Commerce Bldg,76102,COB,COGENT,-118.254,34.0455,2020/5/24 11:50
1554,United States,"Fort Worth, TX",309 W 7th St,Oil & Gas Building,76102,COB,COGENT,-97.4574,39.3456,2020/5/24 11:50
1555,United States,"Fort Worth, TX",777 Main St,Carter Burgess Tower,76102,COB,COGENT,-72.6738,41.7662,2020/5/24 11:50
1556,United States,"Fort Worth, TX",801 Cherry St,Burnett Plaza,76102,COB,COGENT,-97.3346,32.7506,2020/5/24 11:50
1558,United States,"Franklin, TN",311 Eddy Lane,Tierpoint-Nashville DC200 & 300,37054,CNDC,COGENT,-86.8558,35.9245,2020/5/24 11:50
1559,United States,"Franklin, TN",425 Duke Drive,Flexential Cool Springs (fka. Peak 10 / ViaWest -,37067,CNDC,COGENT,-86.8333,35.9511,2020/5/24 11:50
1560,United States,"Franklin, TN",4600 Carothers Parkway,Flexential Franklin (fka. Peak 10 / ViaWest - Nas,37067,CNDC,COGENT,-86.8156,35.9203,2020/5/24 11:50
1561,United States,"Franklin Park, IL",9333 W Grand Ave,Digital Realty,60131,CNDC,COGENT,-87.858,41.927,2020/5/24 11:50
1562,United States,"Franklin Park, IL",9333 W Grand Ave,Tierpoint CHI (West),60131,CNDC,COGENT,-87.858,41.927,2020/5/24 11:50
1563,United States,"Franklin Park, IL",9355 Grand Avenue,Digital Realty,60131,CNDC,COGENT,-87.8592,41.9277,2020/5/24 11:50
1564,United States,"Franklin Park, IL",9377 Grand Avenue,Digital Realty,60131,CNDC,COGENT,-87.8588,41.9286,2020/5/24 11:50
1565,United States,"Franklin Township, NJ",200 Campus Dr,Rackspace (fka. Datapipe NJ),8873,CNDC,COGENT,-76.6734,40.2629,2020/5/24 11:50
1566,United States,"Garland, TX",2008 Lookout Drive,RagingWire TX1,75044,CNDC,COGENT,-96.6578,32.9811,2020/5/24 11:50
1567,United States,"Greensboro, NC",301 S Elm Street,Spectrum (fka. TWC/DukeNet),27401,CNDC,COGENT,-96.8496,32.3847,2020/5/24 11:50
1568,United States,"Greenwood Village, CO",5350 S Valentia Way,H5 Denver,80111,CNDC,COGENT,-104.892,39.6181,2020/5/24 11:50
1570,United States,"Hamilton, OH",5307 Muhlhauser Road,Flexential CIN01-02,45011,CNDC,COGENT,-84.4489,39.3158,2020/5/24 11:50
1571,United States,"Hawthorne, NY",11 Skyline Drive,Tierpoint -HWT (Hawthorne),10532,CNDC,COGENT,-73.8148,41.089,2020/5/24 11:50
1572,United States,"Hawthorne, NY",17 Skyline Drive,Tierpoint-HW2,10532,CNDC,COGENT,-73.4388,41.4945,2020/5/24 11:50
1573,United States,"Herndon, VA",510 Huntmar Park Dr (Hosting Ctr),Cogent Data Center - Herndon,20170,CDC,COGENT,-77.3803,38.9627,2020/5/24 11:50
1574,United States,"Highlands Ranch, CO",9110 Commerce Center Circle,Cyxtera DN2,80129,CNDC,COGENT,-105.034,39.5502,2020/5/24 11:50
1575,United States,"Highlands Ranch, CO",9180 Commerce Center Circle,Cyxtera DN1,80129,CNDC,COGENT,-105.033,39.5499,2020/5/24 11:50
1576,United States,"Hillsboro, OR",3825 NW Aloclek Place,Digital Realty PRT1 (fka. Telx),97124,CNDC,COGENT,-122.892,45.5476,2020/5/24 11:50
1577,United States,"Hillsboro, OR",3825 NW Aloclek Place,Digital Realty,97124,CNDC,COGENT,-122.892,45.5476,2020/5/24 11:50
1578,United States,"Hillsboro, OR",3935 NW Aloclek Place,Flexential Hillsboro (fka. Peak 10 / ViaWest),97124,CNDC,COGENT,-122.891,45.5485,2020/5/24 11:50
1580,United States,"Hillsboro, OR",6327 NW Evergreen Parkway,EdgeConneX EDCPOR01,97124,CNDC,COGENT,-122.915,45.5517,2020/5/24 11:50
1581,United States,"Hillsboro, OR",8135 NE Evergreen Pkwy (fka 21515 NW Evergreen Pkw,Infomart Portland,97124,CNDC,COGENT,-105.35,39.6573,2020/5/24 11:50
1582,United States,"Houston, TX",1 Greenway Plz,1 Greenway Plz,77046,COB,COGENT,-95.4326,29.7328,2020/5/24 11:50
1583,United States,"Houston, TX",1 Riverway,One Riverway,77056,COB,COGENT,-95.4628,29.7611,2020/5/24 11:50
1584,United States,"Houston, TX",1000 Louisiana St,Wells Fargo Plaza,77002,COB,COGENT,-95.3684,29.7585,2020/5/24 11:50
1585,United States,"Houston, TX",1000 Main St,Reliant Energy Plaza,77002,COB,COGENT,-95.3659,29.7568,2020/5/24 11:50
1586,United States,"Houston, TX",1001 Fannin St,1001 Fannin,77002,COB,COGENT,-95.3639,29.7558,2020/5/24 11:50
1587,United States,"Houston, TX",1001 McKinney St,1001 McKinney St,77002,COB,COGENT,-95.364,29.7571,2020/5/24 11:50
1588,United States,"Houston, TX","1001 Texas Ave, Suite 330",Quasar Data Center (fka. Lakota),77002,CNDC,COGENT,-95.3621,29.7601,2020/5/24 11:50
1589,United States,"Houston, TX",1021 Main St,One City Centre,77002,COB,COGENT,-92.4884,35.6796,2020/5/24 11:50
1590,United States,"Houston, TX",11 Greenway Plz,Summit Tower,77046,COB,COGENT,-95.4354,29.7321,2020/5/24 11:50
1591,United States,"Houston, TX",1100 Louisiana St,1100 Louisiana St,77002,COB,COGENT,-95.3689,29.7576,2020/5/24 11:50
1592,United States,"Houston, TX",11003 Corporate Centre Drive,CyrusOne Houston West III,77041,CNDC,COGENT,-95.5609,29.8439,2020/5/24 11:50
1593,United States,"Houston, TX",1111 Bagby St,Heritage Plaza,77002,COB,COGENT,-95.3707,29.7588,2020/5/24 11:50
1594,United States,"Houston, TX",12 Greenway Plz,Summit Plaza West,77046,COB,COGENT,-95.4356,29.7331,2020/5/24 11:50
1595,United States,"Houston, TX",1200 Smith St,Two Allen Center,77002,COB,COGENT,-95.371,29.757,2020/5/24 11:50
1596,United States,"Houston, TX",12001 North Fwy,CenturyLink (fka. Level 3),77060,CNDC,COGENT,-95.4168,29.9432,2020/5/24 11:50
1597,United States,"Houston, TX",1201 Louisiana St,Total Plaza,77002,COB,COGENT,-90.4523,29.8401,2020/5/24 11:50
1598,United States,"Houston, TX",12031 North Freeway,Fibertown - Houston Data Center,77060,CNDC,COGENT,-95.4164,29.9431,2020/5/24 11:50
1599,United States,"Houston, TX",12061 North Fwy,Cogent Data Center - Houston,77060,CDC,COGENT,-95.4164,29.9441,2020/5/24 11:50
1600,United States,"Houston, TX",12175 North Fwy,Sungard - Houston N2,77060,CNDC,COGENT,-95.4158,29.9425,2020/5/24 11:50
1601,United States,"Houston, TX",1221 McKinney St,Houston Center 1,77010,COB,COGENT,-95.3624,29.7562,2020/5/24 11:50
1602,United States,"Houston, TX",1221/1331 Lamar St,Four Houston Center,77010,COB,COGENT,-95.3623,29.7546,2020/5/24 11:50
1603,United States,"Houston, TX",1300 Post Oak Blvd,Wells Fargo Bank Tower,77056,COB,COGENT,-95.4603,29.7532,2020/5/24 11:50
1604,United States,"Houston, TX",1301 Fannin St,Netrality,77002,COB,COGENT,-95.3657,29.7534,2020/5/24 11:50
1605,United States,"Houston, TX",1301 Fannin St 1100,Internap,77002,CNDC,COGENT,-95.3658,29.7538,2020/5/24 11:50
1606,United States,"Houston, TX",1301 Fannin St - Level 3,CenturyLink (fka. Level 3 / Looking Glass),77002,CNDC,COGENT,-95.3657,29.7534,2020/5/24 11:50
1607,United States,"Houston, TX",1301 Fannin Street,Netrality,77002,CNDC,COGENT,-95.3657,29.7534,2020/5/24 11:50
1608,United States,"Houston, TX",1301 McKinney St,Fulbright Tower,77010,COB,COGENT,-95.3615,29.7555,2020/5/24 11:50
1609,United States,"Houston, TX",1330 Post Oak Blvd,Union Texas Center,77056,COB,COGENT,-95.4612,29.7534,2020/5/24 11:50
1610,United States,"Houston, TX",1360 Post Oak Blvd,BHP Tower,77056,COB,COGENT,-95.4621,29.7531,2020/5/24 11:50
1611,United States,"Houston, TX",1400 Post Oak Blvd,Interfin Bldg,77056,COB,COGENT,-95.4625,29.7521,2020/5/24 11:50
1613,United States,"Houston, TX",1415 Louisiana St,Wedge International Tower,77002,COB,COGENT,-95.3699,29.7548,2020/5/24 11:50
1614,United States,"Houston, TX",1515 Aldine Meadows Road,Equinix HO1 (Fka Verizon),77032,CNDC,COGENT,-95.3643,29.9358,2020/5/24 11:50
1615,United States,"Houston, TX",1600 Smith St,Continental Center I,77002,COB,COGENT,-95.3727,29.7545,2020/5/24 11:50
1616,United States,"Houston, TX",16825 Northchase Dr,Two Greenspoint Plaza,77060,COB,COGENT,-95.4036,29.943,2020/5/24 11:50
1617,United States,"Houston, TX",16945 Northchase Dr,Four Greenspoint Plaza,77060,COB,COGENT,-95.4036,29.943,2020/5/24 11:50
1618,United States,"Houston, TX",1800 West Loop S,US Home Bldg,77027,COB,COGENT,-95.4572,29.7491,2020/5/24 11:50
1619,United States,"Houston, TX",1900 West Loop S,3D/International Tower,77027,COB,COGENT,-95.4583,29.7473,2020/5/24 11:50
1620,United States,"Houston, TX",1980 Post Oak Blvd,Two Post Oak Central,77056,COB,COGENT,-95.4626,29.7462,2020/5/24 11:50
1621,United States,"Houston, TX",1990 Post Oak Blvd,Three Post Oak Central,77056,COB,COGENT,-95.4636,29.7455,2020/5/24 11:50
1622,United States,"Houston, TX",2 Riverway,Two Riverway,77056,COB,COGENT,-95.4625,29.7618,2020/5/24 11:50
1623,United States,"Houston, TX",2000 Post Oak Blvd,One Post Oak Central,77056,COB,COGENT,-90.0049,29.8512,2020/5/24 11:50
1624,United States,"Houston, TX",2000 West Loop S,2000 West Loop S,77027,COB,COGENT,-95.4579,29.7466,2020/5/24 11:50
1625,United States,"Houston, TX",24 Greenway Plz,Weslayan Tower,77046,COB,COGENT,-95.4405,29.73,2020/5/24 11:50
1626,United States,"Houston, TX",2500 CityWest Blvd,2500 CityWest Blvd,77042,COB,COGENT,-95.5615,29.7383,2020/5/24 11:50
1627,United States,"Houston, TX",2700 Post Oak Blvd,Galleria Tower I,77056,COB,COGENT,-95.4643,29.739,2020/5/24 11:50
1628,United States,"Houston, TX",2800 North Loop W,Brookhollow Central I,77092,COB,COGENT,-95.447,29.8059,2020/5/24 11:50
1629,United States,"Houston, TX",2800 Post Oak Blvd,Williams Tower,77056,COB,COGENT,-95.4622,29.7371,2020/5/24 11:50
1630,United States,"Houston, TX",2900 North Loop W.,Brookhollow Central II,77092,COB,COGENT,-95.4478,29.8059,2020/5/24 11:50
1631,United States,"Houston, TX",2929 Allen Pky,America Tower,77019,COB,COGENT,-95.3975,29.7607,2020/5/24 11:50
1632,United States,"Houston, TX",2950 North Loop W.,Brookhollow Central III,77092,COB,COGENT,-95.4486,29.8048,2020/5/24 11:50
1633,United States,"Houston, TX",3 Greenway Plaza,Three Greenway Plaza,77046,COB,COGENT,-95.4313,29.7316,2020/5/24 11:50
1634,United States,"Houston, TX",3 Riverway,Three Riverway,77056,COB,COGENT,-95.4616,29.7606,2020/5/24 11:50
1635,United States,"Houston, TX",3040 Post Oak Blvd,Thirty Forty Post Oak,77056,COB,COGENT,-95.4628,29.735,2020/5/24 11:50
1636,United States,"Houston, TX",3050 Post Oak Blvd,3050 Post Oak Blvd,77056,COB,COGENT,-95.4628,29.7339,2020/5/24 11:50
1637,United States,"Houston, TX",3200 Southwest Fwy,The Phoenix Tower,77027,COB,COGENT,-95.4289,29.731,2020/5/24 11:50
1638,United States,"Houston, TX",333 Clay St,333 Clay St,77002,COB,COGENT,-103.792,43.8102,2020/5/24 11:50
1639,United States,"Houston, TX",3700 Southwest Freeway,Lakewood Church,77027,COB,COGENT,-95.4347,29.7303,2020/5/24 11:50
1640,United States,"Houston, TX",3800 Buffalo Speedway,3800 Buffalo Speedway,77098,COB,COGENT,-95.4294,29.7317,2020/5/24 11:50
1641,United States,"Houston, TX",4201 Southwest Frwy,CyrusOne Houston - Galleria,77027,CNDC,COGENT,-95.4451,29.7281,2020/5/24 11:50
1642,United States,"Houston, TX",4211A Southwest Frwy,CyrusOne Houston - Galleria,77027,CNDC,COGENT,-95.4457,29.7281,2020/5/24 11:50
1643,United States,"Houston, TX",4211B Southwest Frwy,CyrusOne Houston - Galleria,77027,CNDC,COGENT,-95.4457,29.7281,2020/5/24 11:50
1644,United States,"Houston, TX",4400 Post Oak Pky,5 Post Oak Park,77027,COB,COGENT,-95.4517,29.7477,2020/5/24 11:50
1645,United States,"Houston, TX",4801 Woodway Dr,4801 Woodway Dr,77056,COB,COGENT,-95.4602,29.7636,2020/5/24 11:50
1646,United States,"Houston, TX",5 Greenway Plz,Greenway Plaza 5,77046,COB,COGENT,-95.4394,29.7312,2020/5/24 11:50
1647,United States,"Houston, TX",500 Dallas St,One Allen Center,77002,COB,COGENT,-88.2181,41.9397,2020/5/24 11:50
1648,United States,"Houston, TX",5051 Westheimer Rd,Galleria Tower II,77056,COB,COGENT,-95.4626,29.7389,2020/5/24 11:50
1649,United States,"Houston, TX",5060 W Alabama,Westin Hotel,77027,COB,COGENT,-95.4646,29.7384,2020/5/24 11:50
1650,United States,"Houston, TX",5065-5075 Westheimer Rd,Galleria Financial Center,77056,COB,COGENT,-95.4634,29.7392,2020/5/24 11:50
1651,United States,"Houston, TX",5150 Westway Park Blvd,CyrusOne Houston West I,77041,CNDC,COGENT,-95.5566,29.8442,2020/5/24 11:50
1652,United States,"Houston, TX",5151 San Felipe St,Sage One Plaza,77056,COB,COGENT,-95.4648,29.7497,2020/5/24 11:50
1653,United States,"Houston, TX",5170 Westway Park Blvd,CyrusOne Houston West II,77041,CNDC,COGENT,-95.5564,29.8443,2020/5/24 11:50
1654,United States,"Houston, TX",5444 Westheimer Rd,CMS Energy Tower,77056,COB,COGENT,-95.4736,29.739,2020/5/24 11:50
1655,United States,"Houston, TX",5555 San Felipe St,Marathon Oil Tower,77056,COB,COGENT,-95.4716,29.7494,2020/5/24 11:50
1656,United States,"Houston, TX",5599 San Felipe St,GeoQuest Center,77056,COB,COGENT,-95.474,29.7496,2020/5/24 11:50
1657,United States,"Houston, TX",5718 Westheimer Rd,Capital One Plaza,77057,COB,COGENT,-95.4789,29.7385,2020/5/24 11:50
1658,United States,"Houston, TX",5847 San Felipe St,San Felipe Plaza,77057,COB,COGENT,-95.482,29.7495,2020/5/24 11:50
1659,United States,"Houston, TX",600 Travis St,JP Morgan Chase Tower,77002,COB,COGENT,-95.3641,29.7605,2020/5/24 11:50
1660,United States,"Houston, TX",601 Jefferson St,Kellogg Brown & Root Tower,77002,COB,COGENT,-89.8498,43.5455,2020/5/24 11:50
1661,United States,"Houston, TX",609 Main Street,609 Main Street,77002,COB,COGENT,-95.3625,29.7593,2020/5/24 11:50
1662,United States,"Houston, TX",660 Greens Parkway,Data Foundry Houston 2,77067,CNDC,COGENT,-95.4242,29.9452,2020/5/24 11:50
1663,United States,"Houston, TX",700 Louisiana St,Bank of America Center,77002,COB,COGENT,-95.3662,29.7606,2020/5/24 11:50
1664,United States,"Houston, TX",700 Milam St,Pennzoil Place - North Tower,77002,COB,COGENT,-95.3656,29.7601,2020/5/24 11:50
1665,United States,"Houston, TX",711 Louisiana Street,Pennzoil Place - South Tower,77002,COB,COGENT,-95.3658,29.7604,2020/5/24 11:50
1666,United States,"Houston, TX",712 Main St,JP Morgan Chase Building,77002,COB,COGENT,-74.0416,40.1706,2020/5/24 11:50
1667,United States,"Houston, TX",717 Texas Avenue,Calpine Center,77002,COB,COGENT,-95.3642,29.7616,2020/5/24 11:50
1668,United States,"Houston, TX",737 N Eldridge Parkway,Eldridge Place III,77079,COB,COGENT,-95.6194,29.7762,2020/5/24 11:50
1669,United States,"Houston, TX",757 N Eldridge Parkway,Two Eldridge Place,77079,COB,COGENT,-95.6194,29.7771,2020/5/24 11:50
1670,United States,"Houston, TX",777 N Eldridge Parkway,One Eldridge Place,77079,COB,COGENT,-95.6188,29.7776,2020/5/24 11:50
1671,United States,"Houston, TX",8 Greenway Plz,Greenway Plaza 8,77046,COB,COGENT,-95.4316,29.7315,2020/5/24 11:50
1672,United States,"Houston, TX",800 Capitol St,Capitol Tower,77002,COB,COGENT,-95.3648,29.7594,2020/5/24 11:50
1673,United States,"Houston, TX",800 Gessner Dr,One Memorial City Plaza,77024,COB,COGENT,-95.5439,29.7777,2020/5/24 11:50
1674,United States,"Houston, TX",801 Travis St,Wells Fargo Center,77002,COB,COGENT,-95.3646,29.7585,2020/5/24 11:50
1675,United States,"Houston, TX",808 Travis St,Niels Esperson Building,77002,COB,COGENT,-95.3652,29.759,2020/5/24 11:50
1676,United States,"Houston, TX",811 Louisiana Street aka 777 Walker St,Two Shell Plaza,77002,COB,COGENT,-95.3666,29.7593,2020/5/24 11:50
1677,United States,"Houston, TX",811 Main Street,BGP PLACE,77002,COB,COGENT,-95.3633,29.7577,2020/5/24 11:50
1678,United States,"Houston, TX",815 Walker,Mellie Esperson Building,77002,COB,COGENT,-95.3653,29.7585,2020/5/24 11:50
1679,United States,"Houston, TX",820 Gessner Dr,Two Memorial City Plaza,77024,COB,COGENT,-95.5433,29.7776,2020/5/24 11:50
1680,United States,"Houston, TX",840 Gessner Dr,Three Memorial City Plaza,77024,COB,COGENT,-95.5426,29.7777,2020/5/24 11:50
1681,United States,"Houston, TX",9 Greenway Plz,Greenway Plaza 9,77046,COB,COGENT,-95.4345,29.7319,2020/5/24 11:50
1682,United States,"Houston, TX",909 Fannin St,Two Houston Center,77010,COB,COGENT,-95.3633,29.7564,2020/5/24 11:50
1683,United States,"Houston, TX",910 Louisiana St,One Shell Plaza,77002,COB,COGENT,-95.3677,29.7591,2020/5/24 11:50
1684,United States,"Houston, TX",919 Milam St (aka 910 Travis St),Bank One Center,77002,COB,COGENT,-95.3659,29.7582,2020/5/24 11:50
1685,United States,"Houston, TX",9821 Katy Fwy,Memorial City Place,77024,COB,COGENT,-95.5374,29.7835,2020/5/24 11:50
1686,United States,"Indianapolis, IN",10 W. Market Street,Market Tower,46204,COB,COGENT,-81.5188,41.0856,2020/5/24 11:50
1688,United States,"Indianapolis, IN",101 W Ohio St,101 West Ohio,46204,COB,COGENT,-86.1602,39.7697,2020/5/24 11:50
1689,United States,"Indianapolis, IN",101-115 W. Washington Street,National City Center,46204,COB,COGENT,-86.1604,39.7663,2020/5/24 11:50
1690,United States,"Indianapolis, IN",111 Monument Circle (aka 1 E. Ohio),Chase Tower,46204,COB,COGENT,-86.157,39.769,2020/5/24 11:50
1691,United States,"Indianapolis, IN",135 N Pennsylvania St,135 N Pennsylvania St,46204,COB,COGENT,-86.1558,39.7695,2020/5/24 11:50
1692,United States,"Indianapolis, IN",151 W Ohio Street,151 W Ohio Street,46204,COB,COGENT,-86.1609,39.7697,2020/5/24 11:50
1693,United States,"Indianapolis, IN",201 N. Illinois St,Captial Center South,46204,COB,COGENT,-97.1417,39.6323,2020/5/24 11:50
1694,United States,"Indianapolis, IN",211 N Pennsylvania (fka 1 Indiana Square),Regions Tower,46204,COB,COGENT,-86.155,39.7705,2020/5/24 11:50
1695,United States,"Indianapolis, IN",251 East Ohio,251 East Ohio,46204,COB,COGENT,-86.1529,39.7694,2020/5/24 11:50
1696,United States,"Indianapolis, IN",251 N. Illinois St,Capital Center North Tower,46204,COB,COGENT,-86.1601,39.7629,2020/5/24 11:50
1697,United States,"Indianapolis, IN",300 N. Meridian St,300 N. Meridian St,46204,COB,COGENT,-86.1584,39.7716,2020/5/24 11:50
1698,United States,"Indianapolis, IN",401 N Shadeland Ave,Lifeline Data Centers,46219,CNDC,COGENT,-86.0427,39.776,2020/5/24 11:50
1699,United States,"Indianapolis, IN",505 W. Merrill St.,Online Tech,46225,CNDC,COGENT,-86.1678,39.7587,2020/5/24 11:50
1700,United States,"Indianapolis, IN",701 W Henry St,365 Data Centers (fka. Equinix - S&D),46225,CNDC,COGENT,-86.1708,39.7596,2020/5/24 11:50
1701,United States,"Indianapolis, IN",733 W Henry Street,Lifeline Data Centers,46225,CNDC,COGENT,-86.1724,39.7597,2020/5/24 11:50
1702,United States,"Indianapolis, IN",800 Oliver Ave,Paragon Colocation,46225,CNDC,COGENT,-117.254,32.7919,2020/5/24 11:50
1703,United States,"Indianapolis,, IN",151 N Delaware St,Market Square,46204,COB,COGENT,-86.154,39.7694,2020/5/24 11:50
1704,United States,"Irvine, CA",1 Park Plaza,Jamboree Center-1 Park Bldg.,92614,COB,COGENT,-117.84,33.6797,2020/5/24 11:50
1705,United States,"Irvine, CA",17222 Von Karman,zColo (fka. Latisys),92614,CNDC,COGENT,-117.85,33.6815,2020/5/24 11:50
1706,United States,"Irvine, CA",17400 Von Karman Avenue,zColo (fka. Latisys),92614,CNDC,COGENT,-117.842,33.6879,2020/5/24 11:50
1707,United States,"Irvine, CA",17836 Gillette Ave,Cyxtera OC2,92614,CNDC,COGENT,-117.853,33.6906,2020/5/24 11:50
1708,United States,"Irvine, CA",18200 Von Karman Ave,Irvine Towers - 18200 Bldg.,92612,COB,COGENT,-117.851,33.6777,2020/5/24 11:50
1709,United States,"Irvine, CA",18300 Von Karman Ave,Irvine Towers- 18300 Bldg. Union Bank,92612,COB,COGENT,-117.851,33.6769,2020/5/24 11:50
1710,United States,"Irvine, CA",18400 Von Karman Ave,Irvine Towers -18400 Bldg.,92612,COB,COGENT,-117.852,33.677,2020/5/24 11:50
1711,United States,"Irvine, CA",18500 Von Karman Ave,Irvine Towers Bldg 5,92612,COB,COGENT,-117.852,33.6763,2020/5/24 11:50
1712,United States,"Irvine, CA",1920 Main Street,1920 Main Street,92612,COB,COGENT,-117.853,33.6841,2020/5/24 11:50
1713,United States,"Irvine, CA",19800 MacArthur,Macarthur Tower 1,92612,COB,COGENT,-117.86,33.6553,2020/5/24 11:50
1714,United States,"Irvine, CA",19900 MacArthur Boulevard,Macarthur Blvd.-Tower 2,92612,COB,COGENT,-117.86,33.656,2020/5/24 11:50
1715,United States,"Irvine, CA",2 Park Plaza,Jamboree Center-2 Park Plaza,92614,COB,COGENT,-117.838,33.6794,2020/5/24 11:50
1716,United States,"Irvine, CA",2010 Main St,2010 Main Plaza,92614,COB,COGENT,-117.851,33.6843,2020/5/24 11:50
1717,United States,"Irvine, CA",2030 Main St.,Irvine Concourse,92614,COB,COGENT,-117.853,33.6832,2020/5/24 11:50
1718,United States,"Irvine, CA",2040 Main Street,Irvine Concourse Bldg 2,92614,COB,COGENT,-117.852,33.6829,2020/5/24 11:50
1719,United States,"Irvine, CA",2050 Main Street,Irvine Concourse -Phase III,92614,COB,COGENT,-117.852,33.6829,2020/5/24 11:50
1720,United States,"Irvine, CA",2525 Main Street,Grupo SMS Datacenter,92614,CNDC,COGENT,-117.866,33.7705,2020/5/24 11:50
1721,United States,"Irvine, CA",2600 Michelson Drive,2600 Michelson,92612,COB,COGENT,-117.85,33.6737,2020/5/24 11:50
1722,United States,"Irvine, CA",2640 Main Street,Alchemy Irvine,92612,CNDC,COGENT,-117.842,33.6811,2020/5/24 11:50
1723,United States,"Irvine, CA",2681 Kelvin Ave,Evoque IRVNCAOW,92614,CNDC,COGENT,-117.837,33.6847,2020/5/24 11:50
1724,United States,"Irvine, CA",3 Park Plaza,Jamboree Center-3 Park Bldg,92614,COB,COGENT,-117.839,33.6769,2020/5/24 11:50
1725,United States,"Irvine, CA",4 Park Plaza,Jamboree Center,92614,COB,COGENT,-117.837,33.6783,2020/5/24 11:50
1726,United States,"Irvine, CA",5 Park Plaza,Jamboree Center-5 Park Bldg.,92614,COB,COGENT,-117.837,33.6767,2020/5/24 11:50
1727,United States,"Irving, TX",122 W John Carpenter Fwy,122 W Carpenter Fwy,75039,COB,COGENT,-96.9462,32.8701,2020/5/24 11:50
1728,United States,"Irving, TX",125 E John Carpenter Fwy,Caltex House,75062,COB,COGENT,-96.9408,32.8679,2020/5/24 11:50
1729,United States,"Irving, TX",222 W Las Colinas Blvd,Millennium Center,75039,COB,COGENT,-96.9434,32.8711,2020/5/24 11:50
1730,United States,"Irving, TX",2222 E Grauwyler Rd,Equinix DA9 (Fka: Verizon),75061,CNDC,COGENT,-96.9127,32.8303,2020/5/24 11:50
1731,United States,"Irving, TX",400 East Las Colinas Blvd,Canal Plaza,75039,COB,COGENT,-96.9356,32.8625,2020/5/24 11:50
1732,United States,"Irving, TX",5215 N OConnor Blvd,Williams Square - Central Tower,75039,COB,COGENT,-96.9388,32.8712,2020/5/24 11:50
1733,United States,"Irving, TX",5221 N OConnor Blvd,Williams Square - East Tower,75039,COB,COGENT,-96.9379,32.8707,2020/5/24 11:50
1734,United States,"Irving, TX",545 E John Carpenter Fwy,545 E John Carpenter Fwy,75062,COB,COGENT,-96.9329,32.8593,2020/5/24 11:50
1735,United States,"Irving, TX",600 E Las Colinas Blvd,600 E Las Colinas Blvd,75039,COB,COGENT,-96.9301,32.8603,2020/5/24 11:50
1736,United States,"Irving, TX",6431 Longhorn Dr,QTS Datacenter,75063,CNDC,COGENT,-96.983,32.8971,2020/5/24 11:50
1737,United States,"Jackson, MS",111 East Capitol Street,111 East Capitol Street,39201,COB,COGENT,-90.188,32.2997,2020/5/24 11:50
1738,United States,"Jacksonville, FL",1 Independent Drive,Wells Fargo Center,32202,COB,COGENT,-81.6563,30.3256,2020/5/24 11:50
1739,United States,"Jacksonville, FL",1301 Riverplace Blvd,Riverplace Tower,32207,COB,COGENT,-81.484,31.1519,2020/5/24 11:50
1740,United States,"Jacksonville, FL",200 W Forsyth Street,200 W Forsyth Street,32202,COB,COGENT,-81.6609,30.3272,2020/5/24 11:50
1741,United States,"Jacksonville, FL",301 West Bay Street,EverBank Center,32202,COB,COGENT,-81.6624,30.3276,2020/5/24 11:50
1742,United States,"Jacksonville, FL",421 West Church Street,Jax Telecom Center,32202,CNDC,COGENT,-81.6625,30.3313,2020/5/24 11:50
1743,United States,"Jacksonville, FL",421 West Church Street,Cologix JAX1,32202,CNDC,COGENT,-81.6625,30.3313,2020/5/24 11:50
1744,United States,"Jacksonville, FL",421 West Church Street,GoRack,32202,CNDC,COGENT,-81.6625,30.3313,2020/5/24 11:50
1745,United States,"Jacksonville, FL",4800 Spring Park Rd,Cologix JAX2,32207,CNDC,COGENT,-81.6139,30.2732,2020/5/24 11:50
1746,United States,"Jacksonville, FL",4905 Belfort Rd.,Flexential JAX01-02,32256,CNDC,COGENT,-81.5819,30.2457,2020/5/24 11:50
1747,United States,"Jacksonville, FL",50 North Laura Street,Bank of America Tower,32202,COB,COGENT,-81.6601,30.3271,2020/5/24 11:50
1748,United States,"JACKSONVILLE, FL",6602 Executive Park Court 105,Edgeconnex EDCJAX01,32216,CNDC,COGENT,-81.6041,30.2543,2020/5/24 11:50
1749,United States,"Jacksonville, FL",76 South Laura Street,Suntrust Tower,32202,COB,COGENT,-81.6599,30.3259,2020/5/24 11:50
1750,United States,"Jacksonville, FL",8324 Baymeadows Way,Tierpoint JAX,32256,CNDC,COGENT,-81.5819,30.2172,2020/5/24 11:50
1751,United States,"Jacksonville, FL",841 Prudential Drive,Aetna,32207,COB,COGENT,-81.6634,30.3161,2020/5/24 11:50
1752,United States,"Jersey City, NJ",10 2nd Street / 3 Second Street,Harborside Plaza 10,7302,COB,COGENT,-74.0322,40.7209,2020/5/24 11:50
1753,United States,"Jersey City, NJ",10 Exchange Place,10 Exchange Place,7302,COB,COGENT,-74.0334,40.7164,2020/5/24 11:50
1754,United States,"Jersey City, NJ",101 Hudson St,101 Hudson St,7302,COB,COGENT,-74.0351,40.7161,2020/5/24 11:50
1757,United States,"Jersey City, NJ","185 Hudson St( fka 34 Exchange Place, Plaza V","Harborside Financial Center, Plaza V",7302,COB,COGENT,-74.0344,40.719,2020/5/24 11:50
1758,United States,"Jersey City, NJ",200 Hudson Street / 34 Exchange Place,"Harborside Financial Center, Plaza II",7311,COB,COGENT,-74.033,40.7187,2020/5/24 11:50
1759,United States,"Jersey City, NJ",210 Hudson St / 34 Exchange Place,"Harborside Financial Center, Plaza III",7311,COB,COGENT,-74.0331,40.7192,2020/5/24 11:50
1761,United States,"Jersey City, NJ",525 Washington Blvd,Newport Tower,7310,COB,COGENT,-74.0356,40.7277,2020/5/24 11:50
1762,United States,"Jersey City, NJ",95 Christopher Columbus Dr,International Financial Tower,7302,COB,COGENT,-74.042,40.7188,2020/5/24 11:50
1763,United States,"Jersey City, NJ",95 Christopher Columbus Dr 19th Flr,QTS Datacenter,7302,CNDC,COGENT,-74.0422,40.7191,2020/5/24 11:50
1764,United States,"Kansas City, MO",1100 Main St,City Center Square,64105,COB,COGENT,-94.5836,39.1007,2020/5/24 11:50
1765,United States,"Kansas City, MO",1100 Walnut St,Town Pavilion,64106,COB,COGENT,-75.1592,39.9484,2020/5/24 11:50
1766,United States,"Kansas City, MO",1102 Grand Blvd,1102 Grand,64106,COB,COGENT,-94.5811,39.1008,2020/5/24 11:50
1767,United States,"Kansas City, MO",1102 Grand Blvd MMR,Building MMR,64106,CNDC,COGENT,-94.5811,39.1008,2020/5/24 11:50
1768,United States,"Kansas City, MO",2345 Grand Blvd,2345 Tower,64108,COB,COGENT,-94.5809,39.0852,2020/5/24 11:50
1769,United States,"Kansas City, MO",324 E 11th St,Netsolus,64106,CNDC,COGENT,-91.5587,41.6843,2020/5/24 11:50
1772,United States,"Kent, WA",6906 South 204th Street,Equinix SE4 (fka. Verizon),98032,CNDC,COGENT,-122.248,47.4206,2020/5/24 11:50
1773,United States,"La Jolla, CA",4250 Executive Square,Executive Square,92037,COB,COGENT,-117.215,32.8739,2020/5/24 11:50
1774,United States,"La Jolla, CA",4275 Executive Square,Executive Square,92037,COB,COGENT,-117.215,32.8735,2020/5/24 11:50
1775,United States,"La Mirada, CA",16680 Valley View Avenue,Cogent Data Center - La Mirada,90638,CDC,COGENT,-118.027,33.8798,2020/5/24 11:50
1776,United States,"Laredo, TX",520 Matamoros,VTX,78040,CNDC,COGENT,-99.5014,27.5065,2020/5/24 11:50
1777,United States,"Las Vegas, NV",1110 Palms Airport Drive,VegasNAP / Fiberhub,89119,CNDC,COGENT,-115.138,36.066,2020/5/24 11:50
1778,United States,"Las Vegas, NV",1541 Pama Lane,EdgeConneX EDCLAS01,89119,CNDC,COGENT,-115.131,36.0652,2020/5/24 11:50
1779,United States,"Las Vegas, NV",2595 Fremont St,LasVegas.net,89104,CNDC,COGENT,-115.114,36.1575,2020/5/24 11:50
1780,United States,"Las Vegas, NV",302 E Carson Street,Flexential Dowtown (fka. Peak 10 / ViaWest),89101,CNDC,COGENT,-115.144,36.1686,2020/5/24 11:50
1781,United States,"Las Vegas, NV",304 E Carson Avenue,Flexential (fka. Peak 10 / ViaWest / Corelink),89101,CNDC,COGENT,-115.144,36.1677,2020/5/24 11:50
1782,United States,"Las Vegas, NV",3330 E. Lone Mountain Rd,"Flexential North (Golden Triangle Industrial Park,",89081,CNDC,COGENT,-115.101,36.248,2020/5/24 11:50
1783,United States,"Las Vegas, NV",7135 S Decatur Blvd,Swith LAS VEGAS,89118,CNDC,COGENT,-115.21,36.0598,2020/5/24 11:50
1784,United States,"Las Vegas, NV",7185 Pollock Drive,zColo,89119,CNDC,COGENT,-115.146,36.0592,2020/5/24 11:50
1785,United States,"Laurel, MD",312 Laurel Avenue,AiNET,20707,CNDC,COGENT,-71.383,41.8388,2020/5/24 11:50
1786,United States,"Lenexa, KS",11200 Lakeview Avenue,Databank - South Lake Data Center,66219,CNDC,COGENT,-94.7703,38.9252,2020/5/24 11:50
1787,United States,"Lenexa, KS",14500 W 105th St,Tierpoint - LEN (Lenexa),66215,CNDC,COGENT,-94.755,38.9407,2020/5/24 11:50
1788,United States,"Lewis Center, OH",8180 Green Meadows Drive North,Cyxtera CL1,43035,CNDC,COGENT,-83.0108,40.1688,2020/5/24 11:50
1789,United States,"Lewisville, TX",2501 South State Highway 121 BUS,CyrusOne Dallas - Lewisville,75067,CNDC,COGENT,-96.9977,33.0002,2020/5/24 11:50
1790,United States,"Linthicum Heights, MD",813 Pinnacle Drive,Tierpoint - BWI (Baltimore),21090,CNDC,COGENT,-76.6799,39.2138,2020/5/24 11:50
1791,United States,"Lisle, IL",4513 Western Avenue,Evoque LSLEIAA,60532,CNDC,COGENT,-88.0947,41.8024,2020/5/24 11:50
1792,United States,"Lithia Springs, GA",375 Riverside Parkway,Cyxtera ATL1 (fka. CenturyLink AT1),30122,CNDC,COGENT,-84.5815,33.7438,2020/5/24 11:50
1793,United States,"Lombard, IL",1850 Springer Dr.,CyrusOne Chicago - Lombard,60148,CNDC,COGENT,-88.0296,41.8492,2020/5/24 11:50
1794,United States,"Los Angeles, CA",1000 Wilshire Blvd,1000 Wilshire Blvd,90017,COB,COGENT,-118.261,34.0504,2020/5/24 11:50
1795,United States,"Los Angeles, CA",10100 Santa Monica Blvd,10100 Santa Monica Blvd,90067,COB,COGENT,-118.416,34.0617,2020/5/24 11:50
1796,United States,"Los Angeles, CA",10250 Constellation Blvd,MGM Tower,90067,COB,COGENT,-118.417,34.0571,2020/5/24 11:50
1797,United States,"Los Angeles, CA",1055 W 7th St,1055 W 7th St,90017,COB,COGENT,-118.263,34.0508,2020/5/24 11:50
1798,United States,"Los Angeles, CA",10880 Wilshire Blvd,Oppenheimer Plaza,90024,COB,COGENT,-118.443,34.0587,2020/5/24 11:50
1799,United States,"Los Angeles, CA",10960 Wilshire Blvd,Saban Plaza,90024,COB,COGENT,-118.446,34.0578,2020/5/24 11:50
1800,United States,"Los Angeles, CA",11100 Santa Monica Blvd,Westwood Gateway 2 East Tower,90025,COB,COGENT,-118.444,34.0465,2020/5/24 11:50
1801,United States,"Los Angeles, CA",11111 Santa Monica Blvd,Westwood Gateway 1,90025,COB,COGENT,-118.445,34.0484,2020/5/24 11:50
1802,United States,"Los Angeles, CA",11150 Santa Monica Blvd,Westwood Gateway 2 North Tower,90025,COB,COGENT,-118.445,34.0469,2020/5/24 11:50
1803,United States,"Los Angeles, CA",1200 W 7th St,Garland Connect (building MMR),90017,CNDC,COGENT,-118.266,34.0511,2020/5/24 11:50
1804,United States,"Los Angeles, CA",1200 W 7th St - Suite L1-100,Alchemy Downtown LA,90017,CNDC,COGENT,-118.266,34.0511,2020/5/24 11:50
1805,United States,"Los Angeles, CA",1200 W 7th St Lower Level 1,Evocative LA1 (fka. net2ez),90017,CNDC,COGENT,-118.266,34.0511,2020/5/24 11:50
1806,United States,"Los Angeles, CA",12910 Culver Road,12910 Culver Boulevard,90066,COB,COGENT,-118.42,33.9923,2020/5/24 11:50
1807,United States,"Los Angeles, CA",12950 Culver Blvd,AMV Digital Media (fka. AEGDM),90066,CNDC,COGENT,-118.42,33.9923,2020/5/24 11:50
1808,United States,"Los Angeles, CA",1800 Avenue of the Stars,Gateway East Bldg,90067,COB,COGENT,-118.417,34.061,2020/5/24 11:50
1809,United States,"Los Angeles, CA",1800 Century Park E,Northrop Plaza I,90067,COB,COGENT,-118.415,34.0627,2020/5/24 11:50
1810,United States,"Los Angeles, CA",1801 Century Park E,Century Park Plaza,90067,COB,COGENT,-118.416,34.0622,2020/5/24 11:50
1811,United States,"Los Angeles, CA",1840 Century Park E,Northrop Grumman Plaza II,90067,COB,COGENT,-118.415,34.0622,2020/5/24 11:50
1812,United States,"Los Angeles, CA",1875 Century Park E,Watt Plaza - North Tower,90067,COB,COGENT,-118.415,34.061,2020/5/24 11:50
1813,United States,"Los Angeles, CA",1880 Century Park E,1880 Century Park E,90067,COB,COGENT,-118.414,34.0619,2020/5/24 11:50
1814,United States,"Los Angeles, CA",1888 Century Park E,Eighteen Eighty Eight Bldg,90067,COB,COGENT,-118.414,34.061,2020/5/24 11:50
1815,United States,"Los Angeles, CA",1900 Avenue of the Stars,1900 Avenue of the Stars,90067,COB,COGENT,-118.417,34.0603,2020/5/24 11:50
1816,United States,"Los Angeles, CA",1901 Avenue of the Stars,1901 Avenue of the Stars,90067,COB,COGENT,-118.418,34.0596,2020/5/24 11:50
1817,United States,"Los Angeles, CA",1925 Century Park E,Watt Plaza - South Tower,90067,COB,COGENT,-118.415,34.0606,2020/5/24 11:50
1818,United States,"Los Angeles, CA",1999 Avenue of the Stars,SunAmerica Center,90067,COB,COGENT,-118.417,34.0589,2020/5/24 11:50
1819,United States,"Los Angeles, CA",2000 Avenue of the Stars,2000 Avenue of the Stars,90067,COB,COGENT,-118.415,34.0582,2020/5/24 11:50
1820,United States,"Los Angeles, CA",2029 Century Park E,Century Plaza Towers North Tower,90067,COB,COGENT,-118.414,34.0594,2020/5/24 11:50
1821,United States,"Los Angeles, CA",2049 Century Park E,Century Plaza Towers South Tower,90067,COB,COGENT,-118.413,34.0586,2020/5/24 11:50
1822,United States,"Los Angeles, CA",2121 Avenue of the Stars,Fox Plaza,90067,COB,COGENT,-118.413,34.0552,2020/5/24 11:50
1823,United States,"Los Angeles, CA",300 S Grand Ave,One California Plaza,90071,COB,COGENT,-118.251,34.0533,2020/5/24 11:50
1824,United States,"Los Angeles, CA",3200 Wilshire Blvd,Towers On Wilshire-North,90010,COB,COGENT,-117.988,33.8746,2020/5/24 11:50
1825,United States,"Los Angeles, CA",3250 Wilshire Blvd,One Park Plaza,90010,COB,COGENT,-118.293,34.0611,2020/5/24 11:50
1826,United States,"Los Angeles, CA",333 S Grand Ave,"Wells Fargo Center 1, North Tower",90071,COB,COGENT,-118.252,34.0529,2020/5/24 11:50
1827,United States,"Los Angeles, CA",333 S Hope St,Bank of America Plaza,90071,COB,COGENT,-118.253,34.0537,2020/5/24 11:50
1828,United States,"Los Angeles, CA",3333 Wilshire Blvd,Wilshire Square Two,90010,COB,COGENT,-118.296,34.0622,2020/5/24 11:50
1829,United States,"Los Angeles, CA",3435 Wilshire Blvd,Equitable Plaza Bldg,90010,COB,COGENT,-118.298,34.0623,2020/5/24 11:50
1830,United States,"Los Angeles, CA",3440 Wilshire Blvd,Bldg 1,90010,COB,COGENT,-118.299,34.0614,2020/5/24 11:50
1831,United States,"Los Angeles, CA",3450 Wilshire Blvd,Bldg 2,90010,COB,COGENT,-118.3,34.0615,2020/5/24 11:50
1832,United States,"Los Angeles, CA",3460 Wilshire Blvd,Bldg 3,90010,COB,COGENT,-118.3,34.0613,2020/5/24 11:50
1833,United States,"Los Angeles, CA",350 S Grand Ave,Two California Plaza,90071,COB,COGENT,-118.251,34.0522,2020/5/24 11:50
1834,United States,"Los Angeles, CA",355 S Grand Ave,Wells Fargo Center - South Tower,90071,COB,COGENT,-89.2958,43.1273,2020/5/24 11:50
1835,United States,"Los Angeles, CA",3580 Wilshire Blvd,Paramount Plaza,90010,COB,COGENT,-118.303,34.0611,2020/5/24 11:50
1836,United States,"Los Angeles, CA",3700 Wilshire Blvd,Wilshire Park Place,90010,COB,COGENT,-118.307,34.0605,2020/5/24 11:50
1837,United States,"Los Angeles, CA",400 S Hope St,Mellon Bank Center,90071,COB,COGENT,-118.254,34.0516,2020/5/24 11:50
1838,United States,"Los Angeles, CA",444 S Flower St,Citibank Center,90071,COB,COGENT,-118.255,34.0516,2020/5/24 11:50
1839,United States,"Los Angeles, CA",445 S Figueroa St,Union Bank Plaza,90071,COB,COGENT,-118.257,34.0536,2020/5/24 11:50
1840,United States,"Los Angeles, CA",4929 Wilshire Blvd,Wilshire/Highland Bldg,90010,COB,COGENT,-118.338,34.0625,2020/5/24 11:50
1841,United States,"Los Angeles, CA",500 S Grand Ave,Biltmore Tower,90071,COB,COGENT,-118.254,34.0498,2020/5/24 11:50
1842,United States,"Los Angeles, CA",515 S Figueroa St,Manulife Plaza,90071,COB,COGENT,-118.258,34.0521,2020/5/24 11:50
1843,United States,"Los Angeles, CA",515 S Flower St,North Tower,90071,COB,COGENT,-118.257,34.0514,2020/5/24 11:50
1844,United States,"Los Angeles, CA",520 S Grand Ave,Biltmore Court,90071,COB,COGENT,-118.254,34.0497,2020/5/24 11:50
1845,United States,"Los Angeles, CA",530 W 6th St,Cogent Data Center - Los Angeles,90014,CDC,COGENT,-117.638,35.3726,2020/5/24 11:50
1846,United States,"Los Angeles, CA",530 W 6th St,Last Mile IX2,90014,CNDC,COGENT,-117.638,35.3726,2020/5/24 11:50
1847,United States,"Los Angeles, CA",530 W 6th St,Building MMR,90014,CNDC,COGENT,-117.638,35.3726,2020/5/24 11:50
1848,United States,"Los Angeles, CA",530 W 6th St - Penthouse,QuadraNet (fka. MultiPoint),90014,CNDC,COGENT,-117.638,35.3726,2020/5/24 11:50
1849,United States,"Los Angeles, CA",5455 Wilshire Blvd,5455 Wilshire Blvd,90036,COB,COGENT,-118.348,34.0628,2020/5/24 11:50
1850,United States,"Los Angeles, CA",550 S Hope St,550 S Hope St,90071,COB,COGENT,-118.256,34.0497,2020/5/24 11:50
1851,United States,"Los Angeles, CA",555 S Flower St,Bank of America Tower,90071,COB,COGENT,-118.258,34.0508,2020/5/24 11:50
1852,United States,"Los Angeles, CA",555 W 5th St,The Gas Company Tower,90013,COB,COGENT,-118.253,34.0501,2020/5/24 11:50
1853,United States,"Los Angeles, CA",5670 Wilshire Blvd,California Federal,90036,COB,COGENT,-118.352,34.062,2020/5/24 11:50
1854,United States,"Los Angeles, CA",5700 Wilshire Blvd,Wilshire Courtyard East,90036,COB,COGENT,-118.353,34.0618,2020/5/24 11:50
1855,United States,"Los Angeles, CA",5750 Wilshire Blvd,Wilshire Courtyard West,90036,COB,COGENT,-118.355,34.0612,2020/5/24 11:50
1856,United States,"Los Angeles, CA",5757 Wilshire Blvd,Museum Square,90036,COB,COGENT,-118.353,34.0626,2020/5/24 11:50
1857,United States,"Los Angeles, CA",5900 Wilshire Blvd,5900 Wilshire Blvd,90036,COB,COGENT,-118.359,34.0621,2020/5/24 11:50
1858,United States,"Los Angeles, CA","600 W 7th Street, 6th Floor",Equinix LA1,90017,CNDC,COGENT,-118.257,34.0474,2020/5/24 11:50
1859,United States,"Los Angeles, CA","600 W 7th Street, Suite 550",Digital Realty (fka. Telx Building MMR),90017,CNDC,COGENT,-118.257,34.0474,2020/5/24 11:50
1860,United States,"Los Angeles, CA",600 Wilshire Blvd,Wilshire-Grand Bldg,90017,COB,COGENT,-118.257,34.048,2020/5/24 11:50
1861,United States,"Los Angeles, CA",601 S Figueroa St,Sanwa Bank Plaza,90017,COB,COGENT,-118.259,34.0508,2020/5/24 11:50
1862,United States,"Los Angeles, CA",606 S Olive St,City National Bank Bldg,90014,COB,COGENT,-118.254,34.0477,2020/5/24 11:50
1863,United States,"Los Angeles, CA",6060 Center Drive West,6060 Center Dr,90045,COB,COGENT,-118.391,33.976,2020/5/24 11:50
1864,United States,"Los Angeles, CA",6080 Center Drive West,Howard Hughes Center-6080,90045,COB,COGENT,-118.392,33.9768,2020/5/24 11:50
1865,United States,"Los Angeles, CA",609 W 7th Street,Crown Castle (fka. Wilcon/IX2),90017,CNDC,COGENT,-118.257,34.0476,2020/5/24 11:50
1866,United States,"Los Angeles, CA",6100 Center Drive West,Howard Hughes-6100 Bldg,90045,COB,COGENT,-118.392,33.9772,2020/5/24 11:50
1867,United States,"Los Angeles, CA",6171 W Century Boulevard,Alchemy LAX,90045,CNDC,COGENT,-118.394,33.9458,2020/5/24 11:50
1868,United States,"Los Angeles, CA",624 S Grand Ave,CoreSite LA1 (One Wilshire),90017,CNDC,COGENT,-118.256,34.048,2020/5/24 11:50
1869,United States,"Los Angeles, CA",626 Wilshire Boulevard,Telehouse LA Center,90017,CNDC,COGENT,-118.257,34.0483,2020/5/24 11:50
1870,United States,"Los Angeles, CA",633 W 5th St,US Bank Tower,90071,COB,COGENT,-118.254,34.0511,2020/5/24 11:50
1872,United States,"Los Angeles, CA",6601 Center Drive West,Howard Hughes Center- 6601 Bldg.,90045,COB,COGENT,-118.394,33.9801,2020/5/24 11:50
1873,United States,"Los Angeles, CA",6701 Center Drive West,Howard Hughes Tower,90045,COB,COGENT,-118.394,33.979,2020/5/24 11:50
1874,United States,"Los Angeles, CA",700 S Flower St,The Bloc,90017,COB,COGENT,-118.259,34.0483,2020/5/24 11:50
1875,United States,"Los Angeles, CA",700 Wilshire Blvd - 1st Floor,Meet Me Room (Summit Co-Locate),90017,CNDC,COGENT,-118.257,34.0486,2020/5/24 11:50
1876,United States,"Los Angeles, CA",707 Wilshire Blvd,707 Wilshire Blvd,90017,COB,COGENT,-118.257,34.0492,2020/5/24 11:50
1877,United States,"Los Angeles, CA",725 S Figueroa St,Ernst & Young Plaza,90017,COB,COGENT,-118.261,34.0493,2020/5/24 11:50
1878,United States,"Los Angeles, CA",777 S Figueroa St,777 Tower,90017,COB,COGENT,-118.261,34.0484,2020/5/24 11:50
1879,United States,"Los Angeles, CA",800 S Hope St,CoreSite LA4,90017,CNDC,COGENT,-118.259,34.046,2020/5/24 11:50
1880,United States,"Los Angeles, CA",800 W 6th St,Pacific Financial Center,90017,COB,COGENT,-118.258,34.0501,2020/5/24 11:50
1881,United States,"Los Angeles, CA",801 S Figueroa St,801 Tower,90017,COB,COGENT,-118.262,34.0479,2020/5/24 11:50
1882,United States,"Los Angeles, CA",818 W 7th St - 11th Floor,CenturyLink (fka. Level 3),90017,CNDC,COGENT,-118.26,34.0489,2020/5/24 11:50
1883,United States,"Los Angeles, CA",818 W 7th St - 6th Floor,Equinix LA2 (fka. Pihana),90017,CNDC,COGENT,-118.26,34.0489,2020/5/24 11:50
1884,United States,"Los Angeles, CA",865 S Figueroa St,TCW Bldg,90017,COB,COGENT,-118.263,34.0469,2020/5/24 11:50
1885,United States,"Los Angeles, CA",900 N Alameda,CoreSite LA2 (Wilshire Annex),90012,CNDC,COGENT,-118.235,34.0579,2020/5/24 11:50
1886,United States,"Los Angeles, CA",900 Wilshire Blvd,Wilshire Grand Centre,90017,COB,COGENT,-118.493,34.023,2020/5/24 11:50
1887,United States,"Louisville, KY",101 South Fifth Street,National City Tower,40202,COB,COGENT,-85.7579,38.2562,2020/5/24 11:50
1888,United States,"Louisville, KY",220 West Main Street,220 West Main Street,40202,COB,COGENT,-78.4828,38.0311,2020/5/24 11:50
1889,United States,"Louisville, KY",321 W Main St,Waterfront Plaza - East Tower,40202,COB,COGENT,-85.1762,40.571,2020/5/24 11:50
1890,United States,"Louisville, KY",323 West Main,Waterfront Plaza - Center Tower,40202,COB,COGENT,-97.6586,35.1396,2020/5/24 11:50
1891,United States,"Louisville, KY",325 West Main,Waterfront Plaza - West Tower,40202,COB,COGENT,-81.2427,33.9847,2020/5/24 11:50
1892,United States,"Louisville, KY",332 W. Broadway,Cogent Data Center - Louisville,40202,CDC,COGENT,-85.758,38.2459,2020/5/24 11:50
1893,United States,"Louisville, KY",332 W. Broadway,Heyburn Building,40202,COB,COGENT,-85.758,38.2459,2020/5/24 11:50
1894,United States,"Louisville, KY",400 W Market St,National City Tower,40202,COB,COGENT,-81.5312,41.0904,2020/5/24 11:50
1895,United States,"Louisville, KY",401 South 4th Street,Brown & Williamson Tower,40202,COB,COGENT,-97.3375,31.1002,2020/5/24 11:50
1896,United States,"Louisville, KY",401 West Main Street,One Riverfront Plaza,40202,COB,COGENT,-78.4865,38.031,2020/5/24 11:50
1897,United States,"Louisville, KY",462 S 4th St,Meidinger Tower,40202,COB,COGENT,-85.7576,38.2514,2020/5/24 11:50
1898,United States,"Louisville, KY",500 W. Jefferson Street,PNC Plaza,40202,COB,COGENT,-85.7592,38.2538,2020/5/24 11:50
1900,United States,"Manassas, VA",11680 Hayden Rd,Iron Mountain,20109,CNDC,COGENT,-77.5453,38.774,2020/5/24 11:50
1901,United States,"Manassas, VA",7400 Infantry Ridge Rd,Equinix DC 14,20109,CNDC,COGENT,-77.5105,38.8035,2020/5/24 11:50
1902,United States,"Manassas, VA",9651 Hornbaker Road,COPT-Powerloft,20109,CNDC,COGENT,-77.5317,38.7475,2020/5/24 11:50
1903,United States,"Marietta, GA",1130 Powers Ferry Place,DataSite,30067,CNDC,COGENT,-84.4816,33.9255,2020/5/24 11:50
1904,United States,"Marina Del Rey, CA",4640 Admiralty Way,Building MMR,90292,COB,COGENT,-118.44,33.9812,2020/5/24 11:50
1905,United States,"Marina Del Rey, CA",4676 Admiralty Way,Marina Tower South,90292,,COGENT,-118.44,33.9801,2020/5/24 11:50
1906,United States,"Marina Del Rey, CA",4676 Admiralty Way,USC/ISI,90292,CNDC,COGENT,-118.44,33.9801,2020/5/24 11:50
1907,United States,"Marina Del Rey, CA",4676 Admiralty Way,USC/ISI Colo 2 MMR,90292,CNDC,COGENT,-118.44,33.9801,2020/5/24 11:50
1908,United States,"Marlborough, MA",34 Saint Martin Dr,Tierpoint MRL,1752,CNDC,COGENT,-71.5793,42.3129,2020/5/24 11:50
1909,United States,"McAllen, TX",200 S 10th Street,CarrierCom,78501,CNDC,COGENT,-77.4379,37.5359,2020/5/24 11:50
1910,United States,"McAllen, TX",200 S 10th Street,McAllen Data Center,78501,CNDC,COGENT,-77.4379,37.5359,2020/5/24 11:50
1911,United States,"McAllen, TX",900 1/2 Beech Avenue,VTX,78501,CNDC,COGENT,-89.4977,31.5636,2020/5/24 11:50
1912,United States,"McClellan, CA",4235 Forcum Ave,Quest Data Center,95652,CNDC,COGENT,-121.407,38.6448,2020/5/24 11:50
1913,United States,"McLean, VA",1600 Tysons Blvd,1600 Tysons Blvd,22102,COB,COGENT,-77.2255,38.9264,2020/5/24 11:50
1914,United States,"McLean, VA",1650 Tysons Blvd,The Corp. Office Centre @ Tysons II,22102,COB,COGENT,-77.2234,38.925,2020/5/24 11:50
1915,United States,"McLean, VA",1750 Tysons Blvd,The Corp. Office Centre @ Tysons II,22102,COB,COGENT,-77.2228,38.924,2020/5/24 11:50
1916,United States,"McLean, VA",1755 Old Meadow Rd - 1st Floor,CenturyLink (fka. Level 3),22102,CNDC,COGENT,-77.2116,38.92,2020/5/24 11:50
1917,United States,"McLean, VA",1765 Greensboro Station Place,Boro Station l,22102,COB,COGENT,-77.2334,38.922,2020/5/24 11:50
1918,United States,"Mclean, VA",1775 Greensboro Station Place,Boro Station lll,22102,COB,COGENT,-77.2333,38.9214,2020/5/24 11:50
1919,United States,"McLean, VA",1785 Greensboro Station Place,Boro Station ll,22102,COB,COGENT,-77.2333,38.9219,2020/5/24 11:50
1920,United States,"McLean, VA",8180 Greensboro Dr,Greensboro Park,22102,COB,COGENT,-77.2277,38.9236,2020/5/24 11:50
1921,United States,"McLean, VA",8200 Greensboro Dr,Greensboro Park,22102,COB,COGENT,-77.2284,38.9232,2020/5/24 11:50
1922,United States,"McLean, VA",8300 Greensboro Dr,8300 Greensboro Dr,22102,COB,COGENT,-77.232,38.9252,2020/5/24 11:50
1924,United States,"Memphis, TN",1715 Aaron Brenner Drive,Renaissance Center,38138,COB,COGENT,-89.8462,35.1025,2020/5/24 11:50
1925,United States,"Memphis, TN",200 Jefferson Avenue,One Memphis Place,38103,COB,COGENT,-76.4109,36.9654,2020/5/24 11:50
1926,United States,"Memphis, TN",22 North Front Street,Falls Building,38103,COB,COGENT,-91.7853,34.542,2020/5/24 11:50
1927,United States,"Memphis, TN",3706 Alumni Avenue,"Univ of Memphis, 101 Jones Hall",38152,CNDC,COGENT,-89.9397,35.1188,2020/5/24 11:50
1928,United States,"Memphis, TN",40 S Main Street,One Commerce Square,38103,COB,COGENT,-90.0535,35.1441,2020/5/24 11:50
1929,United States,"Memphis, TN",50 N Front Street,Raymond James Tower,38103,COB,COGENT,-75.142,39.9513,2020/5/24 11:50
1930,United States,"Memphis, TN",5100 Poplar Avenue,Clark Tower,38137,COB,COGENT,-89.8917,35.1129,2020/5/24 11:50
1931,United States,"Memphis, TN",6750 Poplar Avenue,Forum I,38138,COB,COGENT,-89.8381,35.0989,2020/5/24 11:50
1932,United States,"Memphis, TN",7620 Appling Center Drive 1st floor,zColo,38133,CNDC,COGENT,-89.8086,35.2031,2020/5/24 11:50
1933,United States,"Mesa, AZ",1301 W University Rd,Evoque MESAAZLB,85201,CNDC,COGENT,-111.859,33.4216,2020/5/24 11:50
1934,United States,"Miami, FL",1 SE 3rd Ave,SunTrust International Center,33131,COB,COGENT,-80.1883,25.7739,2020/5/24 11:50
1935,United States,"Miami, FL",100 N Biscayne Blvd,100 N Biscayne Blvd,33132,COB,COGENT,-80.1882,25.7756,2020/5/24 11:50
1936,United States,"Miami, FL",100 N Biscayne Blvd 2nd Floor,Digiport,33132,CNDC,COGENT,-80.1882,25.7756,2020/5/24 11:50
1937,United States,"Miami, FL",100 NE 80th Terrace,Miami Data Vault,33138,CNDC,COGENT,-80.1949,25.8484,2020/5/24 11:50
1938,United States,"Miami, FL",100 S Biscayne Blvd,One Bayfront Plaza,33131,COB,COGENT,-80.1879,25.7732,2020/5/24 11:50
1939,United States,"Miami, FL",100 SE 2nd St,Bank of America Tower,33131,COB,COGENT,-96.7198,33.9959,2020/5/24 11:50
1940,United States,"Miami, FL",11234 NW 20th Street,QTS Datacenter,33172,CNDC,COGENT,-80.3796,25.7915,2020/5/24 11:50
1941,United States,"Miami, FL",11300 NW 25th Street,Telefonica KeyCenter,33172,CNDC,COGENT,-80.38,25.7961,2020/5/24 11:50
1942,United States,"Miami, FL",1221 Brickell Ave,Twelve Twenty-One Brickell Bldg,33131,COB,COGENT,-80.1911,25.7612,2020/5/24 11:50
1943,United States,"Miami, FL",1395 Brickell Avenue,Espirito Santo Plaza,33131,COB,COGENT,-80.1916,25.7607,2020/5/24 11:50
1944,United States,"Miami, FL",1450 Brickell Avenue,1450 Brickell Avenue,33131,COB,COGENT,-80.1931,25.7589,2020/5/24 11:50
1945,United States,"Miami, FL",1961 NW 22nd Street (fka 1953 NW 22nd St),Infolink,33142,CNDC,COGENT,-77.0488,38.9032,2020/5/24 11:50
1946,United States,"Miami, FL",2 S Biscayne Blvd,One Biscayne Tower,33131,COB,COGENT,-80.188,25.7738,2020/5/24 11:50
1947,United States,"Miami, FL",200 S Biscayne Blvd,Southeast Financial Center,33131,COB,COGENT,-80.1882,25.7764,2020/5/24 11:50
1948,United States,"Miami, FL",200 SE 1st St,Cogent Data Center - Miami,33131,CDC,COGENT,-80.1899,25.7732,2020/5/24 11:50
1949,United States,"Miami, FL",200 SE 1st St,200 Southeast 1st Building,33131,CNDC,COGENT,-80.1899,25.7732,2020/5/24 11:50
1950,United States,"Miami, FL",201 S Biscayne Blvd,Miami Center,33131,COB,COGENT,-80.1892,25.7713,2020/5/24 11:50
1951,United States,"Miami, FL",2115 NW 22nd St,CoreSite MI1 (The Miami Exchange),33142,CNDC,COGENT,-80.2301,25.7975,2020/5/24 11:50
1952,United States,"Miami, FL",36 NE 2nd ST,ColoHouse,33132,CNDC,COGENT,-80.1931,25.7759,2020/5/24 11:50
1953,United States,"Miami, FL",36 NE 2ND ST,Equinix MI2 (fka. S&D/L1),33132,CNDC,COGENT,-80.1931,25.7759,2020/5/24 11:50
1954,United States,"Miami, FL",36 NE 2nd ST,Digital Realty MIA1 (fka. Telx MMR),33132,CNDC,COGENT,-80.1931,25.7759,2020/5/24 11:50
1955,United States,"Miami, FL",36 NE 2nd ST,Interconnect Miami,33132,CNDC,COGENT,-80.1931,25.7759,2020/5/24 11:50
1956,United States,"Miami, FL",36 NE 2nd Street,zColo,33132,CNDC,COGENT,-80.1931,25.7759,2020/5/24 11:50
1957,United States,"Miami, FL",444 Brickell Ave,Rivergate Plaza,33131,COB,COGENT,-80.1905,25.7692,2020/5/24 11:50
1958,United States,"Miami, FL",49 NW 5th Street,CenturyLink (fka. Level 3),33128,CNDC,COGENT,-80.1949,25.7792,2020/5/24 11:50
1959,United States,"Miami, FL",50 NE 9th St.,Equinix MI1 (fka. Terremark / Verizon NAP of Ameri,33132,CNDC,COGENT,-80.1931,25.7824,2020/5/24 11:50
1960,United States,"Miami, FL",701 Brickell Ave,NationsBank Tower,33131,COB,COGENT,-80.1887,25.7668,2020/5/24 11:50
1961,United States,"Miami, FL",777 Brickell Ave,Brickell Office Plaza,33131,COB,COGENT,-80.1899,25.7664,2020/5/24 11:50
1962,United States,"Miami, FL",80 SW 8th St,Brickell City Tower,33130,COB,COGENT,-80.1946,25.766,2020/5/24 11:50
1963,United States,"Miami, FL",801 Brickell Ave,One Brickell Square,33131,COB,COGENT,-80.19,25.7658,2020/5/24 11:50
1964,United States,"Miamisburg, OH",3366 South Tech Blvd.,Iron Mountain (fka. IO Data Center),45342,CNDC,COGENT,-84.2364,39.5911,2020/5/24 11:50
1965,United States,"Milpitas, CA",1656 McCarthy Blvd,CoreSite SV2,95035,CNDC,COGENT,-121.916,37.4057,2020/5/24 11:50
1966,United States,"Milwaukee, WI",100 E Wisconsin Ave,The 100 East Building,53202,COB,COGENT,-87.9095,43.039,2020/5/24 11:50
1967,United States,"Milwaukee, WI",111 E Kilbourn Ave,Milwaukee Center,53202,COB,COGENT,-87.9114,43.0417,2020/5/24 11:50
1968,United States,"Milwaukee, WI",111 E Wisconsin Ave,Chase Tower,53202,COB,COGENT,-87.9094,43.0379,2020/5/24 11:50
1969,United States,"Milwaukee, WI",324 E Wisconsin,Cogent Data Center - Milwaukee,53202,CDC,COGENT,-87.907,43.039,2020/5/24 11:50
1970,United States,"Milwaukee, WI",324 E Wisconsin,Wells Building,53202,COB,COGENT,-87.907,43.039,2020/5/24 11:50
1971,United States,"Milwaukee, WI",324 E Wisconsin Suite 825,TSR,53202,CNDC,COGENT,-87.907,43.039,2020/5/24 11:50
1972,United States,"Milwaukee, WI",330 E Kilbourn Ave -Tower One,Plaza East - SE Tower One,53202,COB,COGENT,-87.9077,43.0432,2020/5/24 11:50
1973,United States,"Milwaukee, WI",330 E Kilbourn Ave -Tower Two,Plaza East - NW Tower Two,53202,COB,COGENT,-87.9077,43.0432,2020/5/24 11:50
1974,United States,"Milwaukee, WI",411 E Wisconsin Ave,411 East Wisconsin Center,53202,COB,COGENT,-87.9059,43.0384,2020/5/24 11:50
1975,United States,"Milwaukee, WI",777 E Wisconsin Ave,US Bank Center,53202,COB,COGENT,-87.9019,43.0383,2020/5/24 11:50
1976,United States,"Minneapolis, MN",100 South Fifth Street,5th Street Towers - 100 Tower,55402,COB,COGENT,-93.2687,44.9784,2020/5/24 11:50
1977,United States,"Minneapolis, MN",100 Washington Ave South,100 Washington Ave South,55401,COB,COGENT,-93.2658,44.9812,2020/5/24 11:50
1978,United States,"Minneapolis, MN",111 Washington Ave South,111 Washington Ave South,55401,COB,COGENT,-93.2668,44.9805,2020/5/24 11:50
1979,United States,"Minneapolis, MN",120 South Sixth Street,One Financial Plaza,55402,COB,COGENT,-93.2692,44.9772,2020/5/24 11:50
1980,United States,"Minneapolis, MN",121 South Eighth Street,TCF Tower,55402,COB,COGENT,-93.2712,44.9748,2020/5/24 11:50
1981,United States,"Minneapolis, MN",150 South Fifth Street,5th Street Towers - 150 Tower,55402,COB,COGENT,-93.268,44.9781,2020/5/24 11:50
1982,United States,"Minneapolis, MN",200 South 6th St,US Bank Plaza -North Tower,55402,COB,COGENT,-93.2677,44.9771,2020/5/24 11:50
1983,United States,"Minneapolis, MN",220 South 6th Street,US Bank Plaza - South Tower,55402,COB,COGENT,-93.2676,44.9769,2020/5/24 11:50
1984,United States,"Minneapolis, MN",222 South Ninth Street,Campbell Mithun Tower,55402,COB,COGENT,-93.2699,44.9737,2020/5/24 11:50
1985,United States,"Minneapolis, MN",225 South Sixth Street & 650 3rd Street,Capella Tower,55402,COB,COGENT,-93.2686,44.9762,2020/5/24 11:50
1986,United States,"Minneapolis, MN",250 Marquette Avenue,Cogent Data Center - Minneapolis,55401,CDC,COGENT,-93.2676,44.9809,2020/5/24 11:50
1987,United States,"Minneapolis, MN",250 Marquette Avenue,Marquette Plaza,55401,COB,COGENT,-93.2676,44.9809,2020/5/24 11:50
1988,United States,"Minneapolis, MN",33 South Sixth Street,Minneapolis City Center,55402,COB,COGENT,-93.2728,44.9775,2020/5/24 11:50
1989,United States,"Minneapolis, MN",333 South 7th St,333 South 7th St,55402,COB,COGENT,-93.2677,44.9743,2020/5/24 11:50
1990,United States,"Minneapolis, MN",50 South Sixth Street,50 South Sixth Street,55402,COB,COGENT,-93.2719,44.9783,2020/5/24 11:50
1991,United States,"Minneapolis, MN",50 South Tenth Street,Retek on the Mall,55403,COB,COGENT,-93.2753,44.9746,2020/5/24 11:50
1992,United States,"Minneapolis, MN",511 11th Ave S,Cologix MIN1,55415,CNDC,COGENT,-93.2546,44.9714,2020/5/24 11:50
1993,United States,"Minneapolis, MN",511 11th Ave S,vXchnge,55415,CNDC,COGENT,-93.2546,44.9714,2020/5/24 11:50
1994,United States,"Minneapolis, MN",60 S 6th St,RBC Dain Rauscher Plaza,55402,COB,COGENT,-93.2707,44.9779,2020/5/24 11:50
1995,United States,"Minneapolis, MN",80 South Eighth Street,IDS Center,55402,COB,COGENT,-93.2725,44.976,2020/5/24 11:50
1996,United States,"Minneapolis, MN",800 LaSalle Avenue,LaSalle Plaza,55402,COB,COGENT,-93.2751,44.9764,2020/5/24 11:50
1997,United States,"Minneapolis, MN",800 Nicollet Mall,US Bancorp Center,55402,COB,COGENT,-93.2739,44.9756,2020/5/24 11:50
1998,United States,"Minneapolis, MN",801 South Marquette Avenue,801 Marquette,55402,COB,COGENT,-93.2713,44.9749,2020/5/24 11:50
1999,United States,"Minneapolis, MN",90 South Seventh St.,Wells Fargo Bank Building,55402,COB,COGENT,-93.2709,44.977,2020/5/24 11:50
2000,United States,"Minneapolis, MN",900 Second Ave South,International Centre,55402,COB,COGENT,-93.2719,44.9737,2020/5/24 11:50
2001,United States,"Minneapolis, MN",901 Marquette Avenue,AT&T Tower,55402,COB,COGENT,-93.2725,44.974,2020/5/24 11:50
2002,United States,"Minneapolis, MN",920 Second Avenue South,Oracle Centre,55402,COB,COGENT,-93.2722,44.9733,2020/5/24 11:50
2003,United States,"Murray, UT",3949 S 200 East,Flexential - Millcreek (formerly Peak 10 - East),84107,CNDC,COGENT,-111.885,40.6855,2020/5/24 11:50
2004,United States,"Nashville, TN",147 Fourth Avenue N,365 Data Centers (fka. Equinix - S&D),37219,CNDC,COGENT,-86.7786,36.1627,2020/5/24 11:50
2005,United States,"Nashville, TN",150 3rd Avenue South,The Pinnacle at Symphony Place,37201,COB,COGENT,-86.7744,36.16,2020/5/24 11:50
2006,United States,"Nashville, TN",150 Fourth Avenue North,One Nashville Place,37219,COB,COGENT,-86.778,36.1628,2020/5/24 11:50
2007,United States,"Nashville, TN",201 Fourth Avenue North,Fourth & Church Building,37219,COB,COGENT,-86.7795,36.1638,2020/5/24 11:50
2008,United States,"Nashville, TN",211 Commerce Street,vXchnge,37201,CNDC,COGENT,-86.7767,36.1632,2020/5/24 11:50
2009,United States,"Nashville, TN",222 2ND Avenue South,222 Nashville,37201,COB,COGENT,-86.7732,36.1596,2020/5/24 11:50
2010,United States,"Nashville, TN",401 Commerce Street,SunTrust Plaza,37219,COB,COGENT,-86.7787,36.1617,2020/5/24 11:50
2011,United States,"Nashville, TN",414 Union Street,Bank of America Plaza,37219,COB,COGENT,-86.7809,36.1649,2020/5/24 11:50
2012,United States,"Nashville, TN",511 Union Street,Nashville City Center,37219,COB,COGENT,-86.7818,36.164,2020/5/24 11:50
2013,United States,"Needham, MA",105 Cabot St,Digital Realty,2494,CNDC,COGENT,-71.2196,42.3023,2020/5/24 11:50
2014,United States,"Needham, MA",128 First Ave,Digital Realty,2494,CNDC,COGENT,-73.9852,40.7273,2020/5/24 11:50
2015,United States,"New Orleans, LA",1100 Poydras St,Energy Centre,70163,COB,COGENT,-90.0755,29.9499,2020/5/24 11:50
2016,United States,"New Orleans, LA",1515 Poydras St,1515 Poydras,70112,COB,COGENT,-90.0783,29.9528,2020/5/24 11:50
2017,United States,"New Orleans, LA",1555 Poydras St,1555 Poydras,70112,COB,COGENT,-90.0792,29.9528,2020/5/24 11:50
2018,United States,"New Orleans, LA",1615 Poydras St,Freeport McMoRan Building,70112,COB,COGENT,-90.08,29.9528,2020/5/24 11:50
2019,United States,"New Orleans, LA",400 Poydras Street,400 Poydras Tower,70130,COB,COGENT,-90.0679,29.9485,2020/5/24 11:50
2020,United States,"New Orleans, LA",601 Poydras Street,Pan American Life Center,70130,COB,COGENT,-90.0697,29.95,2020/5/24 11:50
2021,United States,"New Orleans, LA",650 Poydras Street,Cogent Data Center - New Orleans,70130,CDC,COGENT,-90.0701,29.9489,2020/5/24 11:50
2022,United States,"New Orleans, LA",650 Poydras Street,Poydras Center,70130,COB,COGENT,-90.0701,29.9489,2020/5/24 11:50
2023,United States,"New Orleans, LA",701 Poydras,1 Shell Square,70139,COB,COGENT,-90.0714,29.9503,2020/5/24 11:50
2024,United States,"New Orleans, LA",909 Poydras,First Bank and Trust Tower,70112,COB,COGENT,-90.0737,29.9506,2020/5/24 11:50
2025,United States,"New York, NY",1 New York Plz,1 New York Plz,10004,COB,COGENT,-74.0119,40.7021,2020/5/24 11:50
2026,United States,"New York, NY",1 Park Ave,1 Park Ave,10016,COB,COGENT,-73.9815,40.7459,2020/5/24 11:50
2027,United States,"New York, NY",1 Penn Plz,One Penn Plaza,10119,COB,COGENT,-73.9922,40.7512,2020/5/24 11:50
2028,United States,"New York, NY",1 State Street Plz,One State Street Plaza,10004,COB,COGENT,-71.0576,42.3585,2020/5/24 11:50
2029,United States,"New York, NY",1 Whitehall St,Topps Company Bldg,10004,COB,COGENT,-73.6975,42.7629,2020/5/24 11:50
2030,United States,"New York, NY",1 World Trade Center / 285 Fulton St,1 World Trade Center (formerly Freedom Tower),10007,COB,COGENT,-74.0134,40.7127,2020/5/24 11:50
2031,United States,"New York, NY",10 E 53rd St,Harper Collins Bldg,10022,COB,COGENT,-89.9968,39.8832,2020/5/24 11:50
2032,United States,"New York, NY",10 Rockefeller Plaza,10 Rockefeller Plaza,10020,COB,COGENT,-73.9821,40.7606,2020/5/24 11:50
2033,United States,"New York, NY",100 Broadway,100 Broadway,10005,COB,COGENT,-74.0111,40.708,2020/5/24 11:50
2034,United States,"New York, NY",100 Park Ave,100 Park Ave,10017,COB,COGENT,-73.979,40.7514,2020/5/24 11:50
2035,United States,"New York, NY",100 Wall St,100 Wall St,10005,COB,COGENT,-74.0067,40.7051,2020/5/24 11:50
2036,United States,"New York, NY",100 William St,Sylvan Lawrence Bldg,10038,COB,COGENT,-74.0072,40.7084,2020/5/24 11:50
2037,United States,"New York, NY",101 Park Ave,101 Park Ave,10178,COB,COGENT,-73.9776,40.751,2020/5/24 11:50
2039,United States,"New York, NY",11 Broadway,Bowling Green Building,10004,COB,COGENT,-74.014,40.7053,2020/5/24 11:50
2040,United States,"New York, NY",11 Penn Plz,Matthew Bender Bldg,10001,COB,COGENT,-73.9914,40.7491,2020/5/24 11:50
2041,United States,"New York, NY",11 W 42nd St,11 W 42nd St,10036,COB,COGENT,-73.9818,40.7544,2020/5/24 11:50
2042,United States,"New York, NY",110 E 42nd St,Greenpoint Bldg,10017,COB,COGENT,-73.9771,40.7516,2020/5/24 11:50
2043,United States,"New York, NY",110 E 59th St,110 E 59th St,10022,COB,COGENT,-73.9694,40.7623,2020/5/24 11:50
2044,United States,"New York, NY",110 William St,110 William St,10038,COB,COGENT,-74.0068,40.7088,2020/5/24 11:50
2045,United States,"New York, NY",111 Broadway,Trinity Bldg,10006,COB,COGENT,-73.8626,41.0463,2020/5/24 11:50
2046,United States,"New York, NY",111 Eighth Ave,BT - Suite 3113,10011,CNDC,COGENT,-74.0034,40.7415,2020/5/24 11:50
2047,United States,"New York, NY",111 Eighth Ave,Digital Realty NYC2 (fka. Telx MMR),10011,CNDC,COGENT,-74.0034,40.7415,2020/5/24 11:50
2049,United States,"New York, NY",111 Eighth Ave - 15th Floor,Equinix NY9 (15th Floor) (fka. S&D),10011,CNDC,COGENT,-74.0034,40.7415,2020/5/24 11:50
2050,United States,"New York, NY",111 Eighth Ave - 3rd Floor,CenturyLink (fka. Level 3),10011,CNDC,COGENT,-74.0034,40.7415,2020/5/24 11:50
2051,United States,"New York, NY",111 Eighth Ave - Suite 734,Equinix NY9 (7th Floor) (fka. S&D/PAIX),10011,CNDC,COGENT,-74.0045,40.7418,2020/5/24 11:50
2052,United States,"New York, NY",111 W 33rd St,111 W 33rd St / 112 W 34th St,10001,COB,COGENT,-73.9889,40.7497,2020/5/24 11:50
2053,United States,"New York, NY",1114 Avenue of the Americas,Grace Bldg,10036,COB,COGENT,-73.9829,40.7547,2020/5/24 11:50
2054,United States,"New York, NY",1120 Avenue of the Americas,Hippodrome Bldg,10036,COB,COGENT,-73.9832,40.7558,2020/5/24 11:50
2055,United States,"New York, NY",1133 Avenue of the Americas,1133 Avenue of the Americas,10036,COB,COGENT,-73.9837,40.7559,2020/5/24 11:50
2056,United States,"New York, NY",1140 Avenue of the Americas,1140 Avenue of the Americas,10036,COB,COGENT,-73.9828,40.7562,2020/5/24 11:50
2057,United States,"New York, NY",115 Broadway,U.S. Realty Bldg,10006,COB,COGENT,-73.8802,41.008,2020/5/24 11:50
2058,United States,"New York, NY",1155 Avenue of the Americas,1155 Avenue of the Americas,10036,COB,COGENT,-73.9834,40.7567,2020/5/24 11:50
2059,United States,"New York, NY",1177 Avenue of the Americas,Americas Tower,10036,COB,COGENT,-73.9829,40.7573,2020/5/24 11:50
2060,United States,"New York, NY",1180 Avenue of the Americas,1180 Avenue of the Americas,10036,COB,COGENT,-73.9816,40.7574,2020/5/24 11:50
2061,United States,"New York, NY",1185 Avenue of the Americas,Stevens Tower,10036,COB,COGENT,-73.9823,40.758,2020/5/24 11:50
2062,United States,"New York, NY",119 W 40th St (aka 114 West 41st St),119 W 40th St,10018,COB,COGENT,-73.9856,40.7544,2020/5/24 11:50
2063,United States,"New York, NY",12 E 49th Street,Tower 49,10017,COB,COGENT,-73.977,40.7572,2020/5/24 11:50
2064,United States,"New York, NY",120 Broadway,Equitable Bldg,10271,COB,COGENT,-74.0106,40.7085,2020/5/24 11:50
2065,United States,"New York, NY",120 W 45th St,Tower 45,10036,COB,COGENT,-73.9839,40.757,2020/5/24 11:50
2066,United States,"New York, NY",1211 Avenue of the Americas,1211 Avenue of the Americas,10036,COB,COGENT,-73.9824,40.7587,2020/5/24 11:50
2067,United States,"New York, NY",122 E 42nd St (aka 380 Lexington Avenue),Chanin Bldg,10168,COB,COGENT,-73.9764,40.7514,2020/5/24 11:50
2068,United States,"New York, NY",1230 Avenue of the Americas,Simon & Schuster Bldg,10020,COB,COGENT,-73.9808,40.7588,2020/5/24 11:50
2069,United States,"New York, NY",125 Broad St,2 New York Plaza,10004,COB,COGENT,-74.0103,40.7024,2020/5/24 11:50
2070,United States,"New York, NY",125 Park Ave,125 Park Ave,10017,COB,COGENT,-73.9774,40.7517,2020/5/24 11:50
2071,United States,"New York, NY",1250 Broadway,1250 Broadway,10001,COB,COGENT,-73.9879,40.7477,2020/5/24 11:50
2072,United States,"New York, NY",1251 Avenue of the Americas,Mitsui Bldg,10020,COB,COGENT,-73.9818,40.7601,2020/5/24 11:50
2073,United States,"New York, NY",126 E 56th St,Tower 56,10022,COB,COGENT,-73.9705,40.7604,2020/5/24 11:50
2074,United States,"New York, NY",1270 Avenue of the Americas,1270 Avenue of the Americas,10020,COB,COGENT,-73.9798,40.7602,2020/5/24 11:50
2075,United States,"New York, NY",1290 Avenue of the Americas,1290 Avenue of the Americas,10104,COB,COGENT,-73.9796,40.7608,2020/5/24 11:50
2076,United States,"New York, NY",130 W 42nd St,Bush Tower,10036,COB,COGENT,-73.9755,40.7508,2020/5/24 11:50
2077,United States,"New York, NY",1301 Avenue of the Americas,Calyon Building,10019,COB,COGENT,-73.98,40.7618,2020/5/24 11:50
2078,United States,"New York, NY",1325 Avenue of the Americas,1325 Avenue of the Americas,10019,COB,COGENT,-73.9809,40.7627,2020/5/24 11:50
2079,United States,"New York, NY",1330 Avenue of the Americas,Pearson Building,10019,COB,COGENT,-73.9784,40.7619,2020/5/24 11:50
2080,United States,"New York, NY",1333 Broadway,1333 Broadway,10018,COB,COGENT,-73.988,40.7511,2020/5/24 11:50
2081,United States,"New York, NY",1345 Avenue of the Americas,The Alliance Bernstein Bldg,10105,COB,COGENT,-73.9791,40.7631,2020/5/24 11:50
2082,United States,"New York, NY",135 E 57th St,135 E 57th St,10022,COB,COGENT,-73.9695,40.7613,2020/5/24 11:50
2083,United States,"New York, NY",135 W 50th St,Sports Illustrated Bldg,10020,COB,COGENT,-73.9826,40.7609,2020/5/24 11:50
2084,United States,"New York, NY",1350 Avenue of the Americas,MGM Bldg,10019,COB,COGENT,-73.9778,40.7626,2020/5/24 11:50
2085,United States,"New York, NY",1350 Broadway,Herald Square Bldg,10018,COB,COGENT,-73.9873,40.7511,2020/5/24 11:50
2086,United States,"New York, NY",1359 Broadway,1359 Broadway,10018,COB,COGENT,-73.9878,40.7518,2020/5/24 11:50
2087,United States,"New York, NY",1370 Avenue of the Americas,1370 Avenue of the Americas,10019,COB,COGENT,-73.9775,40.7633,2020/5/24 11:50
2088,United States,"New York, NY",1370 Broadway,1370 Broadway,10018,COB,COGENT,-73.9872,40.7518,2020/5/24 11:50
2089,United States,"New York, NY",1372 Broadway,The National Tower Bldg,10018,COB,COGENT,-73.9872,40.7524,2020/5/24 11:50
2090,United States,"New York, NY",14 Wall Street,14 Wall Street,10005,COB,COGENT,-74.0106,40.7076,2020/5/24 11:50
2091,United States,"New York, NY",140 Broadway,140 Broadway,10005,COB,COGENT,-74.0101,40.7087,2020/5/24 11:50
2092,United States,"New York, NY",140 E 45th St,2 Grand Central Tower,10017,COB,COGENT,-73.9741,40.7527,2020/5/24 11:50
2093,United States,"New York, NY",1400 Broadway,1400 Broadway,10018,COB,COGENT,-73.987,40.7531,2020/5/24 11:50
2094,United States,"New York, NY",1407 Broadway,1407 Broadway,10018,COB,COGENT,-73.9874,40.7534,2020/5/24 11:50
2095,United States,"New York, NY",1411 Broadway,World Apparel Center,10018,COB,COGENT,-73.9873,40.7541,2020/5/24 11:50
2096,United States,"New York, NY",1430 Broadway,1430 Broadway,10018,COB,COGENT,-73.9865,40.7539,2020/5/24 11:50
2097,United States,"New York, NY",1440 Broadway,1440 Broadway,10018,COB,COGENT,-73.9865,40.7544,2020/5/24 11:50
2098,United States,"New York, NY",1450 Broadway,1450 Broadway,10018,COB,COGENT,-73.9864,40.7547,2020/5/24 11:50
2099,United States,"New York, NY",1466 Broadway,1466 Broadway,10036,COB,COGENT,-73.9863,40.7556,2020/5/24 11:50
2101,United States,"New York, NY",150 Broadway,150 Broadway,10038,COB,COGENT,-74.0101,40.7092,2020/5/24 11:50
2102,United States,"New York, NY",150 E 42nd St,Hiro Plaza,10017,COB,COGENT,-73.9755,40.7509,2020/5/24 11:50
2103,United States,"New York, NY",150 E 58th St,Architects & Designers Bldg,10155,COB,COGENT,-73.9682,40.7611,2020/5/24 11:50
2104,United States,"New York, NY",150 Greenwich St,4 World Trade Center,10007,COB,COGENT,-74.0123,40.7103,2020/5/24 11:50
2105,United States,"New York, NY",1500 Broadway,Times Square Plaza,10036,COB,COGENT,-73.9857,40.7568,2020/5/24 11:50
2106,United States,"New York, NY",152 W 57th St,Carnegie Hall Tower,10019,COB,COGENT,-73.9795,40.765,2020/5/24 11:50
2107,United States,"New York, NY",1540 Broadway,Bertelsmann Bldg,10036,COB,COGENT,-73.9847,40.758,2020/5/24 11:50
2108,United States,"New York, NY",1633 Broadway,Paramount Plaza,10019,COB,COGENT,-73.8966,40.9089,2020/5/24 11:50
2109,United States,"New York, NY",165 Broadway,One Liberty Plaza,10006,COB,COGENT,-73.8823,40.9858,2020/5/24 11:50
2110,United States,"New York, NY",1675 Broadway,1675 Broadway,10019,COB,COGENT,-73.9835,40.7631,2020/5/24 11:50
2111,United States,"New York, NY",17 State St,17 State St,10004,COB,COGENT,-74.0139,40.7029,2020/5/24 11:50
2112,United States,"New York, NY",1700 Broadway,1700 Broadway,10019,COB,COGENT,-73.9823,40.7635,2020/5/24 11:50
2113,United States,"New York, NY",1740 Broadway,MONY Bldg,10019,COB,COGENT,-73.9815,40.7651,2020/5/24 11:50
2114,United States,"New York, NY",175 Greenwich Street,3 World Trade Center,10007,COB,COGENT,-74.0117,40.711,2020/5/24 11:50
2115,United States,"New York, NY",1775 Broadway (3 Columbus Circle),Columbus Tower,10019,COB,COGENT,-73.9832,40.7686,2020/5/24 11:50
2116,United States,"New York, NY",180 Madison Ave,180 Madison Ave,10016,COB,COGENT,-73.9835,40.7476,2020/5/24 11:50
2117,United States,"New York, NY",180 Maiden Ln,Continental Center,10038,COB,COGENT,-74.0055,40.7053,2020/5/24 11:50
2118,United States,"New York, NY",180 Varick St,180 Varick St,10014,COB,COGENT,-74.0052,40.7275,2020/5/24 11:50
2119,United States,"New York, NY",19 W 44th St,Berkeley Bldg,10036,COB,COGENT,-73.9808,40.7555,2020/5/24 11:50
2120,United States,"New York, NY",195 Broadway,195 Broadway,10007,COB,COGENT,-74.0096,40.7108,2020/5/24 11:50
2121,United States,"New York, NY",199 Water St,1 Seaport Plaza,10038,COB,COGENT,-74.0045,40.7069,2020/5/24 11:50
2122,United States,"New York, NY",2 Park Avenue,2 Park Avenue,10016,COB,COGENT,-73.9823,40.7462,2020/5/24 11:50
2123,United States,"New York, NY",2 Penn Plz,2 Penn Plz,10121,COB,COGENT,-73.9921,40.75,2020/5/24 11:50
2124,United States,"New York, NY",2 Rector St (aka 101 Greenwich St),Wall Street Plaza West,10006,COB,COGENT,-74.0132,40.7082,2020/5/24 11:50
2125,United States,"New York, NY",200 Liberty,One Brookfield Place,10281,COB,COGENT,-74.0156,40.7105,2020/5/24 11:50
2126,United States,"New York, NY",200 Madison Ave,Putnam Berkley,10016,COB,COGENT,-73.9826,40.7491,2020/5/24 11:50
2127,United States,"New York, NY",200 Park Ave,Met Life Bldg,10166,COB,COGENT,-73.9767,40.7533,2020/5/24 11:50
2128,United States,"New York, NY",200 Varick Street,200 Varick St,10014,COB,COGENT,-74.0051,40.7283,2020/5/24 11:50
2129,United States,"New York, NY",200 Vesey Street,Three Brookfield Place,10281,COB,COGENT,-74.0148,40.714,2020/5/24 11:50
2130,United States,"New York, NY",205 E 42nd St,Pfizer Building,10017,COB,COGENT,-73.9736,40.7508,2020/5/24 11:50
2131,United States,"New York, NY",22 Cortlandt St,22 Cortlandt St,10007,COB,COGENT,-74.0106,40.7102,2020/5/24 11:50
2132,United States,"New York, NY",220 E 42nd St,The News Bldg,10017,COB,COGENT,-73.9731,40.7499,2020/5/24 11:50
2133,United States,"New York, NY",225 Liberty Street,Two Brookfield Place,10281,COB,COGENT,-74.0157,40.7126,2020/5/24 11:50
2134,United States,"New York, NY",228 E 45th St,228 E 45th St,10017,COB,COGENT,-73.9718,40.7519,2020/5/24 11:50
2135,United States,"New York, NY",230 Park Ave,Hemsley Building,10169,COB,COGENT,-73.9762,40.7544,2020/5/24 11:50
2136,United States,"New York, NY",237 Park Ave,Park Avenue Atrium,10017,COB,COGENT,-73.9748,40.7536,2020/5/24 11:50
2137,United States,"New York, NY",24 State St,1 Battery Park Plz,10004,COB,COGENT,-71.0577,42.3591,2020/5/24 11:50
2138,United States,"New York, NY",245 Park Ave,The Bear Stearns Bldg,10167,COB,COGENT,-73.9745,40.7545,2020/5/24 11:50
2139,United States,"New York, NY",25 Broadway,Cogent Data Center - New York,10004,CDC,COGENT,-73.8581,41.0772,2020/5/24 11:50
2140,United States,"New York, NY",25 Broadway,Cunard Lines Building,10004,COB,COGENT,-73.8581,41.0772,2020/5/24 11:50
2141,United States,"New York, NY",25 W 45th St,The Central Bldg,10036,COB,COGENT,-73.9807,40.7561,2020/5/24 11:50
2142,United States,"New York, NY",250 Greenwich Street,7 World Trade Center,10007,COB,COGENT,-74.0119,40.7134,2020/5/24 11:50
2143,United States,"New York, NY",250 Park Ave,250 Park Ave,10177,COB,COGENT,-73.9759,40.7551,2020/5/24 11:50
2144,United States,"New York, NY",250 Vesey Street,Four Brookfield Place,10080,COB,COGENT,-74.0161,40.7141,2020/5/24 11:50
2145,United States,"New York, NY",250 W 55th St,250 W 55th,10019,COB,COGENT,-73.9832,40.765,2020/5/24 11:50
2146,United States,"New York, NY",250 W 57th St,Fisk Bldg,10019,COB,COGENT,-73.9824,40.7662,2020/5/24 11:50
2148,United States,"New York, NY",26 Broadway,26 Broadway,10004,COB,COGENT,-74.0131,40.7055,2020/5/24 11:50
2149,United States,"New York, NY",266 West 37th Street,266 West 37th Street,10018,COB,COGENT,-73.9919,40.7538,2020/5/24 11:50
2150,United States,"New York, NY",270 Madison Ave,270 Madison Ave,10016,COB,COGENT,-73.9807,40.7512,2020/5/24 11:50
2151,United States,"New York, NY",275 Madison Ave,275 Madison Ave,10016,COB,COGENT,-73.98,40.7512,2020/5/24 11:50
2152,United States,"New York, NY",277 Park Ave,277 Park Ave,10172,COB,COGENT,-73.9741,40.7552,2020/5/24 11:50
2153,United States,"New York, NY",28 Liberty Street,28 Liberty Street (1 Chase Manhattan Plaza),10005,COB,COGENT,-74.0089,40.7078,2020/5/24 11:50
2154,United States,"New York, NY",280 Park Ave E,Bankers Trust Bldg,10017,COB,COGENT,-73.9756,40.7567,2020/5/24 11:50
2155,United States,"New York, NY",280 Park Ave W,Bankers Trust Bldg,10017,COB,COGENT,-73.9756,40.7567,2020/5/24 11:50
2156,United States,"New York, NY",29 Broadway,29 Broadway,10006,COB,COGENT,-74.0138,40.703,2020/5/24 11:50
2157,United States,"New York, NY",299 Park Ave,The Westvaco Bldg,10171,COB,COGENT,-73.9741,40.756,2020/5/24 11:50
2158,United States,"New York, NY",3 Park Avenue,3 Park Avenue,10016,COB,COGENT,-73.9811,40.7466,2020/5/24 11:50
2159,United States,"New York, NY",3 Times Square,Reuters Building,10036,COB,COGENT,-73.9839,40.7607,2020/5/24 11:50
2160,United States,"New York, NY",30 Broad St,30 Broad St,10004,COB,COGENT,-74.0116,40.7063,2020/5/24 11:50
2161,United States,"New York, NY",30 Rockefeller Plaza,GE Bldg,10112,COB,COGENT,-73.98,40.7594,2020/5/24 11:50
2162,United States,"New York, NY",300 E 42nd St,300 E 42nd St,10017,COB,COGENT,-73.9721,40.7495,2020/5/24 11:50
2163,United States,"New York, NY",300 Park Ave,Colgate Palmolive Building,10022,COB,COGENT,-73.9747,40.757,2020/5/24 11:50
2164,United States,"New York, NY",31 W 52nd St,Deutsche Bank Building,10019,COB,COGENT,-73.9781,40.761,2020/5/24 11:50
2165,United States,"New York, NY",32 Avenue of the Americas,CoreSite NY1,10013,CNDC,COGENT,-74.0047,40.7201,2020/5/24 11:50
2166,United States,"New York, NY",32 Avenue of the Americas,Digital Realty NYC3 (fka. Telx MMR),10013,CNDC,COGENT,-74.0047,40.7201,2020/5/24 11:50
2167,United States,"New York, NY",32 Avenue of the Americas,Digital Realty (fka. TelX),10013,CNDC,COGENT,-74.0047,40.7201,2020/5/24 11:50
2168,United States,"New York, NY",32 Old Slip,Financial Square,10005,COB,COGENT,-74.0077,40.7037,2020/5/24 11:50
2169,United States,"New York, NY",320 Park Ave,Mutual of America Bldg,10022,COB,COGENT,-73.9743,40.7578,2020/5/24 11:50
2170,United States,"New York, NY",325 Hudson St.,Netrality,10013,CNDC,COGENT,-74.0077,40.7269,2020/5/24 11:50
2171,United States,"New York, NY",325 Hudson St.,Netrality MMR,10013,COB,COGENT,-74.0077,40.7269,2020/5/24 11:50
2172,United States,"New York, NY",33 Whitehall St,Broad Financial Center,10004,COB,COGENT,-74.0125,40.7034,2020/5/24 11:50
2173,United States,"New York, NY",330 Madison Ave,330 Madison Ave,10017,COB,COGENT,-73.9791,40.7532,2020/5/24 11:50
2174,United States,"New York, NY",330 Seventh Ave,Gross Bldg,10001,COB,COGENT,-73.9934,40.7477,2020/5/24 11:50
2175,United States,"New York, NY",333 Seventh Ave,333 Seventh Ave,10001,COB,COGENT,-73.9929,40.7475,2020/5/24 11:50
2176,United States,"New York, NY",340 Madison Ave,340 Madison Ave,10173,COB,COGENT,-73.9789,40.7539,2020/5/24 11:50
2177,United States,"New York, NY",345 Park Ave,345 Park Ave,10154,COB,COGENT,-73.9722,40.7579,2020/5/24 11:50
2178,United States,"New York, NY",350 Fifth Ave,Empire State Bldg,10118,COB,COGENT,-73.9851,40.7482,2020/5/24 11:50
2179,United States,"New York, NY",350 Park Avenue,350 Park Avenue,10022,COB,COGENT,-73.9739,40.7583,2020/5/24 11:50
2180,United States,"New York, NY",355 Lexington Avenue,355 Lexington Avenue,10017,COB,COGENT,-73.9765,40.7502,2020/5/24 11:50
2181,United States,"New York, NY",360 Lexington Ave,360 Lexington Ave,10017,COB,COGENT,-73.977,40.7505,2020/5/24 11:50
2182,United States,"New York, NY",370 Lexington Ave,370 Lexington Ave,10017,COB,COGENT,-73.9769,40.7507,2020/5/24 11:50
2183,United States,"New York, NY",375 Hudson St,Saatchi & Saatchi Bldg,10014,COB,COGENT,-74.0081,40.7285,2020/5/24 11:50
2184,United States,"New York, NY",375 Park Ave,Seagram Bldg,10152,COB,COGENT,-73.9722,40.7585,2020/5/24 11:50
2185,United States,"New York, NY",375 Pearl Street,Sabey - Intergate.Manhattan,10038,CNDC,COGENT,-74.0014,40.711,2020/5/24 11:50
2186,United States,"New York, NY",375 Pearl Street,Intergate Manhattan (15-30) MTOB Floors,10038,COB,COGENT,-74.0014,40.711,2020/5/24 11:50
2187,United States,"New York, NY",381 Park Ave South,381 Park Ave S,10016,COB,COGENT,-73.9842,40.7424,2020/5/24 11:50
2188,United States,"New York, NY",395 Hudson St,395 Hudson St,10014,COB,COGENT,-74.0078,40.7291,2020/5/24 11:50
2189,United States,"New York, NY",399 Park Ave,399 Park Ave,10022,COB,COGENT,-73.972,40.7593,2020/5/24 11:50
2190,United States,"New York, NY",4 Times Sq,Conde Nast Bldg,10036,COB,COGENT,-73.9857,40.7561,2020/5/24 11:50
2191,United States,"New York, NY",40 E 52nd St,Security Pacific Plaza,10022,COB,COGENT,-73.974,40.7587,2020/5/24 11:50
2192,United States,"New York, NY",40 Fulton St,40 Fulton St,10038,COB,COGENT,-74.0047,40.7079,2020/5/24 11:50
2193,United States,"New York, NY",40 W 57th St,Squibb Bldg,10019,COB,COGENT,-73.9762,40.7637,2020/5/24 11:50
2194,United States,"New York, NY",40 Wall St,The Trump Bldg,10005,COB,COGENT,-74.0097,40.7069,2020/5/24 11:50
2196,United States,"New York, NY",405 Lexington Ave,Chrysler Bldg,10174,COB,COGENT,-73.9755,40.7516,2020/5/24 11:50
2197,United States,"New York, NY",405 Park Ave,405 Park Ave,10022,COB,COGENT,-73.9717,40.7597,2020/5/24 11:50
2198,United States,"New York, NY",41 Madison Ave,New York Merchandise Mart,10010,COB,COGENT,-73.753,42.6453,2020/5/24 11:50
2199,United States,"New York, NY",410 Park Ave,Chase Manhattan Bank Bldg,10022,COB,COGENT,-73.9722,40.7603,2020/5/24 11:50
2200,United States,"New York, NY",415 Madison Avenue,415 Madison Avenue,10017,COB,COGENT,-73.9761,40.7567,2020/5/24 11:50
2201,United States,"New York, NY",417 Fifth Ave,417 Fifth Ave,10016,COB,COGENT,-73.9827,40.7506,2020/5/24 11:50
2202,United States,"New York, NY",420 Fifth Ave,420 Fifth Ave,10018,COB,COGENT,-73.9833,40.7509,2020/5/24 11:50
2203,United States,"New York, NY",420 Lexington Ave,Graybar Bldg,10170,COB,COGENT,-73.9759,40.7528,2020/5/24 11:50
2204,United States,"New York, NY",425 Lexington Ave,Commerce Plaza,10017,COB,COGENT,-73.975,40.7523,2020/5/24 11:50
2205,United States,"New York, NY",432 Park Avenue South,432 Park Avenue South,10016,COB,COGENT,-73.9718,40.7616,2020/5/24 11:50
2206,United States,"New York, NY",437 Madison Ave,437 Madison Ave,10022,COB,COGENT,-73.9752,40.7574,2020/5/24 11:50
2207,United States,"New York, NY",44 Wall St,44 Wall St,10005,COB,COGENT,-74.0094,40.7068,2020/5/24 11:50
2208,United States,"New York, NY",440 Park Avenue South,440 Park Avenue South,10013,COB,COGENT,-73.9714,40.7615,2020/5/24 11:50
2209,United States,"New York, NY",450 Lexington Ave,450 Lexington Ave,10017,COB,COGENT,-73.9754,40.7533,2020/5/24 11:50
2210,United States,"New York, NY",450 Park Ave,450 Park Ave,10022,COB,COGENT,-73.9714,40.7615,2020/5/24 11:50
2211,United States,"New York, NY",450 W 33rd St,Westyard Bldg,10001,COB,COGENT,-87.6363,41.8291,2020/5/24 11:50
2212,United States,"New York, NY",452 Fifth Ave,HSBC Bank World Headquarters,10018,COB,COGENT,-73.9823,40.7521,2020/5/24 11:50
2213,United States,"New York, NY",470 Park Avenue South,South Wing,10016,COB,COGENT,-73.9827,40.7457,2020/5/24 11:50
2214,United States,"New York, NY",475 Park Ave S,475 Park Ave S,10016,COB,COGENT,-73.9822,40.7454,2020/5/24 11:50
2215,United States,"New York, NY",48 Wall St,48 Wall St,10005,COB,COGENT,-74.0091,40.7065,2020/5/24 11:50
2216,United States,"New York, NY",485 Lexington Ave,Grand Central Square,10017,COB,COGENT,-73.9733,40.7538,2020/5/24 11:50
2217,United States,"New York, NY",488 Madison Ave,Institutional Investors Bldg,10022,COB,COGENT,-73.9751,40.7589,2020/5/24 11:50
2218,United States,"New York, NY",498 Seventh Ave,498 Seventh Ave,10018,COB,COGENT,-73.9897,40.7528,2020/5/24 11:50
2219,United States,"New York, NY",499 Park Ave,499 Park Ave,10022,COB,COGENT,-73.9696,40.7625,2020/5/24 11:50
2220,United States,"New York, NY",500 Fifth Ave,Salmon Tower,10110,COB,COGENT,-73.9809,40.7537,2020/5/24 11:50
2221,United States,"New York, NY",501 Seventh Ave,501 Seventh Ave,10018,COB,COGENT,-73.9889,40.753,2020/5/24 11:50
2222,United States,"New York, NY",501 West 30th St,10 Hudson Yards,10001,COB,COGENT,-74.0011,40.7525,2020/5/24 11:50
2223,United States,"New York, NY",51 W 52nd St,CBS Bldg,10019,COB,COGENT,-73.9788,40.7612,2020/5/24 11:50
2224,United States,"New York, NY",510 Madison Avenue,510 Madison Avenue,10022,COB,COGENT,-73.9746,40.7598,2020/5/24 11:50
2225,United States,"New York, NY",515 Madison Ave,515 Madison Ave,10022,COB,COGENT,-73.9738,40.7598,2020/5/24 11:50
2226,United States,"New York, NY",520 Eighth Ave,520 Eighth Ave,10018,COB,COGENT,-73.992,40.7536,2020/5/24 11:50
2227,United States,"New York, NY",520 Madison Ave,520 Madison Ave,10022,COB,COGENT,-73.9748,40.7602,2020/5/24 11:50
2228,United States,"New York, NY",521 Fifth Ave,Lefcourt National Bldg,10175,COB,COGENT,-73.9798,40.7542,2020/5/24 11:50
2229,United States,"New York, NY",535 Fifth Ave,535 Fifth Ave,10017,COB,COGENT,-73.9794,40.7548,2020/5/24 11:50
2230,United States,"New York, NY",535 Madison Ave,Dillon Read Bldg,10022,COB,COGENT,-73.9732,40.7605,2020/5/24 11:50
2231,United States,"New York, NY",540 Madison Ave,540 Madison Ave,10022,COB,COGENT,-73.9738,40.7609,2020/5/24 11:50
2232,United States,"New York, NY",545 Fifth Ave,Lorraine Bldg,10017,COB,COGENT,-73.9793,40.7551,2020/5/24 11:50
2233,United States,"New York, NY",55 Broad St,NY Information Tech Ctr,10004,COB,COGENT,-74.0111,40.7053,2020/5/24 11:50
2234,United States,"New York, NY",55 Broadway,One Exchange Plaza,10006,COB,COGENT,-74.013,40.7068,2020/5/24 11:50
2235,United States,"New York, NY",55 E 52nd St,Park Avenue Plaza,10055,COB,COGENT,-73.9735,40.7592,2020/5/24 11:50
2236,United States,"New York, NY",55 Water St,55 Water St,10041,COB,COGENT,-74.0089,40.7033,2020/5/24 11:50
2237,United States,"New York, NY",55 Water St 21st Floor,DTCC,10041,CNDC,COGENT,-74.0098,40.7034,2020/5/24 11:50
2238,United States,"New York, NY",550 West 34th St,55 Hudson Yards,10001,COB,COGENT,-74.0011,40.7553,2020/5/24 11:50
2239,United States,"New York, NY",551 Fifth Ave,Fred F. French Bldg,10176,COB,COGENT,-73.9789,40.7554,2020/5/24 11:50
2240,United States,"New York, NY",551 Madison Ave,551 Madison Ave,10022,COB,COGENT,-73.9729,40.7613,2020/5/24 11:50
2241,United States,"New York, NY",555 Fifth Ave,555 Fifth Ave,10017,COB,COGENT,-73.9789,40.7557,2020/5/24 11:50
2242,United States,"New York, NY",560 Lexington Avenue,560 Lexington Avenue,10022,COB,COGENT,-73.9727,40.7568,2020/5/24 11:50
2243,United States,"New York, NY",570 Lexington Ave,Tower 570,10022,COB,COGENT,-73.9726,40.7572,2020/5/24 11:50
2244,United States,"New York, NY",575 Lexington Ave,575 Lexington Ave,10022,COB,COGENT,-73.9713,40.7573,2020/5/24 11:50
2245,United States,"New York, NY",575 Madison Ave,575 Madison Ave,10022,COB,COGENT,-73.9724,40.7618,2020/5/24 11:50
2246,United States,"New York, NY",59 Maiden Lane,59 Maiden Lane,10038,COB,COGENT,-74.0081,40.7087,2020/5/24 11:50
2247,United States,"New York, NY",590 Madison Ave,IBM Bldg,10022,COB,COGENT,-73.9731,40.7621,2020/5/24 11:50
2248,United States,"New York, NY",599 Lexington Ave,599 Lexington Ave,10022,COB,COGENT,-73.9708,40.758,2020/5/24 11:50
2249,United States,"New York, NY",60 Broad St,25 Beaver St,10004,COB,COGENT,-74.0118,40.7052,2020/5/24 11:50
2250,United States,"New York, NY",60 E 42nd St,One Grand Central Place,10165,COB,COGENT,-73.9787,40.7522,2020/5/24 11:50
2251,United States,"New York, NY",60 Hudson St,zColo,10013,CNDC,COGENT,-74.0083,40.7177,2020/5/24 11:50
2252,United States,"New York, NY",60 Hudson St,Epsilon (fka. Metcom),10013,CNDC,COGENT,-74.0083,40.7177,2020/5/24 11:50
2253,United States,"New York, NY","60 Hudson St - Suite 1904, 1510 & 1602",Equinix NY8 (fka. S&D/RACO),10013,CNDC,COGENT,-74.0083,40.7177,2020/5/24 11:50
2254,United States,"New York, NY",60 Hudson St - Suite 900,Digital Realty NYC1 (fka. TelX),10013,CNDC,COGENT,-74.0083,40.7177,2020/5/24 11:50
2255,United States,"New York, NY",60 Madison Ave,60 Madison Ave,10010,COB,COGENT,-73.9865,40.7432,2020/5/24 11:50
2257,United States,"New York, NY",600 Third Ave,600 Third Ave,10016,COB,COGENT,-73.9759,40.7493,2020/5/24 11:50
2258,United States,"New York, NY",601 Lexington Avenue (aka 153 E 53rd St),Citigroup Center,10022,COB,COGENT,-73.9715,40.7585,2020/5/24 11:50
2259,United States,"New York, NY",605 Third Ave,John Wiley Bldg,10158,COB,COGENT,-73.9752,40.749,2020/5/24 11:50
2260,United States,"New York, NY",61 Broadway,61 Broadway,10006,COB,COGENT,-74.0126,40.7071,2020/5/24 11:50
2261,United States,"New York, NY",620 Eighth Avenue,New York Times Building,10018,COB,COGENT,-73.9902,40.7561,2020/5/24 11:50
2262,United States,"New York, NY",622 Third Ave,Grand Centeral Plaza,10017,COB,COGENT,-73.9757,40.7498,2020/5/24 11:50
2263,United States,"New York, NY",623 Fifth Avenue,Swiss Bank Tower,10022,COB,COGENT,-73.9764,40.7578,2020/5/24 11:50
2264,United States,"New York, NY",625 Madison Avenue,625 Madison Avenue,10022,COB,COGENT,-73.9713,40.7631,2020/5/24 11:50
2265,United States,"New York, NY",630 Fifth Ave,45 Rockefeller (International Building),10111,COB,COGENT,-73.9773,40.759,2020/5/24 11:50
2266,United States,"New York, NY",630 Ninth Ave,Film & Video Center Bldg,10036,COB,COGENT,-73.9912,40.76,2020/5/24 11:50
2267,United States,"New York, NY",630 Third Ave,630 Third Ave,10017,COB,COGENT,-73.9754,40.75,2020/5/24 11:50
2268,United States,"New York, NY",641 Lexington Ave,641 Lexington Ave,10022,COB,COGENT,-73.9696,40.7589,2020/5/24 11:50
2269,United States,"New York, NY",65 Broadway - 3rd floor,365 Data Centers (fka. Equinix NY10 - S&D),10006,CNDC,COGENT,-73.8594,41.0741,2020/5/24 11:50
2270,United States,"New York, NY",65 E 55th St,Park Avenue Tower,10022,COB,COGENT,-73.9721,40.761,2020/5/24 11:50
2271,United States,"New York, NY",650 Fifth Ave,650 Fifth Ave,10019,COB,COGENT,-73.9768,40.7597,2020/5/24 11:50
2272,United States,"New York, NY",650 Madison Ave,650 Madison Ave,10022,COB,COGENT,-73.9713,40.764,2020/5/24 11:50
2273,United States,"New York, NY",655 Third Ave,TAMS Bldg,10017,COB,COGENT,-73.9744,40.7503,2020/5/24 11:50
2274,United States,"New York, NY",660 Madison Ave,660 Madison Ave,10065,COB,COGENT,-73.971,40.7646,2020/5/24 11:50
2275,United States,"New York, NY",666 Fifth Ave,666 Fifth Ave,10103,COB,COGENT,-73.9769,40.7605,2020/5/24 11:50
2276,United States,"New York, NY",666 Third Ave,Chrysler Bldg East,10017,COB,COGENT,-73.9746,40.7512,2020/5/24 11:50
2277,United States,"New York, NY",667 Madison Ave,667 Madison Ave,10065,COB,COGENT,-73.9702,40.7644,2020/5/24 11:50
2278,United States,"New York, NY",675 Third Ave,675 Third Ave,10017,COB,COGENT,-73.9739,40.7509,2020/5/24 11:50
2279,United States,"New York, NY",685 Third Ave,685 Third Ave,10017,COB,COGENT,-73.9736,40.7515,2020/5/24 11:50
2280,United States,"New York, NY",7 Times Square,Times Square Tower,10036,COB,COGENT,-73.9869,40.7555,2020/5/24 11:50
2281,United States,"New York, NY",711 Third Ave,711 Third Ave,10017,COB,COGENT,-73.973,40.7522,2020/5/24 11:50
2282,United States,"New York, NY",712 Fifth Ave,712 Fifth Ave,10019,COB,COGENT,-73.975,40.7623,2020/5/24 11:50
2283,United States,"New York, NY",717 Fifth Ave,Merrill Lynch Bldg,10022,COB,COGENT,-73.9741,40.7619,2020/5/24 11:50
2284,United States,"New York, NY",733 Third Ave,Sun America Bldg,10017,COB,COGENT,-73.9722,40.7529,2020/5/24 11:50
2285,United States,"New York, NY",745 Fifth Ave,745 Fifth Ave,10151,COB,COGENT,-73.9734,40.7632,2020/5/24 11:50
2286,United States,"New York, NY",747 Third Ave,747 Third Ave,10017,COB,COGENT,-73.9721,40.7535,2020/5/24 11:50
2287,United States,"New York, NY",75 Broad St,New York Telecom Exchange,10004,COB,COGENT,-74.0113,40.7075,2020/5/24 11:50
2288,United States,"New York, NY",75 Broad St - 19th Floor,vXchnge (19th floor),10004,CNDC,COGENT,-74.0113,40.7075,2020/5/24 11:50
2289,United States,"New York, NY",75 Ninth Avenue,The Chelsea Market,10011,COB,COGENT,-74.006,40.7425,2020/5/24 11:50
2290,United States,"New York, NY",75 Varick Street,1 Hudson Square,10013,COB,COGENT,-74.0068,40.7235,2020/5/24 11:50
2291,United States,"New York, NY",750 Lexington Ave,International Plaza,10022,COB,COGENT,-73.9683,40.7627,2020/5/24 11:50
2292,United States,"New York, NY",750 Seventh Ave,Seventh Avenue Center,10019,COB,COGENT,-73.9838,40.7612,2020/5/24 11:50
2293,United States,"New York, NY",750 Third Ave,Grand Central Square,10017,COB,COGENT,-73.9726,40.7537,2020/5/24 11:50
2294,United States,"New York, NY",757 Third Ave,757 Third Ave,10017,COB,COGENT,-73.9715,40.7539,2020/5/24 11:50
2295,United States,"New York, NY",767 Fifth Ave,General Motors Bldg,10153,COB,COGENT,-73.9727,40.7636,2020/5/24 11:50
2296,United States,"New York, NY",767 Third Ave,767 Third Ave,10017,COB,COGENT,-73.973,40.7549,2020/5/24 11:50
2297,United States,"New York, NY",777 Third Ave,777 Third Ave,10017,COB,COGENT,-73.9711,40.7547,2020/5/24 11:50
2298,United States,"New York, NY",780 Third Avenue,780 Third Avenue,10017,COB,COGENT,-73.9715,40.755,2020/5/24 11:50
2299,United States,"New York, NY",787 Seventh Ave,Equitable Bldg,10019,COB,COGENT,-73.9822,40.7619,2020/5/24 11:50
2300,United States,"New York, NY",79 Madison Ave,79 Madison Ave,10016,COB,COGENT,-73.9853,40.7442,2020/5/24 11:50
2301,United States,"New York, NY",80 Pine St,80 Pine St,10005,COB,COGENT,-73.9022,42.2522,2020/5/24 11:50
2302,United States,"New York, NY",800 Third Ave,800 Third Ave,10022,COB,COGENT,-73.9714,40.7556,2020/5/24 11:50
2303,United States,"New York, NY",805 Third Ave,Crystal Pavilion,10022,COB,COGENT,-73.9706,40.7554,2020/5/24 11:50
2304,United States,"New York, NY",810 Seventh Ave,810 Seventh Ave,10019,COB,COGENT,-73.9826,40.7629,2020/5/24 11:50
2305,United States,"New York, NY",825 Eighth Ave,One Worldwide Plaza,10019,COB,COGENT,-73.9869,40.7621,2020/5/24 11:50
2306,United States,"New York, NY",841 Broadway,Roosevelt Building,10003,COB,COGENT,-73.9912,40.7343,2020/5/24 11:50
2307,United States,"New York, NY",845 Third Ave,845 Third Ave,10022,COB,COGENT,-73.9698,40.7566,2020/5/24 11:50
2308,United States,"New York, NY",85 10th Ave,Telehouse Chelsea,10011,CNDC,COGENT,-74.0079,40.7433,2020/5/24 11:50
2309,United States,"New York, NY",85 10th Ave,CenturyLink (fka. Level 3),10011,CNDC,COGENT,-74.0079,40.7433,2020/5/24 11:50
2310,United States,"New York, NY",850 Third Ave,850 Third Ave,10022,COB,COGENT,-73.9703,40.7569,2020/5/24 11:50
2311,United States,"New York, NY",853 Broadway,853 Broadway,10003,COB,COGENT,-73.9911,40.7347,2020/5/24 11:50
2312,United States,"New York, NY",875 Third Ave,875 Third Ave,10022,COB,COGENT,-73.9692,40.7572,2020/5/24 11:50
2313,United States,"New York, NY",885 Second Ave,1 Dag Hammarskjold Plz,10017,COB,COGENT,-73.9699,40.7535,2020/5/24 11:50
2314,United States,"New York, NY",885 Third Ave,Lipstick Bldg,10022,COB,COGENT,-73.9685,40.758,2020/5/24 11:50
2315,United States,"New York, NY",888 Seventh Ave,888 Seventh Ave,10106,COB,COGENT,-73.9809,40.7655,2020/5/24 11:50
2316,United States,"New York, NY",9 W 57th St,Solow Bldg,10019,COB,COGENT,-73.9749,40.7637,2020/5/24 11:50
2317,United States,"New York, NY",90 John St,The Renaissance,10038,COB,COGENT,-74.0064,40.7082,2020/5/24 11:50
2318,United States,"New York, NY",90 Park Ave,90 Park Ave,10016,COB,COGENT,-73.9794,40.7508,2020/5/24 11:50
2319,United States,"New York, NY",900 Third Ave,900 Third Ave,10022,COB,COGENT,-73.9692,40.7587,2020/5/24 11:50
2320,United States,"New York, NY",909 Third Ave,909 Third Ave,10022,COB,COGENT,-73.968,40.7583,2020/5/24 11:50
2321,United States,"New York, NY",919 Third Ave,919 Third Ave,10022,COB,COGENT,-73.9675,40.7589,2020/5/24 11:50
2322,United States,"New York, NY",99 Park Avenue,99 Park Avenue,10016,COB,COGENT,-73.9783,40.7501,2020/5/24 11:50
2323,United States,"New York, NY",One Bryant Park,Bank of America Tower,10036,COB,COGENT,-73.9849,40.7556,2020/5/24 11:50
2324,United States,"NEWARK, NJ",1085 Raymond Blvd,One Newark Center,7102,COB,COGENT,-74.1663,40.7365,2020/5/24 11:50
2325,United States,"Newark, NJ",165 Halsey St,165 Halsey,7102,COB,COGENT,-74.1734,40.7368,2020/5/24 11:50
2326,United States,"Newark, NJ",165 Halsey St - 5th Floor,zColo (fka. Fibernet/gateway - 5th floor),7102,CNDC,COGENT,-74.1734,40.7368,2020/5/24 11:50
2328,United States,"Newark, NJ",165 Halsey Street,Equinix NY1,7102,CNDC,COGENT,-74.1734,40.7368,2020/5/24 11:50
2329,United States,"Newark, NJ",7-45 Raymond Boulevard,1 Gateway Center,7102,COB,COGENT,-74.1189,40.7334,2020/5/24 11:50
2330,United States,"Newark, NJ",744 Broad St,National Newark Building,7102,COB,COGENT,-74.1708,40.7367,2020/5/24 11:50
2331,United States,"Norcross, GA",2775 Northwoods Parkway NW,Flexential ATL01-02,30071,CNDC,COGENT,-84.1986,33.9548,2020/5/24 11:50
2332,United States,"Norristown, PA",1000 Adams Ave,Tierpoint VFO,19403,CNDC,COGENT,-75.4211,40.1216,2020/5/24 11:50
2333,United States,"North Bergen, NJ",5851 Westside Ave,Equinix NY7 (fka. S&D),7047,CNDC,COGENT,-74.0311,40.7969,2020/5/24 11:50
2334,United States,"Norwalk, CT",6 Norden Place,CyrusOne Norwalk,6855,CNDC,COGENT,-73.3915,41.1116,2020/5/24 11:50
2335,United States,"Oak Brook, IL",1111 W 22nd St,AT&T Plaza,60523,COB,COGENT,-87.946,41.8464,2020/5/24 11:50
2336,United States,"Oak Brook, IL",1808 Swift Drive,zColo (fka. Latisys),60523,CNDC,COGENT,-87.9209,41.854,2020/5/24 11:50
2337,United States,"Oak Brook, IL","800 Jorie Blvd, Suite 120",Navisite,60523,CNDC,COGENT,-87.9415,41.8453,2020/5/24 11:50
2338,United States,"Oakland, CA",1624 Franklin Street,Cogent Data Center - Oakland,94612,CDC,COGENT,-122.268,37.8058,2020/5/24 11:50
2339,United States,"Oakland, CA",1624 Franklin Street,Franklin Building,94612,COB,COGENT,-122.268,37.8058,2020/5/24 11:50
2341,United States,"Oklahoma City, OK",201 Robert S. Kerr Avenue,FullTel Communications,73102,CNDC,COGENT,-97.5169,35.4704,2020/5/24 11:50
2342,United States,"Oklahoma City, OK",4114 Perimeter Center Place,Tierpoint OK2,73112,CNDC,COGENT,-97.5951,35.514,2020/5/24 11:50
2343,United States,"Oklahoma City, OK",4121 Perimeter Center Place,Tierpoint OKC (fka. Perimeter Technology),73122,CNDC,COGENT,-97.5968,35.5142,2020/5/24 11:50
2344,United States,"Oklahoma City, OK",7725 W Reno Ave,RACK59 Data Center,73127,CNDC,COGENT,-97.6493,35.4665,2020/5/24 11:50
2345,United States,"Omaha, NE",1299 Farnam St,1200 Landmark Center,68102,COB,COGENT,-95.9327,41.2571,2020/5/24 11:50
2346,United States,"Omaha, NE",1623 Farnam Street,Building MMR,68102,CNDC,COGENT,-95.9381,41.2574,2020/5/24 11:50
2347,United States,"Omaha, NE",1623 Farnam Street,Great Plains (fka. Pinpoint),68102,CNDC,COGENT,-95.9381,41.2574,2020/5/24 11:50
2348,United States,"Omaha, NE",1700 Farnam Street,Woodman Tower,68102,COB,COGENT,-95.9389,41.2579,2020/5/24 11:50
2349,United States,"Omaha, NE",222 S. 15th Street CPP-I,Central Park Plaza I,68102,COB,COGENT,-95.9363,41.2584,2020/5/24 11:50
2350,United States,"Omaha, NE",222 S. 15th Street CPP-II,Central Park Plaza 2,68102,COB,COGENT,-90.2028,38.6275,2020/5/24 11:50
2351,United States,"Orangeburg, NY",1 Ramland Rd,Green House Data,10962,CNDC,COGENT,-73.9751,41.0345,2020/5/24 11:50
2352,United States,"Orlando, FL",100 West Lucerne Circle,Colo Solutions,32801,CNDC,COGENT,-81.3804,28.5347,2020/5/24 11:50
2353,United States,"Orlando, FL",200 S Orange Ave (and 250 S Orange),SunTrust Center,32801,COB,COGENT,-81.3781,28.5401,2020/5/24 11:50
2354,United States,"Orlando, FL",201 E Pine St,Capital Plaza 1,32801,COB,COGENT,-81.3749,28.5418,2020/5/24 11:50
2355,United States,"Orlando, FL",201 S Orange Ave,Signature Plaza,32801,COB,COGENT,-81.3781,28.5401,2020/5/24 11:50
2356,United States,"Orlando, FL",225 E Robinson St,Landmark Center 2,32801,COB,COGENT,-81.3745,28.5462,2020/5/24 11:50
2357,United States,"Orlando, FL",255 S Orange Ave,The Citrus Center,32801,COB,COGENT,-81.3785,28.5394,2020/5/24 11:50
2358,United States,"Orlando, FL",300 S. Orange Avenue,Lincoln Plaza,32801,COB,COGENT,-81.3794,28.5388,2020/5/24 11:50
2359,United States,"Orlando, FL",301 E Pine St,Capital Plaza 2,32801,COB,COGENT,-81.3741,28.5417,2020/5/24 11:50
2360,United States,"Orlando, FL",315 E Robinson St,Landmark Center 1,32801,COB,COGENT,-81.3735,28.5461,2020/5/24 11:50
2361,United States,"Orlando, FL",390 N Orange Ave,Bank of America Center,32801,COB,COGENT,-81.3798,28.5382,2020/5/24 11:50
2362,United States,"Orlando, FL",420 S Orange Avenue,CNL Center II at City Commons,32801,COB,COGENT,-81.3798,28.5383,2020/5/24 11:50
2363,United States,"Orlando, FL",440 West Kennedy Blvd.,Hostdime,32810,CNDC,COGENT,-81.3932,28.6175,2020/5/24 11:50
2364,United States,"Orlando, FL",440 West Kennedy Blvd.,Atlantic.net,32810,CNDC,COGENT,-81.3932,28.6175,2020/5/24 11:50
2365,United States,"Orlando, FL",450 S Orange Ave,CNL Center l,32801,COB,COGENT,-81.3793,28.5371,2020/5/24 11:50
2366,United States,"Orlando, FL",800 N Magnolia Ave,One Orlando Centre,32803,COB,COGENT,-81.3775,28.5555,2020/5/24 11:50
2367,United States,"Orlando, FL",9701 S. John Young Pkwy,DataSite,32819,CNDC,COGENT,-81.4198,28.4271,2020/5/24 11:50
2368,United States,"Overland Park, KS",6800 College Boulevard,Financial Plaza II,66212,COB,COGENT,-94.6645,38.9285,2020/5/24 11:50
2369,United States,"Overland Park, KS",6900 College Boulevard,Financial Plaza III,66210,COB,COGENT,-94.6653,38.9284,2020/5/24 11:50
2370,United States,"Overland Park, KS",7101 College Blvd,7101 Tower,66210,COB,COGENT,-94.6688,38.9262,2020/5/24 11:50
2371,United States,"Overland Park, KS",7300 College Blvd,Lighton Plaza I,66210,COB,COGENT,-94.672,38.9294,2020/5/24 11:50
2372,United States,"Overland Park, KS",7300 W 110th St,Commerce Plaza I,66210,COB,COGENT,-94.6706,38.9311,2020/5/24 11:50
2373,United States,"Overland Park, KS",7400 College Blvd,Lighton Plaza II,66210,COB,COGENT,-94.6721,38.9287,2020/5/24 11:50
2374,United States,"Overland Park, KS",7400 West 110th Street,Commerce Plaza II,66210,COB,COGENT,-94.6716,38.9313,2020/5/24 11:50
2375,United States,"Overland Park, KS",7500 College Blvd,Lighton Tower,66210,COB,COGENT,-94.6732,38.9291,2020/5/24 11:50
2376,United States,"Palo Alto, CA",525 University Ave,Palo Alto Office Center,94301,COB,COGENT,-122.159,37.4487,2020/5/24 11:50
2377,United States,"Palo Alto, CA",529 Bryant St,Equinix SV8 (fka. S&D),94301,CNDC,COGENT,-122.161,37.4458,2020/5/24 11:50
2378,United States,"Papillion, NE",11425 S 84th St,Tierpoint-MID (Midlands),68046,CNDC,COGENT,-96.039,41.1345,2020/5/24 11:50
2379,United States,"Parsippany, NJ",200 Webro Rd,Cologix NNJ3,7054,CNDC,COGENT,-74.4189,40.8571,2020/5/24 11:50
2380,United States,"Pasadena, CA",201 S Lake Ave,201 S Lake,91101,COB,COGENT,-118.133,34.1422,2020/5/24 11:50
2381,United States,"Pasadena, CA",225 S Lake Ave,225 S Lake Ave,91101,COB,COGENT,-118.134,34.1419,2020/5/24 11:50
2382,United States,"Pasadena, CA",251 S Lake Ave,Corporate Center Pasadena,91101,COB,COGENT,-118.132,34.1417,2020/5/24 11:50
2383,United States,"Pasadena, CA",2947 Bradley Street,Bradley Street Buildings / 2947,91107,COB,COGENT,-118.089,34.1687,2020/5/24 11:50
2384,United States,"Pasadena, CA",2947 Bradley Street 1st Floor,Cogent Data Center - Pasadena,91107,CDC,COGENT,-118.089,34.1687,2020/5/24 11:50
2385,United States,"Pasadena, CA",55 S Lake Ave,Pasadena Tower 2,91101,COB,COGENT,-118.133,34.1449,2020/5/24 11:50
2386,United States,"Pasadena, CA",800 E Colorado Blvd,Tower 1,91101,COB,COGENT,-118.133,34.1456,2020/5/24 11:50
2387,United States,"Philadelphia, PA",1 S Broad St,One South Broad,19107,COB,COGENT,-75.1635,39.9512,2020/5/24 11:50
2388,United States,"Philadelphia, PA",1 S Penn Square (aka 1339 Chestnut St),The Widener Bldg,19107,COB,COGENT,-75.1635,39.951,2020/5/24 11:50
2389,United States,"Philadelphia, PA",100 N 18th St,Two Logan Square,19103,COB,COGENT,-75.1698,39.9557,2020/5/24 11:50
2391,United States,"Philadelphia, PA",123 S Broad St,Wells Fargo Building,19109,COB,COGENT,-75.1634,39.9498,2020/5/24 11:50
2392,United States,"Philadelphia, PA",130 N 18th St,One Logan Square,19103,COB,COGENT,-75.17,39.9564,2020/5/24 11:50
2393,United States,"Philadelphia, PA",1500 John F Kennedy Blvd,2 Penn Center Plaza,19102,COB,COGENT,-75.1659,39.9535,2020/5/24 11:50
2394,United States,"Philadelphia, PA",1500 Market St,Centre Square,19102,COB,COGENT,-90.2176,29.9915,2020/5/24 11:50
2395,United States,"Philadelphia, PA",1500 Spring Garden St,SunGard phl-1500,19130,CNDC,COGENT,-75.1641,39.9621,2020/5/24 11:50
2396,United States,"Philadelphia, PA",1500 Spring Garden St,vXchnge,19130,CNDC,COGENT,-75.1641,39.9621,2020/5/24 11:50
2397,United States,"Philadelphia, PA",1500 Spring Garden St,1500 Spring Garden,19130,COB,COGENT,-75.1641,39.9621,2020/5/24 11:50
2398,United States,"Philadelphia, PA",1515 Market St,1515 Market St,19102,COB,COGENT,-75.1659,39.953,2020/5/24 11:50
2399,United States,"Philadelphia, PA",1600 John F Kennedy Blvd,4 Penn Center,19103,COB,COGENT,-75.1671,39.9536,2020/5/24 11:50
2400,United States,"Philadelphia, PA",1600 Market St,PNC Bank Bldg,19103,COB,COGENT,-75.1674,39.9525,2020/5/24 11:50
2401,United States,"Philadelphia, PA",1601 Cherry Street,3 Parkway,19102,COB,COGENT,-75.1669,39.9559,2020/5/24 11:50
2402,United States,"Philadelphia, PA",1601 Chestnut St (50 South 16th St),Two Liberty Place,19103,COB,COGENT,-75.1672,39.9517,2020/5/24 11:50
2403,United States,"Philadelphia, PA",1601 Market St,1601 Market St,19103,COB,COGENT,-75.1674,39.9531,2020/5/24 11:50
2404,United States,"Philadelphia, PA",1617 John F Kennedy Blvd,One Penn Center at Suburban Station,19103,COB,COGENT,-75.1679,39.9543,2020/5/24 11:50
2405,United States,"Philadelphia, PA",1628 John F Kennedy Blvd,8 Penn Center,19103,COB,COGENT,-75.1681,39.9537,2020/5/24 11:50
2406,United States,"Philadelphia, PA",1635 Market St,7 Penn Center,19103,COB,COGENT,-75.1679,39.9533,2020/5/24 11:50
2407,United States,"Philadelphia, PA",1650 Arch St,1650 Arch St,19103,COB,COGENT,-75.1678,39.9547,2020/5/24 11:50
2408,United States,"Philadelphia, PA",1650 Market St,1 Liberty Place,19103,COB,COGENT,-75.1682,39.9527,2020/5/24 11:50
2409,United States,"Philadelphia, PA",1700 Market St,1700 Market St,19103,COB,COGENT,-75.1696,39.9528,2020/5/24 11:50
2410,United States,"Philadelphia, PA",1717 Arch St,Three Logan Square (Formerly Bell Atlantic Tower),19103,COB,COGENT,-75.1693,39.9553,2020/5/24 11:50
2411,United States,"Philadelphia, PA",1735 Market St,Mellon Bank Center,19103,COB,COGENT,-75.1696,39.9537,2020/5/24 11:50
2412,United States,"Philadelphia, PA",1800 John F Kennedy Blvd,1800 John F Kennedy Blvd,19103,COB,COGENT,-75.17,39.954,2020/5/24 11:50
2413,United States,"Philadelphia, PA",1801 Market St,10 Penn Center,19103,COB,COGENT,-75.1703,39.9536,2020/5/24 11:50
2414,United States,"Philadelphia, PA",1818 Market St,1818 Beneficial Bank Place,19103,COB,COGENT,-75.171,39.9529,2020/5/24 11:50
2415,United States,"Philadelphia, PA",1835 Market St,11 Penn Center,19103,COB,COGENT,-75.1711,39.9537,2020/5/24 11:50
2416,United States,"Philadelphia, PA",1845 Walnut St,1845 Walnut Street,19103,COB,COGENT,-75.1715,39.9507,2020/5/24 11:50
2417,United States,"Philadelphia, PA",1880 John F Kennedy Blvd,1880 John F Kennedy Blvd,19103,COB,COGENT,-75.171,39.9541,2020/5/24 11:50
2418,United States,"Philadelphia, PA",1900 Market St,Philadelphia Stock Exchange Bldg,19103,COB,COGENT,-77.5828,35.2834,2020/5/24 11:50
2419,United States,"Philadelphia, PA",2000 Market St,2000 Market St,19103,COB,COGENT,-75.1736,39.9533,2020/5/24 11:50
2420,United States,"Philadelphia, PA",2001 Market St,Two Commerce Square,19103,COB,COGENT,-75.1737,39.9541,2020/5/24 11:50
2421,United States,"Philadelphia, PA",2005 Market St,One Commerce Square,19103,COB,COGENT,-75.1745,39.9546,2020/5/24 11:50
2422,United States,"Philadelphia, PA",2401 Locust St,Quonix,19103,CNDC,COGENT,-75.18,39.9503,2020/5/24 11:50
2423,United States,"Philadelphia, PA",2929 Arch Street,Cira Center,19104,COB,COGENT,-75.1817,39.9574,2020/5/24 11:50
2424,United States,"Philadelphia, PA",2929 Walnut Street,FMC Tower at Cira Centre,19104,COB,COGENT,-75.184,39.9521,2020/5/24 11:50
2425,United States,"Philadelphia, PA",30 S 17th St,United Plaza,19103,COB,COGENT,-75.1692,39.9522,2020/5/24 11:50
2426,United States,"Philadelphia, PA",3624 Market St,3624 Market St,19104,COB,COGENT,-75.1953,39.9558,2020/5/24 11:50
2427,United States,"Philadelphia, PA",3675 Market Street,University City Science Center,19104,COB,COGENT,-75.1952,39.9567,2020/5/24 11:50
2428,United States,"Philadelphia, PA",3701 Market St,365 Data Centers (fka. Equinix PH2 - S&D),19104,CNDC,COGENT,-75.1962,39.9568,2020/5/24 11:50
2429,United States,"Philadelphia, PA",3701 Market St,Port of Technology Bldg,19104,COB,COGENT,-75.1962,39.9568,2020/5/24 11:50
2430,United States,"Philadelphia, PA",3711 Market Street,The Science Center Incubator,19104,COB,COGENT,-75.1976,39.9572,2020/5/24 11:50
2432,United States,"Philadelphia, PA",401 N Broad St - 3rd Floor,CenturyLink (fka. Level 3),19108,CNDC,COGENT,-75.1625,39.9564,2020/5/24 11:50
2433,United States,"Philadelphia, PA",401 N Broad Street,Equinix PH1 (fka. S&D),19108,CNDC,COGENT,-75.1617,39.9599,2020/5/24 11:50
2434,United States,"Philadelphia, PA",401 N Broad Street,Netrality,19108,CNDC,COGENT,-75.1617,39.9599,2020/5/24 11:50
2435,United States,"Philadelphia, PA",401 N. Broad Street,Crown Castle (fka. Lightower),19108,CNDC,COGENT,-75.1617,39.9599,2020/5/24 11:50
2437,United States,"Philadelphia, PA",4775 League Island Blvd,Tierpoint PHL,19112,CNDC,COGENT,-75.1654,39.8925,2020/5/24 11:50
2438,United States,"Philadelphia, PA",510 Walnut St,Penn Mutual Towers,19103,COB,COGENT,-92.0976,37.1234,2020/5/24 11:50
2439,United States,"Philadelphia, PA",520 Walnut St,520 Walnut St,19106,COB,COGENT,-87.9499,41.9104,2020/5/24 11:50
2440,United States,"Philadelphia, PA",530 Walnut St,Penn Mutual Towers,19106,COB,COGENT,-75.1501,39.9472,2020/5/24 11:50
2443,United States,"Philadelphia, PA",718 Arch St,718 Arch St,19106,COB,COGENT,-75.1523,39.9529,2020/5/24 11:50
2444,United States,"Philadelphia, PA",833 Chestnut St,HCP (fka. DRT MMR),19107,CNDC,COGENT,-75.1549,39.9501,2020/5/24 11:50
2445,United States,"Philadelphia, PA",833 Chestnut St,833 Chestnut East,19107,COB,COGENT,-75.1549,39.9501,2020/5/24 11:50
2446,United States,"Phoenix, AZ",101 N First Ave,US Bank Center,85003,COB,COGENT,-112.075,33.4499,2020/5/24 11:50
2447,United States,"Phoenix, AZ",120 E Van Buren St,Digital Realty (fka. TelX),85004,CNDC,COGENT,-112.072,33.452,2020/5/24 11:50
2448,United States,"Phoenix, AZ",1850 N Central Ave,Viad Corporate Center,85004,COB,COGENT,-112.075,33.4684,2020/5/24 11:50
2449,United States,"Phoenix, AZ",1850 W Deer Valley Rd,Flexential Deer Valley (fka. Peak 10 / ViaWest),85027,CNDC,COGENT,-112.098,33.6847,2020/5/24 11:50
2450,United States,"Phoenix, AZ",2 N Central Ave,One Renaissance Square,85004,COB,COGENT,-112.074,33.4485,2020/5/24 11:50
2451,United States,"Phoenix, AZ",201 E Washington St,Collier Center,85004,COB,COGENT,-112.071,33.4475,2020/5/24 11:50
2452,United States,"Phoenix, AZ",201 N Central Ave,Chase Tower,85004,COB,COGENT,-112.073,33.4509,2020/5/24 11:50
2453,United States,"Phoenix, AZ",2201 E Camelback Rd,Anchor Centre West,85016,COB,COGENT,-112.034,33.5086,2020/5/24 11:50
2454,United States,"Phoenix, AZ",2231 E Camelback Rd,Anchor Centre East,85016,COB,COGENT,-112.033,33.5089,2020/5/24 11:50
2455,United States,"Phoenix, AZ",2325 E Camelback Road,24th At Camelback II,85016,COB,COGENT,-112.032,33.5086,2020/5/24 11:50
2456,United States,"Phoenix, AZ",2375 E Camelback Rd,24th at Camelback I,85016,COB,COGENT,-112.031,33.5087,2020/5/24 11:50
2457,United States,"Phoenix, AZ",2390 E Camelback Rd,Biltmore Financial Center,85016,COB,COGENT,-112.032,33.5097,2020/5/24 11:50
2458,United States,"Phoenix, AZ",2394 E Camelback Rd,Biltmore Financial Center,85016,COB,COGENT,-112.032,33.5108,2020/5/24 11:50
2459,United States,"Phoenix, AZ",2398 E Camelback Rd,Biltmore Financial Center,85016,COB,COGENT,-112.031,33.5101,2020/5/24 11:50
2460,United States,"Phoenix, AZ",2415 E Camelback Rd,Camelback Esplanade III,85016,COB,COGENT,-112.029,33.5081,2020/5/24 11:50
2461,United States,"Phoenix, AZ",2425 E Camelback Rd,Camelback Esplanade I,85016,COB,COGENT,-112.028,33.5087,2020/5/24 11:50
2462,United States,"Phoenix, AZ",2500 West Union Hills Drive,Aligned Energy,85027,CNDC,COGENT,-112.111,33.6589,2020/5/24 11:50
2463,United States,"Phoenix, AZ",2525 E Camelback Rd,Camelback Esplanade II,85016,COB,COGENT,-112.028,33.5089,2020/5/24 11:50
2464,United States,"Phoenix, AZ",2555 E Camelback Rd,Camelback Esplanade V,85016,COB,COGENT,-112.026,33.509,2020/5/24 11:50
2465,United States,"Phoenix, AZ",2575 E Camelback Rd,Camelback Esplanade IV,85016,COB,COGENT,-112.026,33.5082,2020/5/24 11:50
2466,United States,"Phoenix, AZ",2600 N Central Ave,2600 Tower,85004,COB,COGENT,-112.074,33.4773,2020/5/24 11:50
2467,United States,"Phoenix, AZ",2700 N Central Ave,2700 Tower,85004,COB,COGENT,-112.074,33.478,2020/5/24 11:50
2468,United States,"Phoenix, AZ",2800 N Central Ave,2800 Tower,85004,COB,COGENT,-112.074,33.4787,2020/5/24 11:50
2469,United States,"Phoenix, AZ",3101-3111 N Central Avenue,3101-3111 North Central Avenue,85012,COB,COGENT,-112.073,33.4843,2020/5/24 11:50
2470,United States,"Phoenix, AZ",3110 N Central Avenue,zColo (fka. CoreLink),85012,CNDC,COGENT,-112.076,33.4846,2020/5/24 11:50
2471,United States,"Phoenix, AZ",3402 East University Drive,Phoenix NAP,85034,CNDC,COGENT,-112.009,33.416,2020/5/24 11:50
2472,United States,"Phoenix, AZ",3410 E University Drive,Cogent Data Center - Phoenix,85034,CDC,COGENT,-112.008,33.417,2020/5/24 11:50
2473,United States,"Phoenix, AZ",3410 E University Drive,Corporate Center At South Bank Bldg# 1,85034,COB,COGENT,-112.008,33.417,2020/5/24 11:50
2474,United States,"Phoenix, AZ",40 N Central Ave,Two Rennasiance Sq,85004,COB,COGENT,-112.075,33.449,2020/5/24 11:50
2476,United States,"Phoenix, AZ",615 N 48th Street,Iron Mountain (fka. IO Data Center),85008,CNDC,COGENT,-111.976,33.4545,2020/5/24 11:50
2477,United States,"Phoenix, AZ",615 N 48th Street,Cyxtera PHX1 (fka. CenturyLink PH1),85008,CNDC,COGENT,-111.976,33.4545,2020/5/24 11:50
2478,United States,"Piscataway, NJ",101 Possumtown Rd,QTS Datacenter,8854,CNDC,COGENT,-74.4847,40.5569,2020/5/24 11:50
2479,United States,"Piscataway, NJ",101 Possumtown Rd MER1,QTS Datacenter (MER1),8854,CNDC,COGENT,-74.4847,40.5569,2020/5/24 11:50
2480,United States,"Piscataway, NJ",15 Corporate Place S,Open Data Centers,8854,CNDC,COGENT,-74.4567,40.55,2020/5/24 11:50
2481,United States,"Piscataway, NJ",3 Corporate Place,Evoque PSWYNJ63,8854,CNDC,COGENT,-74.4569,40.5545,2020/5/24 11:50
2482,United States,"Piscataway, NJ",3 Corporate Place,Digital Realty,8854,CNDC,COGENT,-74.4569,40.5545,2020/5/24 11:50
2483,United States,"Piscataway, NJ",3 Corporate Place,Cyxtera NJ3,8854,CNDC,COGENT,-74.4569,40.5545,2020/5/24 11:50
2484,United States,"Piscataway, NJ",365 South Randolphville Road,Digital Realty,8854,CNDC,COGENT,-74.4595,40.5527,2020/5/24 11:50
2485,United States,"Piscataway, NJ",365 South Randolphville Road,Cyxtera NJ4,8854,CNDC,COGENT,-74.4595,40.5527,2020/5/24 11:50
2486,United States,"Pittsburgh, PA",1 PPG Place,One PPG Place,15222-5405,COB,COGENT,-80.0035,40.4399,2020/5/24 11:50
2487,United States,"Pittsburgh, PA",100 South Commons,DataBank CLE1 (fka. 365 Data Center),15212,CNDC,COGENT,-80.0048,40.4512,2020/5/24 11:50
2488,United States,"Pittsburgh, PA",1001 Liberty Ave,Federated Tower at Liberty Center,15222,COB,COGENT,-79.9946,40.4444,2020/5/24 11:50
2489,United States,"Pittsburgh, PA",11 Stanwix St,11 Stanwix St,15222,COB,COGENT,-80.0061,40.4391,2020/5/24 11:50
2490,United States,"Pittsburgh, PA",2 PPG Place,2 PPG Place,15222,COB,COGENT,-80.0032,40.4405,2020/5/24 11:50
2491,United States,"Pittsburgh, PA",20 Stanwix Street,PNC Center,15222,COB,COGENT,-80.0051,40.4387,2020/5/24 11:50
2492,United States,"Pittsburgh, PA",2202 Liberty Ave,vXchnge,15222,CNDC,COGENT,-79.9811,40.4515,2020/5/24 11:50
2493,United States,"Pittsburgh, PA",301 Grant Street,One Oxford Centre,15219,COB,COGENT,-79.9991,40.4383,2020/5/24 11:50
2494,United States,"Pittsburgh, PA","322 4th Ave, 4th Floor",DQE POP,15222,CNDC,COGENT,-80.0005,40.4389,2020/5/24 11:50
2495,United States,"Pittsburgh, PA",4 PPG Place,4 PPG Place,15222,COB,COGENT,-80.0035,40.4399,2020/5/24 11:50
2496,United States,"Pittsburgh, PA",401 Liberty Ave,Gateway Center III,15222,COB,COGENT,-80.0053,40.4419,2020/5/24 11:50
2497,United States,"Pittsburgh, PA",420 Fort Duquesne Blvd,Gateway Center I,15222,COB,COGENT,-80.0055,40.4428,2020/5/24 11:50
2498,United States,"Pittsburgh, PA",436 Seventh Avenue,Koppers Building,15219,COB,COGENT,-79.9956,40.4421,2020/5/24 11:50
2499,United States,"Pittsburgh, PA",444 Liberty Ave,Gateway Center IV,15222,COB,COGENT,-80.0046,40.4409,2020/5/24 11:50
2500,United States,"Pittsburgh, PA",5 PPG Place,5 PPG Place,15222,COB,COGENT,-80.0032,40.4392,2020/5/24 11:50
2501,United States,"Pittsburgh, PA",500 Grant Street,BNY Mellon Center,15219,COB,COGENT,-79.9958,40.4397,2020/5/24 11:50
2502,United States,"Pittsburgh, PA",525 William Penn Place,Three Mellon Center,15219,COB,COGENT,-79.9979,40.4401,2020/5/24 11:50
2503,United States,"Pittsburgh, PA",6 PPG Place,6 PPG Place,15222,COB,COGENT,-80.0039,40.4396,2020/5/24 11:50
2504,United States,"Pittsburgh, PA",600 Grant Street,US Steel Tower (UPMC),15219,COB,COGENT,-79.9948,40.4409,2020/5/24 11:50
2505,United States,"Pittsburgh, PA",603 Stanwix St,Gateway Center II,15222,COB,COGENT,-80.0046,40.4422,2020/5/24 11:50
2506,United States,"Pittsburgh, PA",625 Liberty Avenue,EQT Plaza,15222,COB,COGENT,-80.001,40.4425,2020/5/24 11:50
2507,United States,"Pittsburgh, PA",650 Smithfield St,Centre City Tower,15222,COB,COGENT,-81.2081,39.9101,2020/5/24 11:50
2508,United States,"Pittsburgh, PA",707 Grant Street,Gulf Tower,15219,COB,COGENT,-79.9953,40.4425,2020/5/24 11:50
2509,United States,"Plano, TX",1009 Jupiter Rd.,Global IP Networks,75074,CNDC,COGENT,-96.6831,33.0122,2020/5/24 11:50
2510,United States,"Plano, TX",1221 Coit Rd,Internap,75075,CNDC,COGENT,-96.7684,33.0157,2020/5/24 11:50
2511,United States,"PLANO, TX",1500 EAST PLANO PARKWAY,Ascent DAL1,75074,CNDC,COGENT,-96.6917,33.0064,2020/5/24 11:50
2512,United States,"Plano, TX",2800 Summit Avenue,Aligned Energy,75074,CNDC,COGENT,-96.6764,33.0093,2020/5/24 11:50
2513,United States,"Plano, TX",3500 East Plano Parkway,3500 East Plano Parkway,75074,CNDC,COGENT,-96.6663,33.0074,2020/5/24 11:50
2514,United States,"Plano, TX",5700 Granite Parkway,Granite Park Two,75024,COB,COGENT,-96.8189,33.0875,2020/5/24 11:50
2515,United States,"Plano, TX",5800 Granite Parkway,Granite Park One,75024,COB,COGENT,-96.8202,33.0876,2020/5/24 11:50
2516,United States,"Plano, TX",5830 Granite Parkway,Granite Park Five,75024,COB,COGENT,-96.8211,33.0876,2020/5/24 11:50
2517,United States,"Plano, TX",5850 Granite Parkway,Granite Park Four,75024,COB,COGENT,-96.8218,33.0872,2020/5/24 11:50
2518,United States,"Plano, TX",6653 Pinecrest Drive,T5 Data Centers,75024,CNDC,COGENT,-96.811,33.0649,2020/5/24 11:50
2519,United States,"Plano, TX",6653 Pinecrest Drive,Equinix DA7,75024,CNDC,COGENT,-96.811,33.0649,2020/5/24 11:50
2520,United States,"Plano, TX",8375 Dominion Parkway,Data Bank DFW3,75024,CNDC,COGENT,-96.8156,33.0871,2020/5/24 11:50
2521,United States,"Pompano Beach, FL",500 Green Road,Volico Data Center,33064,CNDC,COGENT,-80.13,26.2897,2020/5/24 11:50
2522,United States,"Portland, OR",1233 NW 12th Avenue,vXchnge,97209,CNDC,COGENT,-122.684,45.5317,2020/5/24 11:50
2523,United States,"Portland, OR",707 SW Washington,Union Bank of California Tower,97205,COB,COGENT,-122.679,45.5211,2020/5/24 11:50
2524,United States,"Portland, OR",707 SW Washington St,Building MMR,97205,CNDC,COGENT,-122.679,45.5211,2020/5/24 11:50
2525,United States,"Portland, OR",921 SW Washington St,Pittock Bldg,97205,CNDC,COGENT,-122.681,45.5214,2020/5/24 11:50
2526,United States,"Portland, OR",921 SW Washington St,The Pittock Block Building,97205,COB,COGENT,-122.681,45.5214,2020/5/24 11:50
2527,United States,"Providence, RI",1155 Westminster St,Prov.net PVD1,2909,CNDC,COGENT,-71.4277,41.8175,2020/5/24 11:50
2529,United States,"Providence, RI",304 Carpenter Street,Prov.net PVD2,2909,CNDC,COGENT,-71.4291,41.8178,2020/5/24 11:50
2530,United States,"Purchase, NY",1 Manhattanville Rd,Centre 1,10577,COB,COGENT,-73.7193,41.0248,2020/5/24 11:50
2531,United States,"Purchase, NY",1 Manhattanville Rd 1st Floor,RFA Data Center,10577,CNDC,COGENT,-73.7193,41.0248,2020/5/24 11:50
2532,United States,"Purchase, NY",2 Manhattanville Road,Centre 2,10577,COB,COGENT,-73.7197,41.026,2020/5/24 11:50
2533,United States,"Purchase, NY",3 Manhattanville Road,Centre 3,10577,COB,COGENT,-73.7196,41.0259,2020/5/24 11:50
2534,United States,"Purchase, NY",4 Manhattanville Road,Centre 4,10577,COB,COGENT,-73.7203,41.0253,2020/5/24 11:50
2535,United States,"Raleigh, NC",213 N Harrington St,Spectrum (fka. TWC/DukeNet),27603,CNDC,COGENT,-78.6452,35.7835,2020/5/24 11:50
2536,United States,"Rancho Cordova, CA",11085 Sun Center Dr,SunGard Sacramento,95670,CNDC,COGENT,-121.272,38.5961,2020/5/24 11:50
2537,United States,"Rancho Cordova, CA",2999 Gold Canal Dr,DataCate,95670,CNDC,COGENT,-121.268,38.5944,2020/5/24 11:50
2538,United States,"Reading, PA",2561 Bernville Road,DirectLink Technologies Datacenter,19605,CNDC,COGENT,-75.982,40.3813,2020/5/24 11:50
2539,United States,"Redondo Beach, CA",3690 Redondo Beach Avenue,Internap LAX014,90278,CNDC,COGENT,-118.369,33.8892,2020/5/24 11:50
2540,United States,"Reston, VA",10740 Parkridge Blvd,Parkridge 6,20191,COB,COGENT,-77.3133,38.9462,2020/5/24 11:50
2541,United States,"Reston, VA",10780-10790 Parkridge Blvd,Parkridge 5,20191,COB,COGENT,-77.316,38.9456,2020/5/24 11:50
2542,United States,"Reston, VA",11513 Sunset Hills Rd,Flex Data Center (fka. InfoRelay),20190,CNDC,COGENT,-77.3466,38.9519,2020/5/24 11:50
2543,United States,"Reston, VA",11700 Plaza America Dr,Plaza America Tower 1,20190,COB,COGENT,-77.3499,38.9526,2020/5/24 11:50
2544,United States,"Reston, VA",11710 Plaza America Dr,Plaza America Tower 2,20190,COB,COGENT,-77.3509,38.9524,2020/5/24 11:50
2545,United States,"Reston, VA",11911 Freedom Drive,One Fountain,20190,COB,COGENT,-77.3564,38.9591,2020/5/24 11:50
2546,United States,"Reston, VA",11921 Freedom Drive,Two Fountain,20190,COB,COGENT,-77.3578,38.959,2020/5/24 11:50
2548,United States,"Reston, VA",11951 Freedom Dr,One Freedom Square,20190,COB,COGENT,-77.3586,38.9594,2020/5/24 11:50
2549,United States,"Reston, VA",11955 Democracy Drive,College Board Building,20190,COB,COGENT,-77.3586,38.9579,2020/5/24 11:50
2550,United States,"Reston, VA",11955 Freedom Dr,Two Freedom Square,20190,COB,COGENT,-77.3586,38.9594,2020/5/24 11:50
2551,United States,"Reston, VA",12011 Sunset Hills Rd,One Reston Overlook,22091,COB,COGENT,-77.3574,38.9545,2020/5/24 11:50
2552,United States,"Reston, VA",12021 Sunset Hills Rd,Two Reston Overlook,22091,COB,COGENT,-77.3595,38.9541,2020/5/24 11:50
2553,United States,"Reston, VA",12098 Sunrise Valley Dr,CoreSite VA2,20191,CNDC,COGENT,-77.3639,38.9507,2020/5/24 11:50
2554,United States,"Reston, VA","12100 Sunrise Valley Dr, Stes CR1-CR10",CoreSite VA1,20191,CNDC,COGENT,-77.3641,38.9515,2020/5/24 11:50
2555,United States,"Reston, VA",12369 Sunrise Valley Drive,Coresite VA3,20191,CNDC,COGENT,-77.3717,38.9482,2020/5/24 11:50
2556,United States,"Reston, VA",1780 Business Center Dr,Digital Realty (fka. DuPont Fabros VA3),20190,CNDC,COGENT,-77.3272,38.9485,2020/5/24 11:50
2557,United States,"Reston, VA",1807 Michael Faraday Ct,Evocative DC1,20190,CNDC,COGENT,-77.3293,38.9488,2020/5/24 11:50
2558,United States,"Reston, VA",1818 Library St.,1818 Library St.,20190,COB,COGENT,-77.3589,38.9583,2020/5/24 11:50
2560,United States,"Reston, VA",1875 Explorer St,1875 Explorer St,20190,COB,COGENT,-77.3598,38.9583,2020/5/24 11:50
2561,United States,"Reston Town Center, VA",1750 Presidents Street,1750 Presidents Street,20190,COB,COGENT,-77.3495,38.9585,2020/5/24 11:50
2562,United States,"Richardson, TX",1001 East Campbell Road,SunGard dal-1001,75081,CNDC,COGENT,-96.7074,32.9761,2020/5/24 11:50
2563,United States,"Richardson, TX",1210 Integrity Way,Digital Realty,75081,CNDC,COGENT,-96.712,32.9655,2020/5/24 11:50
2564,United States,"Richardson, TX",1215 Integrity Drive,Digital Realty,75081,CNDC,COGENT,-96.7157,32.9649,2020/5/24 11:50
2565,United States,"Richardson, TX",1232 Alma Rd,Digital Realty,75081,CNDC,COGENT,-96.7155,32.9662,2020/5/24 11:50
2566,United States,"Richardson, TX",1232 Alma Rd,Equinix DA10 (fka. Verizon),75081,CNDC,COGENT,-96.7155,32.9662,2020/5/24 11:50
2567,United States,"Richardson, TX",3010 Waterview Parkway,Flexential (fka. Peak 10 / ViaWest),75080,CNDC,COGENT,-96.7525,32.9948,2020/5/24 11:50
2568,United States,"Richardson, TX",850 E Collins Boulevard,Digital Realty,75081,CNDC,COGENT,-96.7137,32.9673,2020/5/24 11:50
2569,United States,"Richardson, TX",900 Quality Way,Digital Realty (IBM),75081,CNDC,COGENT,-96.7134,32.9672,2020/5/24 11:50
2571,United States,"Richardson, TX",908 Quality Way,Digital Realty,75081,CNDC,COGENT,-96.71,32.9663,2020/5/24 11:50
2572,United States,"Richardson, TX",950 E Collins Boulevard,Digital Realty (LinkedIn),75081,CNDC,COGENT,-96.7106,32.9672,2020/5/24 11:50
2573,United States,"Richmond, VA",1111 East Main Street,The Bank of America Center,23219,COB,COGENT,-77.435,37.5366,2020/5/24 11:50
2574,United States,"Rockville, MD",51 Monroe St,One Metro Square,20850,COB,COGENT,-77.1479,39.0839,2020/5/24 11:50
2575,United States,"Rockville, MD",600 Jefferson Plz,Jefferson Plaza 1,20852,COB,COGENT,-77.1463,39.0811,2020/5/24 11:50
2576,United States,"Royal Oak, MI",4815 Delemere Ave,Liberty Center One,48073,CNDC,COGENT,-83.1774,42.5335,2020/5/24 11:50
2577,United States,"Sacramento, CA",1100 N Market Blvd,QTS Datacenter (fka. Herakles),95834,CNDC,COGENT,-121.486,38.6464,2020/5/24 11:50
2578,United States,"Sacramento, CA",1200 Striker Avenue,RagingWire CA1 (fka. UFO),95834,CNDC,COGENT,-121.489,38.6504,2020/5/24 11:50
2579,United States,"Sacramento, CA",1312 Striker Avenue,RagingWire CA2,95834,CNDC,COGENT,-121.492,38.6503,2020/5/24 11:50
2580,United States,"Sacramento, CA",1625 National Dr.,RagingWire CA3,95834,CNDC,COGENT,-121.499,38.6443,2020/5/24 11:50
2581,United States,"Sacramento, CA",770 L St,770 L St,95814,CNDC,COGENT,-121.498,38.5786,2020/5/24 11:50
2582,United States,"Sacramento, CA",770 L St,770 L St,95814,COB,COGENT,-121.498,38.5786,2020/5/24 11:50
2583,United States,"Saint Paul, MN",101 5th Street East,US Bank Center,55101,COB,COGENT,-100.47,48.968,2020/5/24 11:50
2584,United States,"Saint Paul, MN",30 7th Street E,Wells Fargo Place,55101,COB,COGENT,-93.0956,44.9479,2020/5/24 11:50
2585,United States,"Saint Paul, MN",332 Minnesota Street,First National Bank Building,55101,COB,COGENT,-93.0912,44.9464,2020/5/24 11:50
2586,United States,"Saint Paul, MN",444 Cedar Street,UBS Plaza (aka Piper Jaffray Plaza),55101,COB,COGENT,-93.0941,44.9481,2020/5/24 11:50
2587,United States,"Saint Paul, MN",55 East 5th Street,Alliance Bank,55101,COB,COGENT,-73.9943,40.7346,2020/5/24 11:50
2588,United States,"Salt Lake City, UT",111 Main St,111 Main,84111,COB,COGENT,-111.891,40.7669,2020/5/24 11:50
2589,United States,"Salt Lake City, UT",118 South 1000 West,Flexential SLC07,84104,CNDC,COGENT,-111.92,40.7668,2020/5/24 11:50
2590,United States,"Salt Lake City, UT",136 East South Temple,136 E South Temple,84111,COB,COGENT,-111.887,40.7689,2020/5/24 11:50
2591,United States,"Salt Lake City, UT",136 South Main,Kearns Building,84101,COB,COGENT,-71.0975,41.7809,2020/5/24 11:50
2592,United States,"Salt Lake City, UT",170 South Main Street,US Bank Building,84101,COB,COGENT,-112.411,38.8828,2020/5/24 11:50
2593,United States,"Salt Lake City, UT",179 Social Hall Avenue,DataBank SLC1 (fka. C7 Data Centers),84111,CNDC,COGENT,-111.886,40.7685,2020/5/24 11:50
2594,United States,"Salt Lake City, UT",257 East 200 South Street,257 Tower,84111,COB,COGENT,-111.884,40.7655,2020/5/24 11:50
2595,United States,"Salt Lake City, UT",572 S Delong St,CenturyLink (fka. Level 3),84104,CNDC,COGENT,-111.954,40.7563,2020/5/24 11:50
2596,United States,"Salt Lake City, UT",572 S Delong St,Peak10 - Via West Datacenter,84104,CNDC,COGENT,-111.954,40.7563,2020/5/24 11:50
2597,United States,"San Antonio, TX",100 Taylor Street (AKA 301 & 323 Broadway),Taylor Telecom Ctr Bldg 1,78205,CNDC,COGENT,-86.786,36.1807,2020/5/24 11:50
2598,United States,"San Antonio, TX",112 East Pecan St,Weston Centre,78205,COB,COGENT,-98.4925,29.4282,2020/5/24 11:50
2599,United States,"San Antonio, TX",300 Convent Street,Bank of America Building,78205,COB,COGENT,-98.4909,29.4295,2020/5/24 11:50
2600,United States,"San Antonio, TX",310 S Saint Marys,Tower Life Building,78205,COB,COGENT,-98.4915,29.4229,2020/5/24 11:50
2601,United States,"San Antonio, TX",323 Broadway,Taylor Telecom Ctr Bldg 2 - Grande,78205,CNDC,COGENT,-71.1014,42.3697,2020/5/24 11:50
2602,United States,"San Diego, CA",101 W Broadway,101 W Broadway,92101,COB,COGENT,-117.164,32.7152,2020/5/24 11:50
2603,United States,"San Diego, CA",12270 World Trade Drive,zColo (fka. Castle Access),92128,CNDC,COGENT,-117.073,32.9888,2020/5/24 11:50
2604,United States,"San Diego, CA",225 Broadway,225 Broadway,92101,COB,COGENT,-73.8591,41.0838,2020/5/24 11:50
2605,United States,"San Diego, CA",401 B St,Wells Fargo Plaza,92101,COB,COGENT,-117.161,32.7173,2020/5/24 11:50
2606,United States,"San Diego, CA",401 W A St,Columbia Center,92101,COB,COGENT,-117.167,32.7185,2020/5/24 11:50
2607,United States,"San Diego, CA",402 W Broadway,Emerald Plaza,92101,COB,COGENT,-117.167,32.7162,2020/5/24 11:50
2608,United States,"San Diego, CA",4320 La Jolla Village Drive,The Plaza at La Jolla Village Building,921212123,COB,COGENT,-117.213,32.8732,2020/5/24 11:50
2609,United States,"San Diego, CA",4330 La Jolla Village Drive,Genesee Building,921212123,COB,COGENT,-117.213,32.8748,2020/5/24 11:50
2610,United States,"San Diego, CA",4350 La Jolla Village Dr,Smith Barney Tower,92122,COB,COGENT,-117.213,32.8741,2020/5/24 11:50
2611,United States,"San Diego, CA",4365 Executive Dr,Pacifica Tower,92122,COB,COGENT,-117.212,32.8746,2020/5/24 11:50
2612,United States,"San Diego, CA",4370 La Jolla Village Dr,Northern Trust Tower,92122,COB,COGENT,-117.211,32.874,2020/5/24 11:50
2613,United States,"San Diego, CA",4380 La Jolla Village Dr,Bank of America Building,92122,COB,COGENT,-117.21,32.8742,2020/5/24 11:50
2614,United States,"San Diego, CA",501 W Broadway,501 W Broadway,92101,COB,COGENT,-94.3762,45.9763,2020/5/24 11:50
2615,United States,"San Diego, CA",525 B St,525 B Street,92101,COB,COGENT,-117.16,32.7177,2020/5/24 11:50
2616,United States,"San Diego, CA","525 B St, Suite 1020",Cogent Data Center - San Diego,92101,CDC,COGENT,-117.16,32.7176,2020/5/24 11:50
2617,United States,"San Diego, CA",550 W C St,550 Corporate Center,92101,COB,COGENT,-111.664,40.1099,2020/5/24 11:50
2619,United States,"San Diego, CA",5775 Kearny Villa Road,Scale Matrix,92123,CNDC,COGENT,-117.133,32.8382,2020/5/24 11:50
2620,United States,"San Diego, CA",600 W Broadway,One America Plaza,92101,COB,COGENT,-117.169,32.7165,2020/5/24 11:50
2621,United States,"San Diego, CA",655 W Broadway,First Allied Plaza,92101,COB,COGENT,-73.9956,40.7269,2020/5/24 11:50
2622,United States,"San Diego, CA",701 B St,701 B Street,92101,COB,COGENT,-117.158,32.7174,2020/5/24 11:50
2623,United States,"San Diego, CA",750 B St,Symphony Towers,92101,COB,COGENT,-117.158,32.7183,2020/5/24 11:50
2626,United States,"San Diego, CA",8917 Complex Drive,Fiber Alley,92123,CNDC,COGENT,-117.134,32.8298,2020/5/24 11:50
2627,United States,"San Diego, CA",8929 Aero Dr - 1st Floor,CenturyLink (fka. Level 3),92123,CNDC,COGENT,-117.136,32.8089,2020/5/24 11:50
2628,United States,"San Diego, CA",8939 Complex Drive,AIS Fiber Alley (FADC 2),92123,CNDC,COGENT,-117.134,32.8299,2020/5/24 11:50
2629,United States,"San Diego, CA",8959 Complex Drive,8959 Complex Drive,92123,COB,COGENT,-117.134,32.8298,2020/5/24 11:50
2630,United States,"San Diego, CA",8971 Complex Dr,i2B networks,92123,CNDC,COGENT,-117.134,32.8298,2020/5/24 11:50
2631,United States,"San Diego, CA",9305 Lightwave Ave,AIS Lightwave (LWDC),92123,CNDC,COGENT,-117.13,32.8278,2020/5/24 11:50
2632,United States,"San Diego, CA",9530 Towne Centre Drive,Cogent Data Center - San Diego,92121,CDC,COGENT,-117.21,32.8806,2020/5/24 11:50
2633,United States,"San Diego, CA",9530 Towne Centre Drive,9530 Towne Centre Drive,92121,COB,COGENT,-117.21,32.8806,2020/5/24 11:50
2635,United States,"San Francisco, CA",1 California St,One California,94111,COB,COGENT,-122.397,37.7933,2020/5/24 11:50
2636,United States,"San Francisco, CA",1 Embarcadero Ctr,One Embarcadero Center,94111,COB,COGENT,-122.399,37.795,2020/5/24 11:50
2637,United States,"San Francisco, CA",1 Front St,One Front Street,94111,COB,COGENT,-122.399,37.7918,2020/5/24 11:50
2638,United States,"San Francisco, CA",1 Market Plaza Spear Tower,1 Market Plaza Spear Tower,94105,COB,COGENT,-122.394,37.7931,2020/5/24 11:50
2639,United States,"San Francisco, CA",1 Market Street,1 Market Plaza Steuart Tower,94105,COB,COGENT,-77.3228,37.5334,2020/5/24 11:50
2640,United States,"San Francisco, CA",1 Montgomery Street (120 Kearny St),Post Montgomery,94104,COB,COGENT,-122.404,37.7893,2020/5/24 11:50
2641,United States,"San Francisco, CA",1 Sansome St,Citicorp Center,94104,COB,COGENT,-122.401,37.7904,2020/5/24 11:50
2642,United States,"San Francisco, CA",100 Bush St,Shell Bldg,94104,COB,COGENT,-122.4,37.7914,2020/5/24 11:50
2643,United States,"San Francisco, CA",100 California St,100 California St,94111,COB,COGENT,-122.398,37.7937,2020/5/24 11:50
2644,United States,"San Francisco, CA",100 Montgomery St (fka 120 Montgomery St),Equitable Life Building,94104,COB,COGENT,-86.3095,32.3776,2020/5/24 11:50
2645,United States,"San Francisco, CA",101 California St,101 California St,94111,COB,COGENT,-122.399,37.7932,2020/5/24 11:50
2646,United States,"San Francisco, CA",101 Second Street,101 Second Street,94105,COB,COGENT,-122.399,37.7881,2020/5/24 11:50
2647,United States,"San Francisco, CA",185 Berry St (Berry Street Bldg),China Basin Landing - Berry Street Building,94107,COB,COGENT,-122.392,37.7766,2020/5/24 11:50
2649,United States,"San Francisco, CA",185 Berry St Suite 1820,CenturyLink (fka. Level 3),94107,CNDC,COGENT,-122.392,37.7766,2020/5/24 11:50
2650,United States,"San Francisco, CA",199 Fremont St,199 Fremont St,94105,COB,COGENT,-122.395,37.7899,2020/5/24 11:50
2651,United States,"San Francisco, CA",2 Embarcadero Ctr,Two Embarcadero Center,94111,COB,COGENT,-122.398,37.7948,2020/5/24 11:50
2652,United States,"San Francisco, CA",2 Harrison St,Hills Plaza II,94105,COB,COGENT,-122.389,37.7895,2020/5/24 11:50
2653,United States,"San Francisco, CA",200 Paul Ave - 3rd Floor,Digital Realty (fka. TelX),94124,CNDC,COGENT,-122.398,37.7237,2020/5/24 11:50
2654,United States,"San Francisco, CA",201 Mission St,201 Mission St,94105,COB,COGENT,-122.395,37.791,2020/5/24 11:50
2655,United States,"San Francisco, CA",220 Bush St,220 Bush St,94104,COB,COGENT,-122.401,37.7914,2020/5/24 11:50
2656,United States,"San Francisco, CA",220 Montgomery St,Mills Bldg,94104,COB,COGENT,-122.402,37.7913,2020/5/24 11:50
2657,United States,"San Francisco, CA",222 Sansome St,Mandarin Oriental Hotel,94104,COB,COGENT,-122.401,37.7925,2020/5/24 11:50
2658,United States,"San Francisco, CA",225 Bush St,225 Bush St,94104,COB,COGENT,-122.401,37.7909,2020/5/24 11:50
2659,United States,"San Francisco, CA",235 Montgomery St,Russ Bldg Tower,94104,COB,COGENT,-122.403,37.7914,2020/5/24 11:50
2660,United States,"San Francisco, CA",275 Battery St,Embarcadero Center West,94111,COB,COGENT,-122.4,37.7939,2020/5/24 11:50
2661,United States,"San Francisco, CA",3 Embarcadero Ctr,Three Embarcadero Center,94111,COB,COGENT,-122.397,37.7948,2020/5/24 11:50
2662,United States,"San Francisco, CA",300 Clay Street,One Maritime Plaza,94111,COB,COGENT,-122.399,37.7955,2020/5/24 11:50
2663,United States,"San Francisco, CA",301 Howard St,301 Howard St,94105,COB,COGENT,-122.394,37.7897,2020/5/24 11:50
2664,United States,"San Francisco, CA","303 2nd St, North Tower",Marathon Plaza,94107,COB,COGENT,-122.396,37.7853,2020/5/24 11:50
2665,United States,"San Francisco, CA","303 2nd St, South Tower",Marathon Plaza,94107,COB,COGENT,-122.396,37.7853,2020/5/24 11:50
2666,United States,"San Francisco, CA",333 Bush St,333 Bush St,94104,COB,COGENT,-122.403,37.7906,2020/5/24 11:50
2667,United States,"San Francisco, CA",345 California St,The California Center,94104,COB,COGENT,-122.4,37.7923,2020/5/24 11:50
2668,United States,"San Francisco, CA",345 Spears St,Hills Plaza I,94105,COB,COGENT,-122.39,37.79,2020/5/24 11:50
2669,United States,"San Francisco, CA",365 Main Street,Digital Realty,94105,CNDC,COGENT,-163.539,63.031,2020/5/24 11:50
2670,United States,"San Francisco, CA",388 Market St,388 Market St,94111,COB,COGENT,-122.398,37.7922,2020/5/24 11:50
2671,United States,"San Francisco, CA",4 Embarcadero Ctr,Four Embarcadero Center,94111,COB,COGENT,-122.396,37.7953,2020/5/24 11:50
2672,United States,"San Francisco, CA",415 Mission St,Salesforce Tower,94105,COB,COGENT,-122.397,37.7898,2020/5/24 11:50
2673,United States,"San Francisco, CA",425 Market St,One Metropolitan Plaza,94105,COB,COGENT,-122.398,37.7912,2020/5/24 11:50
2674,United States,"San Francisco, CA",44 Montgomery St,44 Montgomery St,94104,COB,COGENT,-122.402,37.7897,2020/5/24 11:50
2675,United States,"San Francisco, CA",475 Sansome St,475 Sansome St,94111,COB,COGENT,-122.402,37.7947,2020/5/24 11:50
2676,United States,"San Francisco, CA",50 Beale St,Bechtel Building,94105,COB,COGENT,-122.397,37.7912,2020/5/24 11:50
2677,United States,"San Francisco, CA",50 California St,50 California St,94111,COB,COGENT,-122.397,37.794,2020/5/24 11:50
2678,United States,"San Francisco, CA",50 Fremont St,50 Fremont Center,94105,COB,COGENT,-122.397,37.7906,2020/5/24 11:50
2679,United States,"San Francisco, CA",500 Sansome St,500 Sansome St,94111,COB,COGENT,-122.401,37.7952,2020/5/24 11:50
2680,United States,"San Francisco, CA",501 2nd St,501 2nd St,94107,COB,COGENT,-122.393,37.7833,2020/5/24 11:50
2681,United States,"San Francisco, CA",505 Montgomery St,505 Montgomery St,94111,COB,COGENT,-122.403,37.7939,2020/5/24 11:50
2682,United States,"San Francisco, CA",505 Sansome St,505 Sansome St,94111,COB,COGENT,-122.402,37.7952,2020/5/24 11:50
2683,United States,"San Francisco, CA",525 Market St,First Market Tower,94105,COB,COGENT,-122.399,37.7905,2020/5/24 11:50
2684,United States,"San Francisco, CA",555 California St,Bank of America Headquarters,94104,COB,COGENT,-122.404,37.792,2020/5/24 11:50
2685,United States,"San Francisco, CA",555 Market St,555 Market St,94105,COB,COGENT,-122.4,37.79,2020/5/24 11:50
2686,United States,"San Francisco, CA",560 Mission,560 Mission Street,94105,COB,COGENT,-122.4,37.7889,2020/5/24 11:50
2687,United States,"San Francisco, CA",575 Market St,575 Market St,94105,COB,COGENT,-122.4,37.7896,2020/5/24 11:50
2688,United States,"San Francisco, CA",580 California St,580 California St,94104,COB,COGENT,-122.404,37.793,2020/5/24 11:50
2689,United States,"San Francisco, CA",595 Market St,595 Market St,94105,COB,COGENT,-122.401,37.7892,2020/5/24 11:50
2690,United States,"San Francisco, CA",600 Montgomery St,Transamerica Pyramid,94111,COB,COGENT,-89.7309,33.4733,2020/5/24 11:50
2691,United States,"San Francisco, CA",601 California St,International Bldg,94108,COB,COGENT,-122.405,37.7924,2020/5/24 11:50
2693,United States,"San Francisco, CA",77 Battery Street,77 Battery Street,94111,COB,COGENT,-122.4,37.792,2020/5/24 11:50
2694,United States,"San Francisco, CA",88 Kearny St,California Federal Savings,94108,COB,COGENT,-122.403,37.7886,2020/5/24 11:50
2695,United States,"San Jose, CA",1 Almaden Blvd,1 Almaden Blvd,95113,COB,COGENT,-121.895,37.3334,2020/5/24 11:50
2696,United States,"San Jose, CA",11 Great Oaks Blvd,Equinix SV1,95119,CNDC,COGENT,-121.783,37.2415,2020/5/24 11:50
2697,United States,"San Jose, CA",150 South 1st St,Rackspace (fka. Datapipe),95113,CNDC,COGENT,-73.9616,40.7138,2020/5/24 11:50
2698,United States,"San Jose, CA",1735 Lundy Ave,Equinix SV3,95131,CNDC,COGENT,-121.888,37.3882,2020/5/24 11:50
2700,United States,"San Jose, CA",50 W San Fernando St,Fairmont Plaza Ph I,95113,COB,COGENT,-121.889,37.3338,2020/5/24 11:50
2701,United States,"San Jose, CA",534 Stockton Ave 1st Floor,Evocative (fka. 365 Data Centers & Equinix SV7 - S,95126,CNDC,COGENT,-121.91,37.3385,2020/5/24 11:50
2702,United States,"San Jose, CA",55 Almaden,55 Almaden,95110,COB,COGENT,-121.871,37.234,2020/5/24 11:50
2703,United States,"San Jose, CA",55 S Market St,CoreSite SV1,95113,CNDC,COGENT,-82.5499,35.5932,2020/5/24 11:50
2704,United States,"San Jose, CA",55 S Market St,Tower 55,95113,COB,COGENT,-82.5499,35.5932,2020/5/24 11:50
2705,United States,"San Jose, CA",7 Great Oaks Blvd.,Equinix SV10,95119,CNDC,COGENT,-121.785,37.2412,2020/5/24 11:50
2706,United States,"San Jose, CA",9 Great Oaks Boulevard,Equinix SV5,95119,CNDC,COGENT,-121.782,37.2419,2020/5/24 11:50
2707,United States,"San Jose, CA",99 Almaden Blvd,Union Bank Bldg,95113,COB,COGENT,-121.894,37.3325,2020/5/24 11:50
2708,United States,"Sandston, VA",6000 Technology Blvd,QTS Datacenter,23150,CNDC,COGENT,-77.2442,37.4881,2020/5/24 11:50
2709,United States,"Santa Clara, CA",1100 Space Park Drive,Digital Realty SLC1 (fka. Telx MMR),95050,CNDC,COGENT,-121.953,37.3757,2020/5/24 11:50
2710,United States,"Santa Clara, CA",1101 Space Park,Colovore,95054,CNDC,COGENT,-121.953,37.3767,2020/5/24 11:50
2711,United States,"Santa Clara, CA",1350 Duane Ave,Equinix SV2,95054,CNDC,COGENT,-121.955,37.3782,2020/5/24 11:50
2712,United States,"Santa Clara, CA",1525 Comstock St,Internap SJE014,95054,CNDC,COGENT,-121.956,37.3749,2020/5/24 11:50
2713,United States,"Santa Clara, CA",1700 Richard Avenue,EdgeConneX EDCSCV01,95050,CNDC,COGENT,-121.959,37.3642,2020/5/24 11:50
2714,United States,"Santa Clara, CA",2050 Martin Avenue,vXchnge,95050,CNDC,COGENT,-121.962,37.3654,2020/5/24 11:50
2715,United States,"Santa Clara, CA",2151 Mission College Boulevard,Internap,95054,CNDC,COGENT,-121.962,37.3895,2020/5/24 11:50
2716,United States,"Santa Clara, CA",2401 Walsh,Cyxtera SC4,95054,CNDC,COGENT,-121.97,37.3716,2020/5/24 11:50
2717,United States,"Santa Clara, CA",2403 Walsh,Cyxtera SC5,95054,CNDC,COGENT,-121.97,37.3704,2020/5/24 11:50
2719,United States,"Santa Clara, CA",2805 LaFayette St.,Digital Realty,95050,CNDC,COGENT,-121.949,37.3723,2020/5/24 11:50
2720,United States,"Santa Clara, CA",2807 Mission College Boulevard,QTS Datacenter,95054,CNDC,COGENT,-121.978,37.3939,2020/5/24 11:50
2721,United States,"Santa Clara, CA",2820 Northwestern Parkway,Vantage (V1 & V4) (Telx MMR),95051,CNDC,COGENT,-121.974,37.3724,2020/5/24 11:50
2722,United States,"Santa Clara, CA",2880 Northwestern Parkway,Vantage Data Center (V3),95051,CNDC,COGENT,-121.973,37.3736,2020/5/24 11:50
2723,United States,"Santa Clara, CA",2900 Stender Way,CoreSite SV5,95054,CNDC,COGENT,-121.946,37.3549,2020/5/24 11:50
2724,United States,"Santa Clara, CA",2901 Coronado Drive,CoreSite SV3,95054,CNDC,COGENT,-121.972,37.3751,2020/5/24 11:50
2725,United States,"Santa Clara, CA",2950 Stender Way,CoreSite 2950 Stender Way,95054,CNDC,COGENT,-121.971,37.3755,2020/5/24 11:50
2727,United States,"Santa Clara, CA",2970 Corvin Drive,Equinix SV15 (fka. Terremark / Verizon NAP West),95051,CNDC,COGENT,-121.989,37.3755,2020/5/24 11:50
2728,United States,"Santa Clara, CA",2972 Stender Way,CoreSite SV4,95054,CNDC,COGENT,-121.971,37.3762,2020/5/24 11:50
2729,United States,"Santa Clara, CA",3000 Corvin Drive,Equinix SV16 (fka. Terremark / Verizon NAP West),95051,CNDC,COGENT,-121.989,37.3759,2020/5/24 11:50
2731,United States,"Santa Clara, CA",3020 Coronado Dr,CoreSite SV7,95054,CNDC,COGENT,-117.168,32.6948,2020/5/24 11:50
2732,United States,"Santa Clara, CA",3030 Corvin Drive,Equinix SV17 (fka. Terremark NAP),95051,CNDC,COGENT,-121.989,37.3765,2020/5/24 11:50
2733,United States,"Santa Clara, CA",3080 Raymond St,Wave Data Center (fka. Layer 42),95054,CNDC,COGENT,-121.954,37.378,2020/5/24 11:50
2734,United States,"Santa Clara, CA",3105 Alfred St,Digital Realty,95054,CNDC,COGENT,-121.959,37.3781,2020/5/24 11:50
2736,United States,"Santa Clara, CA",3340 Bassett Street,3340 Bassett Street,95054,CNDC,COGENT,-121.95,37.3812,2020/5/24 11:50
2737,United States,"Santa Clara, CA",4650 Old Ironsides Dr,Cyxtera SC9,95054,CNDC,COGENT,-121.981,37.3976,2020/5/24 11:50
2738,United States,"Santa Clara, CA",4700 Old Ironsides Dr,Cyxtera SC8,95054,CNDC,COGENT,-121.981,37.3986,2020/5/24 11:50
2739,United States,"Santa Monica, CA",100 Wilshire Blvd,The 100 Wilshire Bldg,90401,COB,COGENT,-118.501,34.0168,2020/5/24 11:50
2740,United States,"Santa Monica, CA",1601 Cloverfield Blvd,Water Garden Phase II,90404,COB,COGENT,-118.473,34.0282,2020/5/24 11:50
2741,United States,"Santa Monica, CA",1620 26th St,Water Garden Phase I,90404,COB,COGENT,-118.471,34.0294,2020/5/24 11:50
2742,United States,"Santa Monica, CA",2400 Broadway,Colorado Center Bldg D,90404,COB,COGENT,-89.6664,36.1708,2020/5/24 11:50
2743,United States,"Santa Monica, CA",2401 Colorado Ave,Colorado Center Bldg A,90404,COB,COGENT,-118.473,34.0292,2020/5/24 11:50
2744,United States,"Santa Monica, CA",2425 Olympic Blvd,Water Garden Phase I,90404,COB,COGENT,-118.471,34.0281,2020/5/24 11:50
2745,United States,"Santa Monica, CA",2425-2501 Colorado Ave,Colorado Center Bldg B,90404,COB,COGENT,-118.482,34.0217,2020/5/24 11:50
2746,United States,"Santa Monica, CA",2500 Broadway,Colorado Center Bldg F,90404,COB,COGENT,-118.474,34.0307,2020/5/24 11:50
2747,United States,"Scottsdale, AZ",8521 E Princess Drive,Iron Mountain (fka. IO Data Center),85255,CNDC,COGENT,-111.896,33.6445,2020/5/24 11:50
2748,United States,"Seattle, WA",100 Fourth Ave N,Fisher Plaza West,98109,COB,COGENT,-92.7771,45.024,2020/5/24 11:50
2749,United States,"Seattle, WA",1000 2nd Ave,1000 Second Ave,98104,COB,COGENT,-122.335,47.6058,2020/5/24 11:50
2750,United States,"Seattle, WA",1000 Denny Way - 4th Floor,CenturyLink (fka. Level 3),98109,CNDC,COGENT,-122.336,47.619,2020/5/24 11:50
2751,United States,"Seattle, WA",1000 Denny Way-1st Floor,H5 Seattle,98109,CNDC,COGENT,-122.336,47.619,2020/5/24 11:50
2752,United States,"Seattle, WA",1001 4th Ave,1001 Fourth Avenue Plaza,98154,COB,COGENT,-122.334,47.6061,2020/5/24 11:50
2753,United States,"Seattle, WA",1111 3rd Ave,1111 Third Ave Bldg,98101,COB,COGENT,-122.335,47.6066,2020/5/24 11:50
2754,United States,"Seattle, WA",1191 2nd Ave,Safeco Center,98101,COB,COGENT,-122.337,47.606,2020/5/24 11:50
2755,United States,"Seattle, WA",1200 5th Ave,IBM Bldg,98101,COB,COGENT,-73.9531,40.7913,2020/5/24 11:50
2756,United States,"Seattle, WA",1201 3rd Ave,Washington Mutual Tower,98101,COB,COGENT,-122.336,47.6071,2020/5/24 11:50
2757,United States,"Seattle, WA",1220 Howell St,Metropolitan Park/North Tower,98101,COB,COGENT,-122.33,47.6175,2020/5/24 11:50
2758,United States,"Seattle, WA",12201 Tukwila International Blvd,Wowrack,98168,CNDC,COGENT,-122.294,47.4931,2020/5/24 11:50
2759,United States,"Seattle, WA",1301 2nd Avenue,Russell Investment Center,98101,COB,COGENT,-122.338,47.6075,2020/5/24 11:50
2760,United States,"Seattle, WA",1301 5th Ave,Rainier Tower,98101,COB,COGENT,-122.334,47.6089,2020/5/24 11:50
2761,United States,"Seattle, WA",1326 5th Ave,The Skinner Building,98101,COB,COGENT,-122.334,47.6094,2020/5/24 11:50
2762,United States,"Seattle, WA",140 Fourth Ave N,Tierpoint Seattle DC1,98109,CNDC,COGENT,-122.348,47.6194,2020/5/24 11:50
2763,United States,"Seattle, WA",140 Fourth Ave N,Fisher Plaza East,98109,COB,COGENT,-122.348,47.6194,2020/5/24 11:50
2764,United States,"Seattle, WA",140 Fourth Ave. N,KOMO Plaza (fka. Fisher Plz) MMR & Datacenter,98109,CNDC,COGENT,-122.348,47.6194,2020/5/24 11:50
2765,United States,"Seattle, WA",1420 5th Ave,US Bank Centre,98101,COB,COGENT,-122.334,47.6103,2020/5/24 11:50
2766,United States,"Seattle, WA",1500 4th Ave,1500 4th Ave,98101,COB,COGENT,-122.337,47.6106,2020/5/24 11:50
2767,United States,"Seattle, WA",1501 4th Ave,Century Square,98101,COB,COGENT,-122.338,47.6103,2020/5/24 11:50
2768,United States,"Seattle, WA",1501 First Avenue South,Home Plate Center,98134,COB,COGENT,-122.335,47.5898,2020/5/24 11:50
2769,United States,"Seattle, WA",1505 5th Ave,Renke Building,98101,COB,COGENT,-122.336,47.6109,2020/5/24 11:50
2770,United States,"Seattle, WA",1521 First Avenue South,Home Plate Center,98134,COB,COGENT,-122.335,47.5886,2020/5/24 11:50
2772,United States,"Seattle, WA",1730 Minor Ave,Metropolitan Park/East Tower,98101,COB,COGENT,-122.33,47.6169,2020/5/24 11:50
2773,United States,"Seattle, WA",1904 Third Ave,Securities Building,98101,COB,COGENT,-73.9454,40.7911,2020/5/24 11:50
2775,United States,"Seattle, WA",2001 6th Ave,Equinix SE2 (12th floor) (fka. S&D),98121,CNDC,COGENT,-87.5632,33.2102,2020/5/24 11:50
2776,United States,"Seattle, WA",2001 6th Ave,Equinix SE2 (13th floor) (fka. S&D),98121,CNDC,COGENT,-87.5632,33.2102,2020/5/24 11:50
2777,United States,"Seattle, WA",2001 6th Ave,ColoCenters - 18th Floor,98121,CNDC,COGENT,-87.5632,33.2102,2020/5/24 11:50
2778,United States,"Seattle, WA",2001 6th Ave,Swift Ventures - Suite 2008A,98121,CNDC,COGENT,-87.5632,33.2102,2020/5/24 11:50
2779,United States,"Seattle, WA",2001 6th Ave,Westin Building MMR - 19th Floor,98121,CNDC,COGENT,-87.5632,33.2102,2020/5/24 11:50
2780,United States,"Seattle, WA",2001 6th Ave,Westin Building Exchange,98121,COB,COGENT,-87.5632,33.2102,2020/5/24 11:50
2781,United States,"Seattle, WA",2020 5th Ave,Equinix SE3,98121,CNDC,COGENT,-86.8075,33.5188,2020/5/24 11:50
2782,United States,"Seattle, WA",2101 4th Ave,4th & Blanchard,98121,COB,COGENT,-122.343,47.6142,2020/5/24 11:50
2783,United States,"Seattle, WA",2200 6th Ave,Denny Bldg,98121,COB,COGENT,-77.4198,37.5608,2020/5/24 11:50
2784,United States,"Seattle, WA",3101 Western Ave,Digital Fortress,98121,CNDC,COGENT,-122.357,47.6178,2020/5/24 11:50
2785,United States,"Seattle, WA",403 Columbia Street,Columbia House,98104,,COGENT,-121.983,45.6377,2020/5/24 11:50
2786,United States,"Seattle, WA",600 University St,One Union Square,98101,COB,COGENT,-122.332,47.6097,2020/5/24 11:50
2787,United States,"Seattle, WA",601 Union St,Two Union Square,98101,COB,COGENT,-122.332,47.6102,2020/5/24 11:50
2788,United States,"Seattle, WA",701 5th Ave,Columbia Center,98104,COB,COGENT,-122.331,47.6046,2020/5/24 11:50
2789,United States,"Seattle, WA",800 5th Ave,Bank of America Bldg,98104,COB,COGENT,-122.33,47.6058,2020/5/24 11:50
2790,United States,"Seattle, WA",925 Fourth Avenue,Madison Financial IDX Tower,98104,COB,COGENT,-122.332,47.6062,2020/5/24 11:50
2791,United States,"Seattle, WA",999 3rd Ave,Wells Fargo Center,98104,COB,COGENT,-122.334,47.605,2020/5/24 11:50
2792,United States,"Secaucus, NJ",1 Enterprise Avenue North,Internap,7094,CNDC,COGENT,-74.0681,40.7875,2020/5/24 11:50
2793,United States,"Secaucus, NJ",105 Enterprise Avenue South,Equinix NY6,7904,CNDC,COGENT,-90.3355,43.6592,2020/5/24 11:50
2794,United States,"Secaucus, NJ",110 B Meadowlands Pkwy,Interserver TEB-2,7094,CNDC,COGENT,-74.0738,40.7861,2020/5/24 11:50
2795,United States,"Secaucus, NJ",15 Enterprise Ave North,Evoque NYC3,7094,CNDC,COGENT,-74.069,40.7838,2020/5/24 11:50
2796,United States,"Secaucus, NJ",2 Emerson Lane,Coresite NY2,7094,CNDC,COGENT,-74.0643,40.7852,2020/5/24 11:50
2797,United States,"Secaucus, NJ",200B Meadowlands Pkwy,vXchnge,7094,CNDC,COGENT,-74.075,40.7853,2020/5/24 11:50
2798,United States,"Secaucus, NJ",275 Hartz Way 1st Floor,Equinix NY2,7094,CNDC,COGENT,-74.0759,40.7778,2020/5/24 11:50
2799,United States,"Secaucus, NJ",755 Secaucus Rd,Equinix NY4,7094,CNDC,COGENT,-74.0696,40.7773,2020/5/24 11:50
2800,United States,"Secaucus, NJ",800 Secaucus Road,Equinix NY5,7094,CNDC,COGENT,-74.0722,40.7785,2020/5/24 11:50
2801,United States,"Shakopee, MN",4450 Dean Lakes Boulevard,Cyxtera MP2,55379,CNDC,COGENT,-93.462,44.7788,2020/5/24 11:50
2802,United States,"Silver Spring, MD",12401 Prosperity Dr,MDC-1,20904,CNDC,COGENT,-74.0009,41.9377,2020/5/24 11:50
2803,United States,"Silver Spring, MD",8401 Colesville Rd,Silver Spring Metro Plaza 1,20910,COB,COGENT,-77.0315,38.9949,2020/5/24 11:50
2804,United States,"Silver Spring, MD",8403 Colesville Rd,Silver Spring Metro Plaza 2,20910,COB,COGENT,-77.0315,38.9949,2020/5/24 11:50
2805,United States,"Silver Spring, MD",8405 Colesville Rd,Silver Spring Metro Plaza,20910,COB,COGENT,-77.0314,38.995,2020/5/24 11:50
2806,United States,"Somerset, NJ",125 Belmont Drive,Rackspace (fka. Datapipe),8873,CNDC,COGENT,-74.5369,40.5408,2020/5/24 11:50
2807,United States,"Somerset, NJ",800 Cottontail Lane,Sentinel NY-1,8873,CNDC,COGENT,-74.546,40.5382,2020/5/24 11:50
2808,United States,"Somerville, MA",169 Holland St,169 Holland St,2144,COB,COGENT,-71.1252,42.3998,2020/5/24 11:50
2809,United States,"Somerville, MA",50 Innerbelt,Internap - BSN 003,2413,CNDC,COGENT,-71.0804,42.3779,2020/5/24 11:50
2810,United States,"Somerville, MA",70 Innerbelt Road,CoreSite BO1 (fka. CRG West),2143,CNDC,COGENT,-71.0798,42.3766,2020/5/24 11:50
2811,United States,"South Bend, IN",1440 Ignition Drive South,Data Realty,46601,CNDC,COGENT,-86.2587,41.6599,2020/5/24 11:50
2812,United States,"South Bend, IN",506 W South Street,GAP - Union Station Technology Center,46601,CNDC,COGENT,-86.256,41.6693,2020/5/24 11:50
2813,United States,"South Bend, IN",746 South Arnold Street,ColoStore,46619,CNDC,COGENT,-86.2677,41.6676,2020/5/24 11:50
2814,United States,"Southfield, MI",1000 Town Center,1000 Town Center,48075,COB,COGENT,-83.2438,42.4754,2020/5/24 11:50
2815,United States,"Southfield, MI",19675 West 10 Mile Road,CenturyLink (fka. Level 3),48078,CNDC,COGENT,-83.2381,42.4722,2020/5/24 11:50
2816,United States,"Southfield, MI",2000 Town Center,2000 Town Center,48075,COB,COGENT,-83.2457,42.4774,2020/5/24 11:50
2817,United States,"Southfield, MI",21005 Lahser Road,EdgeConneX EDCDET01,48033,CNDC,COGENT,-83.2612,42.4485,2020/5/24 11:50
2818,United States,"Southfield, MI",21420 Melrose Ave,ServOnDemand,48075,CNDC,COGENT,-83.2539,42.4484,2020/5/24 11:50
2819,United States,"Southfield, MI",21700 Melrose Avenue,Nexcess - Southfield,48075,CNDC,COGENT,-83.2565,42.4463,2020/5/24 11:50
2820,United States,"Southfield, MI",24245 Northwestern Highway,123Net - DC3,48075,CNDC,COGENT,-83.2347,42.4669,2020/5/24 11:50
2821,United States,"Southfield, MI",24275 Northwestern Highway,123Net - DC2,48075,CNDC,COGENT,-83.235,42.467,2020/5/24 11:50
2822,United States,"Southfield, MI",24660 Lahser Rd,365 Data Centers (fka. Equinix - S&D),48033,CNDC,COGENT,-83.2598,42.4699,2020/5/24 11:50
2823,United States,"Southfield, MI",24700 Northwestern Highway,123Net - DC1,48075,CNDC,COGENT,-83.2381,42.4716,2020/5/24 11:50
2824,United States,"Southfield, MI",27777 Franklin Road,American Center,48076,COB,COGENT,-83.299,42.4899,2020/5/24 11:50
2825,United States,"Southfield, MI",3000 Town Center,3000 Town Center,48075,COB,COGENT,-83.2462,42.478,2020/5/24 11:50
2826,United States,"Southfield, MI",4000 Town Center,4000 Town Center,48075,COB,COGENT,-83.244,42.4793,2020/5/24 11:50
2827,United States,"Southfield, MI",One Town Square,One Town Square,48076,COB,COGENT,-82.5556,35.486,2020/5/24 11:50
2828,United States,"Southfield, MI",Two Towne Square,Two Towne Square,48076,COB,COGENT,-83.2538,42.4826,2020/5/24 11:50
2829,United States,"Spring, TX",2626 Spring Cypress Road,TRG Data Center,77388,CNDC,COGENT,-95.459,30.066,2020/5/24 11:50
2830,United States,"St. Louis, MO",1 Memorial Dr,Gateway Tower,63101,COB,COGENT,-95.5143,29.7513,2020/5/24 11:50
2831,United States,"St. Louis, MO",10 South Broadway,Equitable Building,63102,COB,COGENT,-90.1895,38.6249,2020/5/24 11:50
2832,United States,"St. Louis, MO",100 North Broadway,Bank of America Tower,63102,COB,COGENT,-99.0955,35.0286,2020/5/24 11:50
2833,United States,"St. Louis, MO",100 South Fourth St,The Deloitte Building,63102,COB,COGENT,-73.9638,40.712,2020/5/24 11:50
2834,United States,"St. Louis, MO",1010 Market Street,1010 Market,63102,COB,COGENT,-90.1967,38.6265,2020/5/24 11:50
2835,United States,"St. Louis, MO",200 North Broadway,St. Louis Place,63102,COB,COGENT,-94.9186,36.6176,2020/5/24 11:50
2836,United States,"St. Louis, MO",210 N Tucker Blvd,Netrality,63101,CNDC,COGENT,-90.1972,38.6293,2020/5/24 11:50
2837,United States,"St. Louis, MO",210 N Tucker Blvd,Netrality (fka. DRT),63101,CNDC,COGENT,-90.1972,38.6293,2020/5/24 11:50
2838,United States,"St. Louis, MO",211 N Broadway,One Metropolitan Square,63102,COB,COGENT,-89.1936,40.1449,2020/5/24 11:50
2839,United States,"St. Louis, MO",600 Washington Ave,One City Center,63101,COB,COGENT,-95.3661,29.7661,2020/5/24 11:50
2840,United States,"St. Louis, MO",710 N Tucker Blvd,vXchnge,63101,CNDC,COGENT,-90.1956,38.6327,2020/5/24 11:50
2841,United States,"St. Louis, MO",710 N Tucker Blvd,vXchnge,63101,CNDC,COGENT,-90.1956,38.6327,2020/5/24 11:50
2842,United States,"St. Louis, MO",720 Olive Street,Laclede Gas Building,63101,COB,COGENT,-90.1923,38.628,2020/5/24 11:50
2843,United States,"St. Louis, MO",800 Market,Bank of America Plaza,63101,COB,COGENT,-81.3723,40.8037,2020/5/24 11:50
2844,United States,"St. Louis, MO",900 Walnut St,Netrality Properties,63102,COB,COGENT,-83.2669,42.7682,2020/5/24 11:50
2845,United States,"St. Louis, MO",900 Walnut St Suite 360,Tierpoint SLW,63102,CNDC,COGENT,-75.1562,39.9481,2020/5/24 11:50
2846,United States,"St. Louis, MO",900 Walnut Street,Netrality (MMR),63102,CNDC,COGENT,-75.1562,39.9481,2020/5/24 11:50
2847,United States,"St. Paul, MN",380-388 St Peter Street,Lawson Commons,55102,COB,COGENT,-93.0958,44.9456,2020/5/24 11:50
2848,United States,"Stamford, CT",1 Landmark Square,Landmark Square,6901,COB,COGENT,-73.6595,41.0058,2020/5/24 11:50
2849,United States,"Stamford, CT",10 Riverbend Road South,CyrusOne Stamford,6907,CNDC,COGENT,-86.1747,36.9271,2020/5/24 11:50
2850,United States,"Stamford, CT",107 Elm Street,Four Stamford Plaza,6901,COB,COGENT,-73.532,41.0526,2020/5/24 11:50
2851,United States,"Stamford, CT",1266 E Main St,Soundview Plaza,6902,COB,COGENT,-73.5097,41.0603,2020/5/24 11:50
2852,United States,"Stamford, CT",1351 Washington Blvd,FirstLight (fka. Fibertech),6902,CNDC,COGENT,-73.5429,41.0581,2020/5/24 11:50
2853,United States,"Stamford, CT",177 Broad Street,177 Broad Street,6901,COB,COGENT,-75.9212,39.2403,2020/5/24 11:50
2854,United States,"Stamford, CT",2 Landmark Square,Landmark Square,6901,COB,COGENT,-73.538,41.0547,2020/5/24 11:50
2855,United States,"Stamford, CT",201 Broad Street,Canterbury Green,6901,COB,COGENT,-75.9217,39.2403,2020/5/24 11:50
2858,United States,"Stamford, CT",232 Harbor Plaza,Harbor Plaza,6902,COB,COGENT,-73.5335,41.0345,2020/5/24 11:50
2859,United States,"Stamford, CT",250 Harbor Drive,Harbor Plaza,6902,COB,COGENT,-80.1706,25.6967,2020/5/24 11:50
2860,United States,"Stamford, CT",262 Harbor Drive,Harbor Plaza,6902,COB,COGENT,-73.5339,41.0333,2020/5/24 11:50
2861,United States,"Stamford, CT",263 Tresser Boulevard,One Stamford Plaza,6901,COB,COGENT,-73.5342,41.0518,2020/5/24 11:50
2862,United States,"Stamford, CT",281 Tresser Boulvevard,Two Stamford Plaza,6901,COB,COGENT,-73.5341,41.052,2020/5/24 11:50
2863,United States,"Stamford, CT",290 Harbor Drive,Harbor Plaza,6902,COB,COGENT,-78.415,36.3619,2020/5/24 11:50
2864,United States,"Stamford, CT",3 Landmark Square,Landmark Square,6901,COB,COGENT,-73.538,41.0548,2020/5/24 11:50
2865,United States,"Stamford, CT",300 Atlantic Street,300 Atlantic Street,6901,COB,COGENT,-73.5393,41.0514,2020/5/24 11:50
2866,United States,"Stamford, CT",301 Tresser Boulevard,Three Stamford Plaza,6901,COB,COGENT,-73.5329,41.0523,2020/5/24 11:50
2867,United States,"Stamford, CT",4 Landmark Square,Landmark Square,6901,COB,COGENT,-73.538,41.0547,2020/5/24 11:50
2868,United States,"Stamford, CT",6 Landmark Square,Landmark Square,6901,COB,COGENT,-73.538,41.0547,2020/5/24 11:50
2869,United States,"Stamford, CT",680 Washington Blvd,Stamford Tower,6901,COB,COGENT,-73.5435,41.049,2020/5/24 11:50
2870,United States,"Stamford, CT",750 Washington Blvd,Stamford Tower,6901,COB,COGENT,-73.5437,41.0498,2020/5/24 11:50
2871,United States,"Stamford, CT",One Station Place,Metro Center,6902,COB,COGENT,-73.5424,41.0458,2020/5/24 11:50
2872,United States,"Staten Island, NY",7 Teleport Drive,Telehouse Teleport,10311,CNDC,COGENT,-74.1748,40.6069,2020/5/24 11:50
2873,United States,"Sterling, VA",1506 Moran Rd,QTS (fka. Carpathia),20166,CNDC,COGENT,-77.4479,38.9946,2020/5/24 11:50
2874,United States,"Sterling, VA",21110 Ridgetop Circle,Cyxtera IAD1 (fka. CenturyLink DC4),20166,CNDC,COGENT,-77.4132,39.0253,2020/5/24 11:50
2875,United States,"Sterling, VA",21111 Ridgetop Circle,"CyrusOne NVA - Sterling I,II,III",20166,CNDC,COGENT,-77.4141,39.0254,2020/5/24 11:50
2876,United States,"Sterling, VA",21350 Pacific Blvd,CyrusOne NVA - Sterling V,20166,CNDC,COGENT,-77.4339,39.0251,2020/5/24 11:50
2877,United States,"Sterling, VA",22810 International Dr,Cyxtera DC5,20166,CNDC,COGENT,-77.4239,38.9844,2020/5/24 11:50
2878,United States,"Sterling, VA",22860 International Dr,Cyxtera DC6,20166,CNDC,COGENT,-77.4247,38.9839,2020/5/24 11:50
2880,United States,"Sterling, VA",45845 Nokes Blvd,Cyxtera IAD1 (fka. CenturyLink DC3),20164,CNDC,COGENT,-77.4189,39.0263,2020/5/24 11:50
2881,United States,"Sterling, VA",45901 Nokes Blvd,Cyxtera IAD1 (fka. CenturyLink DC2),20166,CNDC,COGENT,-77.4134,39.0234,2020/5/24 11:50
2882,United States,"Sunnyvale, CA",1320 Kifer Road,Cyxtera SN2,94086,CNDC,COGENT,-121.989,37.3733,2020/5/24 11:50
2883,United States,"Sunnyvale, CA",1360 Kifer Road,Element Critical SV1 (fka. CentralColo),94086,CNDC,COGENT,-121.988,37.372,2020/5/24 11:50
2884,United States,"Sunnyvale, CA",1380 Kifer Road,CenturyLink (fka. Level 3),94086,CNDC,COGENT,-121.988,37.3734,2020/5/24 11:50
2885,United States,"SUNNYVALE, CA",1400 KIFER RD,Cyxtera SN1,94086,CNDC,COGENT,-121.986,37.3721,2020/5/24 11:50
2886,United States,"Sunnyvale, CA",255 Caspian Dr,Equinix SV4,94089,CNDC,COGENT,-122.015,37.4145,2020/5/24 11:50
2887,United States,"Sunnyvale, CA",444 Toyama 1st Floor,Equinix SV6 (fka. S&D),94089,CNDC,COGENT,-122.015,37.3996,2020/5/24 11:50
2888,United States,"Suwanee, GA",300 Satellite Blvd,QTS Datacenter,30024,CNDC,COGENT,-84.0612,34.0353,2020/5/24 11:50
2889,United States,"Syracuse, NY",100 Madison St,AXA Towers I,13202,COB,COGENT,-83.942,43.4661,2020/5/24 11:50
2890,United States,"Syracuse, NY",109 South Warren St,Westelcom,13202,CNDC,COGENT,-76.1504,43.0503,2020/5/24 11:50
2891,United States,"Syracuse, NY",109 South Warren St,State Tower Building,13202,COB,COGENT,-76.1504,43.0503,2020/5/24 11:50
2892,United States,"Syracuse, NY",109 South Warren St 9th Floor,Wholesale Carrier Services - 9th Floor,13202,CNDC,COGENT,-76.1504,43.0503,2020/5/24 11:50
2893,United States,"Syracuse, NY",110 West Fayette Street,One Lincoln Center,13202,COB,COGENT,-76.1531,43.049,2020/5/24 11:50
2894,United States,"Syracuse, NY",120 Madison St,AXA Towers II,13202,COB,COGENT,-83.942,43.4661,2020/5/24 11:50
2895,United States,"Syracuse, NY",2 Clinton Square,FirstLight (fka. Fibertech),13202,CNDC,COGENT,-76.1527,43.0501,2020/5/24 11:50
2896,United States,"Syracuse, NY",2 Clinton Square,The Atrium,13202,COB,COGENT,-76.1527,43.0501,2020/5/24 11:50
2898,United States,"Syracuse, NY",300 S State St,One Park Place,13202,COB,COGENT,-117.295,34.096,2020/5/24 11:50
2899,United States,"Tampa, FL",100 N Tampa St,100 N Tampa St,33602,COB,COGENT,-82.4583,27.9455,2020/5/24 11:50
2900,United States,"Tampa, FL",100 S Ashley Dr,Wells FArgo Center,33602,COB,COGENT,-82.4578,27.9445,2020/5/24 11:50
2901,United States,"Tampa, FL",101 E Kennedy Blvd,Bank of America Plaza,33602,COB,COGENT,-74.0695,40.7293,2020/5/24 11:50
2903,United States,"Tampa, FL",400 N Ashley Dr,Rivergate Tower,33602,COB,COGENT,-82.4606,27.9472,2020/5/24 11:50
2904,United States,"Tampa, FL",400 N Tampa St,Cogent Data Center - Tampa,33602,CDC,COGENT,-82.4591,27.9477,2020/5/24 11:50
2905,United States,"Tampa, FL",400 N Tampa St,WOW! Business,33602,CNDC,COGENT,-82.4591,27.9477,2020/5/24 11:50
2906,United States,"Tampa, FL",400 N Tampa St,400 N Tampa St,33602,COB,COGENT,-82.4591,27.9477,2020/5/24 11:50
2907,United States,"Tampa, FL",401 E Jackson St,SunTrust Financial Centre,33602,COB,COGENT,-79.2738,36.0948,2020/5/24 11:50
2908,United States,"Tampa, FL",412 East Madison Ave,Hostway,33602,CNDC,COGENT,-82.457,27.949,2020/5/24 11:50
2909,United States,"Tampa, FL",655 N Franklin St,MMR Cage 16,33602,COB,COGENT,-82.4584,27.9496,2020/5/24 11:50
2910,United States,"Tampa, FL",655 N Franklin St 1000,365 Data Centers (fka. Equinix - S&D),33602,CNDC,COGENT,-82.4584,27.9499,2020/5/24 11:50
2911,United States,"Tampa, FL",7909 Woodland Center Blvd - 1st Floor,CenturyLink (fka. Level 3),33614,CNDC,COGENT,-82.523,28.0185,2020/5/24 11:50
2912,United States,"Tampa, FL",8350 Park Edge Dr,Flexential North (fka. Peak 10 / ViaWest),33637,CNDC,COGENT,-82.3683,28.0738,2020/5/24 11:50
2913,United States,"Tampa, FL",9310 Florida Palm Drive,Cyxtera TPA1,33619,CNDC,COGENT,-82.3512,27.9453,2020/5/24 11:50
2914,United States,"Tampa, FL",9413 Corporate Lake Dr,Flexential TPA02,33634,CNDC,COGENT,-82.5405,28.0353,2020/5/24 11:50
2915,United States,"Tampa, FL",9417 Corporate Lake Drive,Flexential TPA01,33634,CNDC,COGENT,-82.54,28.0372,2020/5/24 11:50
2916,United States,"Tempe, AZ",1005 W Geneva Drive,Omnis Data Center,85282,CNDC,COGENT,-111.952,33.3944,2020/5/24 11:50
2917,United States,"Tempe, AZ",3011 S 52nd St,EdgeConneX EDCPHX01,85282,CNDC,COGENT,-111.97,33.3953,2020/5/24 11:50
2918,United States,"Toledo, OH",639 Oliver St,CenturyLink (fka. Level 3 / Wiltel),43609,CNDC,COGENT,-77.0415,41.2438,2020/5/24 11:50
2919,United States,"Totowa, NJ",50 Madison Rd.,CyrusOne Totowa,7512,CNDC,COGENT,-74.2424,40.9072,2020/5/24 11:50
2920,United States,"Troy, MI",101 W. Big Beaver Road,Columbia Center ll,48084,COB,COGENT,-83.1499,42.5608,2020/5/24 11:50
2921,United States,"Troy, MI",201 W Big Beaver Rd,Columbia Center I,48084,COB,COGENT,-83.1509,42.5608,2020/5/24 11:50
2922,United States,"Troy, MI",2855 Coolidge Highway,Troy Place,48084,COB,COGENT,-83.1881,42.559,2020/5/24 11:50
2923,United States,"Troy, MI",3001 W Big Beaver Road,Troy Place,48084,COB,COGENT,-83.189,42.5604,2020/5/24 11:50
2924,United States,"Troy, MI",3155 West Big Beaver Road,Troy Place,48084,COB,COGENT,-83.19,42.5594,2020/5/24 11:50
2925,United States,"Troy, MI",319 Executive Dr,ManagedWay Troy,48083,CNDC,COGENT,-83.0998,42.5371,2020/5/24 11:50
2926,United States,"Troy, MI",3221 W Big Beaver Road,Troy Place,48084,COB,COGENT,-83.1914,42.5593,2020/5/24 11:50
2927,United States,"Troy, MI",3250 W Big Beaver Rd,Sheffield Office Park l,48084,COB,COGENT,-83.1939,42.5622,2020/5/24 11:50
2928,United States,"Troy, MI",3290 W Big Beaver Rd,Sheffield Office Park lll,48084,COB,COGENT,-83.1944,42.5634,2020/5/24 11:50
2929,United States,"Troy, MI",3310 W Big Beaver Rd,Sheffield Office Park IV,48084,COB,COGENT,-83.1943,42.5647,2020/5/24 11:50
2930,United States,"Troy, MI",3331 W Big Beaver Rd,Cogent Data Center - Troy,48084,CDC,COGENT,-83.1932,42.5595,2020/5/24 11:50
2931,United States,"Troy, MI",3331 W Big Beaver Rd,Troy Place,48084,COB,COGENT,-83.1932,42.5595,2020/5/24 11:50
2932,United States,"Troy, MI",600 Executive Drive,ManagedWay Troy #2,48083,CNDC,COGENT,-80.0755,26.7117,2020/5/24 11:50
2933,United States,"Tukwila, WA",12101 Tukwila International Boulevard,Digital Fortress,98168,CNDC,COGENT,-122.294,47.4942,2020/5/24 11:50
2934,United States,"Tukwila, WA",12101 Tukwila International Boulevard,Digital Fortress,98168,CNDC,COGENT,-122.294,47.4942,2020/5/24 11:50
2935,United States,"Tukwila, WA",12101 Tukwila International Boulevard,Intergate West Building C,98168,COB,COGENT,-122.294,47.4942,2020/5/24 11:50
2936,United States,"Tukwila, WA",12301 Tukwila International Blvd TBD,Cyxtera SEA1 (fka. CenturyLink SE2),98168,CNDC,COGENT,-122.288,47.4859,2020/5/24 11:50
2937,United States,"Tukwila, WA",3311 S 120th Place,Sabey - Intergate.Seattle East Carrier MMR,98168,CNDC,COGENT,-122.291,47.4951,2020/5/24 11:50
2938,United States,"Tukwila, WA",3311 S 120th Place,zColo (fka. CoreLink),98168,CNDC,COGENT,-122.291,47.4951,2020/5/24 11:50
2939,United States,"Tukwila, WA",3355 S 120th Place,Sabey - Intergate.Seattle East Bldg. 5 SDC52,98168,CNDC,COGENT,-122.29,47.4919,2020/5/24 11:50
2940,United States,"Tukwila, WA",3355 S 120th Place,Cyxtera SEA1 (fka. CenturyLink SE3),98168,CNDC,COGENT,-122.29,47.4919,2020/5/24 11:50
2942,United States,"Tukwila, WA",3433 S 120th Place,Digital Realty (fka. Intergate East Bldg. 4),98168,CNDC,COGENT,-122.289,47.493,2020/5/24 11:50
2943,United States,"Tulsa, OK",322 E Archer Street,Tierpoint TUL (fka. Perimeter Technology),74120,CNDC,COGENT,-95.9886,36.1588,2020/5/24 11:50
2944,United States,"Tustin, CA",14351 Myford Rd,IC2Net,92780,CNDC,COGENT,-117.803,33.7171,2020/5/24 11:50
2945,United States,"Vienna, VA",1919 Gallows Rd,Tysons International Plaza Tower 1,22182,COB,COGENT,-77.2266,38.913,2020/5/24 11:50
2946,United States,"Vienna, VA",1921 Gallows Rd,Cogent Data Center - Vienna,22182,CDC,COGENT,-77.2273,38.9129,2020/5/24 11:50
2947,United States,"Vienna, VA",1921 Gallows Rd,Tysons International Plaza Tower 2,22182,COB,COGENT,-77.2273,38.9129,2020/5/24 11:50
2948,United States,"Vienna, VA",1951 Kidwell Dr,1951 Kidwell Dr,22182,COB,COGENT,-77.2205,38.9104,2020/5/24 11:50
2949,United States,"Vienna, VA",2070 Chain Bridge Rd,Tycon Courthouse,22182,COB,COGENT,-77.2374,38.9173,2020/5/24 11:50
2951,United States,"Vienna, VA",8150 Leesburg Pike,8150 Leesburg Pike,22182,COB,COGENT,-77.2266,38.9165,2020/5/24 11:50
2952,United States,"Vienna, VA",8229 Boone Blvd,Tycon 3,22182,COB,COGENT,-77.2311,38.9155,2020/5/24 11:50
2954,United States,"Vienna, VA",8245 Boone Blvd,Tycon 2,22182,COB,COGENT,-77.2326,38.9154,2020/5/24 11:50
2955,United States,"Waltham, MA",100 Fifth Avenue,Prospect Hill Office Park,2451,COB,COGENT,-73.9932,40.737,2020/5/24 11:50
2956,United States,"Waltham, MA",1000 Winter St,Bay Colony Corporate Center,2451,COB,COGENT,-71.2746,42.4075,2020/5/24 11:50
2957,United States,"Waltham, MA",1050 Winter St,Bay Colony Corporate Center,2451,COB,COGENT,-71.2767,42.4075,2020/5/24 11:50
2958,United States,"Waltham, MA",1100 Winter St,Bay Colony Corporate Center,2451,COB,COGENT,-71.2758,42.4094,2020/5/24 11:50
2959,United States,"Waltham, MA",115 2nd Ave,Cyxtera BO3,2451,CNDC,COGENT,-78.6417,40.4095,2020/5/24 11:50
2960,United States,"Waltham, MA",1601 Trapelo Rd,Reservoir Place,2451,COB,COGENT,-71.2587,42.4183,2020/5/24 11:50
2961,United States,"Waltham, MA",200 Fifth Avenue,Prospect Hill Office Park,2451,COB,COGENT,-73.9899,40.742,2020/5/24 11:50
2962,United States,"Waltham, MA",265 Winter Street 2,ColoSpace,2453,CNDC,COGENT,-71.2531,42.3999,2020/5/24 11:50
2963,United States,"Waltham, MA",300 Fifth Avenue,Cogent Data Center,2451,CDC,COGENT,-73.986,40.7468,2020/5/24 11:50
2964,United States,"Waltham, MA",300 Fifth Avenue,Prospect Hill Office Park,2451,COB,COGENT,-73.986,40.7468,2020/5/24 11:50
2965,United States,"Waltham, MA",580 Winter St,Cyxtera BO2,2451,CNDC,COGENT,-71.2674,42.3968,2020/5/24 11:50
2966,United States,"Waltham, MA",600 Winter Street,Cyxtera BO1,2451,CNDC,COGENT,-71.2683,42.397,2020/5/24 11:50
2967,United States,"Waltham, MA",74 West Street,Equinix BO1 (fka. S&D),2451,CNDC,COGENT,-84.4289,32.9169,2020/5/24 11:50
2968,United States,"Waltham, MA",950 Winter St,Bay Colony Corporate Center,2451,COB,COGENT,-71.2726,42.4064,2020/5/24 11:50
2969,United States,"Washington, DC",1 Thomas Circle NW,One Thomas Circle,20005,COB,COGENT,-77.0327,38.9052,2020/5/24 11:50
2970,United States,"Washington, DC",1000 Thomas Jefferson St.,1000 Thomas Jefferson St.,20007,COB,COGENT,-77.0605,38.9028,2020/5/24 11:50
2971,United States,"Washington, DC",1001 G St NW,Washington Center West,20001,COB,COGENT,-77.0265,38.8987,2020/5/24 11:50
2972,United States,"Washington, DC",1025 Thomas Jefferson St NW,Jefferson Court,20007,COB,COGENT,-77.0598,38.9032,2020/5/24 11:50
2973,United States,"Washington, DC",1050 Connecticut Ave NW,Washington Square,20036,COB,COGENT,-77.0407,38.9035,2020/5/24 11:50
2974,United States,"Washington, DC",1050 Thomas Jefferson St NW,1050 Thomas Jefferson St NW,20007,COB,COGENT,-77.0604,38.9035,2020/5/24 11:50
2975,United States,"Washington, DC",1055 Thomas Jefferson St NW,Foundry Bldg,20007,COB,COGENT,-77.0597,38.9038,2020/5/24 11:50
2976,United States,"Washington, DC",1100 New York Ave NW,1100 New York Ave NW,20005,COB,COGENT,-77.0276,38.9003,2020/5/24 11:50
2977,United States,"Washington, DC",1101 17th St NW,Morris B. Hariton Bldg,20036,COB,COGENT,-77.038,38.9041,2020/5/24 11:50
2978,United States,"Washington, DC",1101 New York Ave NW,1101 New York Ave NW,20005,COB,COGENT,-77.0276,38.9003,2020/5/24 11:50
2979,United States,"Washington, DC",1101 Pennsylvania Ave NW,1101 Pennsylvania Ave NW,20004,COB,COGENT,-77.0273,38.8955,2020/5/24 11:50
2980,United States,"Washington, DC",1110 Vermont Ave NW,1110 Vermont Ave,20005,COB,COGENT,-77.0331,38.9043,2020/5/24 11:50
2981,United States,"Washington, DC",1111 19th St NW,1111 19th St NW,20036,COB,COGENT,-77.0429,38.9044,2020/5/24 11:50
2982,United States,"Washington, DC",1120 20th Street NW,One Lafayette Centre,20036,COB,COGENT,-77.0454,38.9044,2020/5/24 11:50
2983,United States,"Washington, DC",1120 Vermont Ave NW,"Cogent Data Center - Washington, DC",20005,CDC,COGENT,-77.033,38.9049,2020/5/24 11:50
2984,United States,"Washington, DC",1120 Vermont Ave NW,1120 Vermont Ave NW,20005,COB,COGENT,-77.033,38.9049,2020/5/24 11:50
2985,United States,"Washington, DC",1133 21st Street NW,Two Lafayette Centre,20036,COB,COGENT,-77.0463,38.9045,2020/5/24 11:50
2986,United States,"Washington, DC",1133 Connecticut Ave NW,1133 Connecticut Ave NW,20036,COB,COGENT,-77.0402,38.905,2020/5/24 11:50
2987,United States,"Washington, DC",1140 Connecticut Ave NW,1140 Connecticut Ave NW,20036,COB,COGENT,-77.0411,38.9049,2020/5/24 11:50
2988,United States,"Washington, DC",1155 21st Street NW,Three Lafayette Centre,20036,COB,COGENT,-77.0463,38.9048,2020/5/24 11:50
2989,United States,"Washington, DC",1201 Connecticut Ave NW,1201 Connecticut Ave NW,20036,COB,COGENT,-77.0407,38.9061,2020/5/24 11:50
2990,United States,"Washington, DC",1201 Pennsylvania Ave NW,Heurich Bldg,20004,COB,COGENT,-77.0283,38.8955,2020/5/24 11:50
2991,United States,"Washington, DC",1220 19th St NW,1220 19th St NW,20036,COB,COGENT,-77.0439,38.9066,2020/5/24 11:50
2992,United States,"Washington, DC",1220 L St NW,1220 L St NW,20005,COB,COGENT,-77.029,38.9033,2020/5/24 11:50
2993,United States,"Washington, DC",1250 Connecticut Ave NW,1250 Connecticut Ave NW,20036,COB,COGENT,-77.0422,38.907,2020/5/24 11:50
2994,United States,"Washington, DC",1255 23rd St NW,Floyd D. Akers Bldg,20037,COB,COGENT,-77.0497,38.9063,2020/5/24 11:50
2995,United States,"Washington, DC",1275 K St NW,CoreSite DC1,20005,CNDC,COGENT,-77.0292,38.9028,2020/5/24 11:50
2996,United States,"Washington, DC",1275 K St NW,1275 K St NW,20005,COB,COGENT,-77.0292,38.9028,2020/5/24 11:50
2997,United States,"Washington, DC",1299 Pennsylvania Ave NW,The Warner Building,20004,COB,COGENT,-77.0288,38.8965,2020/5/24 11:50
2998,United States,"Washington, DC",1300 Eye St NW,Franklin Square,20005,COB,COGENT,-77.0301,38.9011,2020/5/24 11:50
2999,United States,"Washington, DC",1301 K St NW,One Franklin Square,20005,COB,COGENT,-77.0306,38.903,2020/5/24 11:50
3000,United States,"Washington, DC",1325 G St NW,1325 G St NW,20005,COB,COGENT,-77.0308,38.8988,2020/5/24 11:50
3001,United States,"Washington, DC",1350 Eye St NW,One Franklin Square,20005,COB,COGENT,-77.0313,38.901,2020/5/24 11:50
3002,United States,"Washington, DC",1401 Eye St NW,Franklin Tower,20005,COB,COGENT,-77.0324,38.9016,2020/5/24 11:50
3003,United States,"Washington, DC",1455 Pennsylvania Ave NW,Willard Office Bldg,20004,COB,COGENT,-77.0329,38.8969,2020/5/24 11:50
3004,United States,"Washington, DC",1615 L St NW,1615 L St NW,20036,COB,COGENT,-77.036,38.9039,2020/5/24 11:50
3005,United States,"Washington, DC",1620 L St NW,1620 L St NW,20036,COB,COGENT,-77.0376,38.9033,2020/5/24 11:50
3006,United States,"Washington, DC",1667 K St NW,Flagship of Farragut Square,20006,COB,COGENT,-77.0381,38.903,2020/5/24 11:50
3007,United States,"Washington, DC",1730 M St NW,1730 M St NW,20036,COB,COGENT,-77.0398,38.9055,2020/5/24 11:50
3008,United States,"Washington, DC",1747 Pennsylvania Ave NW,1747 Pennsylvania Ave NW,20006,COB,COGENT,-77.0408,38.8995,2020/5/24 11:50
3009,United States,"Washington, DC",1800 K St NW,1800 K St NW,20006,COB,COGENT,-77.0421,38.9021,2020/5/24 11:50
3010,United States,"Washington, DC",1801 K St NW,1801 K St NW,20006,COB,COGENT,-77.0419,38.9029,2020/5/24 11:50
3011,United States,"Washington, DC",1825 Eye St NW,International Square 3,20006,COB,COGENT,-77.0432,38.9023,2020/5/24 11:50
3012,United States,"Washington, DC",1828 L St NW,1828 L St NW,20036,COB,COGENT,-77.0429,38.9035,2020/5/24 11:50
3013,United States,"Washington, DC",1850 K St NW,International Square 1,20006,COB,COGENT,-77.0431,38.9023,2020/5/24 11:50
3014,United States,"Washington, DC",1875 Eye St NW,International Square 2,20006,COB,COGENT,-77.0429,38.9017,2020/5/24 11:50
3015,United States,"Washington, DC",1901 Pennsylvania Ave NW,1901 Pennsylvania Ave NW,20006,COB,COGENT,-77.0437,38.9007,2020/5/24 11:50
3016,United States,"Washington, DC",1909 K St NW,The Millennium Bldg,20006,COB,COGENT,-77.0441,38.9029,2020/5/24 11:50
3017,United States,"Washington, DC",1919 M St NW,1919 M St NW,20036,COB,COGENT,-77.0444,38.9059,2020/5/24 11:50
3018,United States,"Washington, DC",1919 Pennsylvania Ave NW,1919 Pennsylvania Ave NW,20006,COB,COGENT,-77.0443,38.9008,2020/5/24 11:50
3019,United States,"Washington, DC",2000 L St NW / 2001 K St NW,2000 L St NW,20036,COB,COGENT,-77.0457,38.9034,2020/5/24 11:50
3020,United States,"Washington, DC",2000 M St NW,2000 M St NW,20036,COB,COGENT,-77.0451,38.9055,2020/5/24 11:50
3021,United States,"Washington, DC",2000 Pennsylvania Ave NW,2000 Pennsylvania Ave NW,20006,COB,COGENT,-77.0456,38.9003,2020/5/24 11:50
3022,United States,"Washington, DC",2020 K St NW,2020 K St NW,20006,COB,COGENT,-77.0459,38.9022,2020/5/24 11:50
3023,United States,"Washington, DC",2100 M St NW,2100 M St NW,20037,COB,COGENT,-77.0469,38.9051,2020/5/24 11:50
3024,United States,"Washington, DC",2112 Pennsylvania Avenue NW,2112 Pennsylvania Avenue,20037,COB,COGENT,-77.0479,38.9014,2020/5/24 11:50
3025,United States,"Washington, DC",2200 Pennsylvania Ave NW,Square 54,20037,COB,COGENT,-77.0491,38.902,2020/5/24 11:50
3026,United States,"Washington, DC",2450 N Street NW,2450 N Street NW,20037,COB,COGENT,-77.0526,38.907,2020/5/24 11:50
3027,United States,"Washington, DC","3600 M Street, NW",Car Barn (Georgetown University),20007,COB,COGENT,-77.0697,38.9051,2020/5/24 11:50
3028,United States,"Washington, DC",400-444 N Capitol St NW,Hall of States Bldg,20001,COB,COGENT,-77.0094,38.8954,2020/5/24 11:50
3029,United States,"Washington, DC",401 9th St NW,Market Square North,20004,COB,COGENT,-92.4776,44.0265,2020/5/24 11:50
3030,United States,"Washington, DC",4200 Connecticut Ave,UDC,20008,COB,COGENT,-77.0654,38.9445,2020/5/24 11:50
3031,United States,"Washington, DC",555 12th St NW,District Center Building,20004,COB,COGENT,-77.0274,38.8968,2020/5/24 11:50
3032,United States,"Washington, DC",575 7th Street NW,575 7th Street NW,20004,COB,COGENT,-77.0217,38.8967,2020/5/24 11:50
3033,United States,"Washington, DC",600 14th St NW,Hamilton Square,20004,COB,COGENT,-77.0323,38.8977,2020/5/24 11:50
3034,United States,"Washington, DC",600 Maryland Ave SW,Capital Gallery,20024,COB,COGENT,-77.0209,38.8863,2020/5/24 11:50
3035,United States,"Washington, DC",600 New Hampshire Ave NW,Watergate 600,20037,COB,COGENT,-77.0551,38.8978,2020/5/24 11:50
3036,United States,"Washington, DC",601 13th St NW,Homer Bldg,20005,COB,COGENT,-77.0293,38.8979,2020/5/24 11:50
3037,United States,"Washington, DC",601 Pennsylvania Ave NW - North Bldg,Pennsylvania Building,20004,COB,COGENT,-77.0204,38.8936,2020/5/24 11:50
3038,United States,"Washington, DC",700 13th St NW,700 13th St NW,20005,COB,COGENT,-77.0298,38.8989,2020/5/24 11:50
3039,United States,"Washington, DC",701 13th Street - 700 12th Street,One Metro Center,20005,COB,COGENT,-90.588,41.4995,2020/5/24 11:50
3040,United States,"Washington, DC",701 Pennsylvania Ave NW,Market Square East,20004,COB,COGENT,-77.0229,38.8941,2020/5/24 11:50
3041,United States,"Washington, DC",801 Pennsylvania Ave NW,Market Square West,20004,COB,COGENT,-77.0236,38.8942,2020/5/24 11:50
3042,United States,"Washington, DC",900 17th St NW,Farragut Bldg,20006,COB,COGENT,-77.04,38.9015,2020/5/24 11:50
3043,United States,"Washington, DC",901 15th St NW (AKA 900 4th),McPherson Bldg,20005,COB,COGENT,-77.0334,38.9017,2020/5/24 11:50
3044,United States,"Weehawken, NJ",1919 Park Avenue,Cyxtera EWR2 (fka. CenturyLink NJ2X),7086,CNDC,COGENT,-73.9363,40.8081,2020/5/24 11:50
3045,United States,"Weehawken, NJ",300 JFK Boulevard East,Digital Realty NJR1 (fka. Telx MMR),7086,CNDC,COGENT,-74.0253,40.7615,2020/5/24 11:50
3046,United States,"Weehawken, NJ",300 JFK Boulevard East,Cyxtera EWR2 (fka. CenturyLink NJ2),7086,CNDC,COGENT,-74.0253,40.7615,2020/5/24 11:50
3047,United States,"West Chicago, IL",603 Discovery Dr,New Continuum,60185,CNDC,COGENT,-88.249,41.8758,2020/5/24 11:50
3048,United States,"WEST JORDAN, UT",3333 W 9000 S,Aligned Energy,84088,CNDC,COGENT,-111.97,40.5859,2020/5/24 11:50
3049,United States,"West Jordan, UT",7202 S Campus View Dr,Flexential South Valley (fka. Peak 10 / ViaWest),84084,CNDC,COGENT,-111.986,40.6202,2020/5/24 11:50
3050,United States,"West Palm Beach, FL",424 Hampton Rd,Cloud South Data Center,33405,CNDC,COGENT,-80.0555,26.6836,2020/5/24 11:50
3052,United States,"Westland, MI",6435 North Hix Road,Online Tech Data Center,48186,CNDC,COGENT,-83.4195,42.3311,2020/5/24 11:50
3053,United States,"White Plains, NY",1 North Broadway,White Plains Plaza - South Tower,10601,COB,COGENT,-96.944,38.6708,2020/5/24 11:50
3054,United States,"White Plains, NY",1 North Lexington Avenue,Gateway Center One,10601,COB,COGENT,-73.7733,41.0322,2020/5/24 11:50
3055,United States,"White Plains, NY",1-11 Martine Avenue,Pace University,10606,COB,COGENT,-73.7694,41.031,2020/5/24 11:50
3056,United States,"White Plains, NY",10 Bank Street,10 Bank Street,10606,COB,COGENT,-73.774,41.0298,2020/5/24 11:50
3057,United States,"White Plains, NY",360 Hamilton Avenue,360 Hamilton Avenue,10601,COB,COGENT,-73.7671,41.0347,2020/5/24 11:50
3058,United States,"White Plains, NY",44 South Broadway,Westchester One,10601,COB,COGENT,-73.7616,41.0309,2020/5/24 11:50
3059,United States,"White Plains, NY",445 Hamilton Avenue,White Plains Plaza - North Tower,10601,COB,COGENT,-81.49,31.1914,2020/5/24 11:50
3060,United States,"White Plains, NY",50 Main Street,Westchester Financial Center,10606,COB,COGENT,-75.5979,41.6489,2020/5/24 11:50
3061,United States,"Wilmington, DE",1201 N Market St,IPR International,19801,CNDC,COGENT,-75.5468,39.7479,2020/5/24 11:50
3062,United States,"Wilmington, DE",1201 N Market St,Chase Manhattan Centre,19801,COB,COGENT,-75.5468,39.7479,2020/5/24 11:50
3063,United States,"Wilmington, DE",1313 N Market St,Hercules Plaza,19801,COB,COGENT,-75.5463,39.749,2020/5/24 11:50
3064,United States,"Worcester, MA",100 Front Street,Worcester Office I,1608,COB,COGENT,-91.4441,33.7041,2020/5/24 11:50
3065,United States,"Worcester, MA",120 Front Street,Worcester Office II,1608,COB,COGENT,-96.756,40.2653,2020/5/24 11:50
3066,United States,"Worcester, MA",474 Main Street,Crown Castle (fka. Lightower),1608,CNDC,COGENT,-73.207,42.177,2020/5/24 11:50`
)

func main() {
	cache := make(map[string]interface{}, 0)
	var addres []*models.GeographicAddress
	var sites []*models.GeographicSite
	reader := csv.NewReader(strings.NewReader(from))
	if records, err := reader.ReadAll(); err == nil {
		for _, record := range records {
			country := record[1]
			if country != "China" && country != "Canada" {
				continue
			}
			city := record[2]
			key := fmt.Sprintf("%s/%s", country, city)
			addr := record[3]
			building := record[4]
			zipCode := record[5]
			longitude := record[8]
			latitude := record[9]
			if _, ok := cache[key]; !ok {
				cache[key] = struct{}{}
				a := handler.GeographicAddress()
				a.FieldedAddress.Country = country
				a.FieldedAddress.City = city
				a.FieldedAddress.GeographicSubAddress[0].BuildingName = building
				a.FieldedAddress.Postcode = zipCode
				a.FieldedAddress.StreetName = addr
				a.FormattedAddress = handler.FormattedAddress(a.FieldedAddress)
				a.GeographicLocation.GeographicPoint[0].Latitude = swag.String(latitude)
				a.GeographicLocation.GeographicPoint[0].Longitude = swag.String(longitude)
				s := handler.GeographicSite()
				s.FieldedAddress = a.FieldedAddress
				s.FormattedAddress = a.FormattedAddress
				s.ReferencedAddress = a.ReferencedAddress
				addres = append(addres, a)
				sites = append(sites, s)
			}
		}
	} else {
		fmt.Println(err)
	}
	fmt.Println(util.ToString(addres))
	fmt.Println(util.ToString(sites))
}

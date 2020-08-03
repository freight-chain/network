# What is an "EDI" Transaction? 
www.freighttrust.com
www.t.me/freighttrust



## Formats of EDI

EDI works based on standards which determine how each message should be formatted.

Four EDI standards exist: UN/EDIFACT, which is the only internationally-recognized standard, used mostly outside of North America; ANSI ASC X12, used within North America, TRADACOM, used by British retail companies, and ODETTE, which is used by European automakers.


**Here is a small EDI sample, an X12 interchange containing a single 824 Application Advice transaction.**


```
ISA*00*          *00*          *08*9254110060     *ZZ*123456789      *041216*0805*U*00501*000095071*0*P*>~
GS*AG*5137624388*123456789*20041216*0805*95071*X*005010~
ST*824*021390001*005010X186A1~
BGN*11*FFA.ABCDEF.123456*20020709*0932**123456789**WQ~
N1*41*ABC INSURANCE*46*111111111~
PER*IC*JOHN JOHNSON*TE*8005551212*EX*1439~
N1*40*SMITHCO*46*A1234~
OTI*TA*TN*NA***20020709*0902*2*0001*834*005010X220A1~
SE*7*021390001~
GE*1*95071~
IEA*1*000095071~
```
### Here is a that EDI transaction in JSON format
```
{
  "interchanges": [
    {
      "ISA": "Interchange Control Header",
      "ISA_01_AuthorizationQualifier": "00",
      "ISA_02_AuthorizationInformation": "          ",
      "ISA_03_SecurityQualifier": "00",
      "ISA_04_SecurityInformation": "          ",
      "ISA_05_SenderQualifier": "08",
      "ISA_06_SenderId": "9254110060     ",
      "ISA_07_ReceiverQualifier": "ZZ",
      "ISA_08_ReceiverId": "123456789      ",
      "ISA_09_Date": "041216",
      "ISA_10_Time": "0805",
      "ISA_11_StandardsId": "U",
      "ISA_12_Version": "00501",
      "ISA_13_InterchangeControlNumber": "000095071",
      "ISA_14_AcknowledgmentRequested": "0",
      "ISA_15_TestIndicator": "P",
      "functional_groups": [
        {
          "GS": "Functional Group Header",
          "GS_01_FunctionalIdentifierCode": "AG",
          "GS_02_ApplicationSenderCode": "5137624388",
          "GS_03_ApplicationReceiverCode": "123456789",
          "GS_04_Date": "20041216",
          "GS_05_Time": "0805",
          "GS_06_GroupControlNumber": "95071",
          "GS_07_ResponsibleAgencyCode": "X",
          "GS_08_Version": "005010",
          "transactions": [
            {
              "824": "Application Advice",
              "ST": "Transaction Set Header",
              "ST_01_TransactionSetIdentifierCode": "824",
              "ST_02_TransactionSetControlNumber": "021390001",
              "ST_03_ImplementationConventionReference": "005010X186A1",
              "segments": [
                {
                  "BGN": "Beginning Segment",
                  "BGN_01": "11",
                  "BGN_02": "FFA.ABCDEF.123456",
                  "BGN_03": "20020709",
                  "BGN_04": "0932",
                  "BGN_06": "123456789",
                  "BGN_08": "WQ"
                },
                {
                  "N1": "Party Identification",
                  "N1_01": "41",
                  "N1_02": "ABC INSURANCE",
                  "N1_03": "46",
                  "N1_04": "111111111"
                },
                {
                  "PER": "Administrative Communications Contact",
                  "PER_01": "IC",
                  "PER_02": "JOHN JOHNSON",
                  "PER_03": "TE",
                  "PER_04": "8005551212",
                  "PER_05": "EX",
                  "PER_06": "1439"
                },
                {
                  "N1": "Party Identification",
                  "N1_01": "40",
                  "N1_02": "SMITHCO",
                  "N1_03": "46",
                  "N1_04": "A1234"
                },
                {
                  "OTI": "Original Transaction Identification",
                  "OTI_01": "TA",
                  "OTI_02": "TN",
                  "OTI_03": "NA",
                  "OTI_06": "20020709",
                  "OTI_07": "0902",
                  "OTI_08": "2",
                  "OTI_09": "0001",
                  "OTI_10": "834",
                  "OTI_11": "005010X220A1"
                }
              ]
            }
          ]
        }
      ]
    }
  ]
}
```



The implementation of EDI is important for companies as it can significantly reduce the cost of sending documents.

**EDI Costs Versus Benefits**
A paper purchase order requires resources to print the document, fax it, or post it to the vendor. EDI automatically sends the electronic document to the vendor thus reducing the cost of sending the PO. Studies of the cost savings of implementing EDI have been performed, including a report from the Aberdeen group in 2008, which highlighted that in the US it cost $37.45 to produce and send a paper PO, while it only cost $23.83 to send it using EDI.

Not all companies use EDI. There is a cost to implement and maintain the technology required to perform EDI. Each trading partner that a company wants to use EDI with may require resources to set up and this can be cost-prohibitive for smaller companies or companies without technical resources.


EDI How It Works
There are a number of ways EDI messages are transmitted between trading partners. The most common method was to *use a value-added network or VAN.* This allowed companies to send a transmission which was then reviewed by the VAN and then sent to the correct recipient.

**Think of Freight Trust Network as a Value Added Network**


More recently a new method for EDI transmission is being used. This is called AS2, which stands for Applicability Statement 2, and was championed by Wal-Mart, who requires all of their vendors to use this method. Using AS2, the EDI documents are transmitted across the internet and the security of the document is achieved by encryption and the use of digital certificates.

There are dozens of EDI documents that can be implemented by a company and their trading partners. Under the ANSI ASC X12 standard, EDI documents are part of a series, for example, such as an order series, a warehousing series, or a financial series.

Many companies will only implement a small number of EDI documents with their trading partners, commonly in the ordering series, material handling series and the delivery series.

**We have over 20,000 mappings and counting, think of us not only as a Value Added Network, but a Marketplace where you can quickly get and start selling to other businesses**

For example a company who is implementing EDI between themselves and a third party logistics company may only implement five EDI documents such as an EDI 940 for a warehouse shipping order, EDI 943 for a warehouse stock transfer shipment advice, EDI 944 for a warehouse stock transfer receipt advice, EDI 945 for a warehouse shipping advice, and EDI 947 for a warehouse inventory adjustment advice.

# Supported EDI Transactions

## Fully Supported 
204 Motor Carrier Shipment Information - Version 4010

210 Freight Details and Invoice - Version 4010

211 Motor Carrier Bill of Lading - Version 4010

214 Shipment Status Message - Version 4010

997 Functional Acknowledgement



# Partial List of additional EDI Transactions
Frequently Used EDI Transactions In Supply Chain
753 Request for Routing Instructions

754 Routing Instructions

816 Organizational Relationships

818 Commission Sales Report

830 Planning Schedule w/ Release Capability

832 Price/Sales Catalog

840 Request for Quotation

841 Specifications/Technical Information

842 Nonconformance Report

843 Response to Request for Quotation

845 Price Authorization Acknowledgment/Status

846 Inventory Inquiry/Advice

847 Material Claim

848 Material Safety Data Sheet

850 Purchase Order

851 Asset Schedule

852 Product Activity Data

853 Routing and Carrier Instruction

855 Purchase Order Acknowledgment

856 Ship Notice/Manifest

857 Shipment and Billing Notice

860 Purchase Order Change Request – Buyer Initiated

861 Receiving Advice/Acceptance Certificate

862 Shipping Schedule

863 Report of Test Results

865 Purchase Order Change Acknowledgment/Request – Seller Initiated

866 Production Sequence

869 Order Status Inquiry

870 Order Status Report

873 Commodity Movement Services

874 Commodity Movement Services Response

878 Product Authorization/De-authorization

879 Price Information

882 Direct Store Delivery Summary Information

885 Retail Account Characteristics

888 Item Maintenance

889 Promotion Announcement

890 Contract & Rebate Management

893 Item Information Request

895 Delivery/Return Acknowledgment or Adjustment

940 Warehouse Shipping Order

943 Warehouse Stock Transfer Shipment Advice

944 Warehouse Stock Transfer Receipt Advice

945 Warehouse Shipping Advice

947 Warehouse Inventory Adjustment Advice


www.freighttrust.com
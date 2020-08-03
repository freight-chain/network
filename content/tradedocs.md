# Freight Trust & Clearing

<!-- COPYRIGHT 2020 - FREIGHTTRUST AND CLEARING CORPORATION
        ALL RIGHTS RESERVED 
    -->
<!-- FREIGHT TRUST HEADER AREA DEFAULTS  -->
<!-- FREIGHT TRUST BANNER IMAGE -->
<p   align="center">

<img   src="https://raw.githubusercontent.com/freight-trust/branding/master/images/optimized_github_repo_card.png">
</p>
<br>
<!-- FREIGHT TRUST BANNER IMAGE -->
<!-- Badges Start -->
<p   align="center">
<img   alt="Twitter Follow"   src="https://img.shields.io/twitter/follow/freighttrustnet?label=%40FreightTrustNet&style=social">
 <a   href="https://t.me/freighttrust">
 <img   alt="Join Freight Trust's Public Telegram"   src="https://img.shields.io/badge/telegram-%40freighttrust-blue">
 </a>
<img   src="https://img.shields.io/github/package-json/keywords/freight-trust/tdocs">
</p>
<!-- Badges End -->
<!-- FREIGHT TRUST HEADER AREA DEFAULTS END -->

## Overview 


- [Introducing TradeDocs](https://github.com/freight-trust/tdocs)
- Releases of new services
    - TradeDocs
    - TradeDocs: FormGen
    - TradeSure 
    - Burgess - Jurisdictional Regimes and Standardized Parameters

- Roadmap Update
    - Customs Integration 
        - Production Ready by late September 
        - That is for United States, Canada 
        - [Mexico Integration, pidemento](https://ft-customs-mx.netlify.app/)
            - Q1 2021 No later (high priority)
        - Japan Integration 
            - Q4 2020 [japan_integration](/qKw7VEo0RFWil5qNXJfZsA)
        - Singapore 
            - Q4 2020
        - Australia 
            - Mid Q1 2021
        - EuroZone
            - Q2 2021 (at the latest)
        - United Kingdom 
            - Q2 2021 (at the latest)
    - SAFERnet Accessibility
- Updated Supported X12 Transaction Segments 
    - [Full list of supported X12 can be found here](https://x12.netlify.app/)

- Community Updates and Discussion 
    - Please Visit [github/freight-chain/network](https://github.com/freight-chain/network)
    - 
- New Staking RFC proposal and Masternode changes 
    - New Structures for Clusters 
    - Reward Structure 
        - [github.com/freight-chain/rfc/issues/22](https://github.com/freight-chain/rfc/issues/22)

- Branding and Corporate Identity 
    - @freight-trust/carbon 
    - New Resources
        - Figma Support
    - Style Guides 
        - General
        - GitHub Public Repos. 
    - Boilerplate Assets
    - Components [.svg, .webp]





---

# Summary Update

In this update we will introduce a few new `modules` and `interfaces`, and how this all ties together.

## Too Long, Dont Understand

We can take a shipment information (through EDI/XML/JSON), convert it the necessary parts through our mappings, generate an NFT-based smart contract that uses our library of terms, conditions, events, and types. We have the option of even converting the `token` into a `physical document` (we call this `retokenization/detokenization`), **which means we can support** **100% of all shipments, period**.

- Here is the Bill of Lading Artefact for a Generated Physical Document: [https://github.com/freight-trust/tradedocs-formgen/blob/master/assets/bol-mock-data.json](https://github.com/freight-trust/tradedocs-formgen/blob/master/assets/bol-mock-data.json)

- Here is the template for the necessary information needed for a legal bill of lading: [https://github.com/freight-trust/tradedocs-formgen/blob/master/views/bol.hbs](https://github.com/freight-trust/tradedocs-formgen/blob/master/views/bol.hbs)


- Here are the `TypeMappings` for the RAW EDI X12 Message: [https://github.com/freight-trust/protocol/blob/master/ROOT/TypeMaps.json](https://github.com/freight-trust/protocol/blob/master/ROOT/TypeMaps.json)


- Here are the `BillofLadingModel` Schema for interfacing: [https://github.com/freight-trust/protocol/blob/master/ROOT/BillOfLadingModel.json](https://github.com/freight-trust/protocol/blob/master/ROOT/BillOfLadingModel.json)

- Here are the `BOLWorkFlow` Schema for inputing your `legal terms` into the `cominbator`, this enriches the RAW EDI/data to include infromation not typically found in the B2B stuff (we attach this to the "trading channel" for the two businesses)  [https://github.com/freight-trust/protocol/blob/master/ROOT/BOLWorkflow.json](https://github.com/freight-trust/protocol/blob/master/ROOT/BOLWorkflow.json)

- Here are the `Solidity` Smart contracts Library used to create the actual `NFT-token` of the data [https://github.com/freight-trust/instruments/tree/master/contracts](https://github.com/freight-trust/instruments/tree/master/contracts)
    - Note: There are additional points of authentication of signatures necessary in some jurisdictions, such as the EuroZone. You can see the signature app for E.U. Digital Signatures here: [eurosign.netlify.app/](https://eurosign.netlify.app/)
<br>
- 100% Legally Compliant and Legal Parity. 
- Better than Bolero, essDocs, Komgo, Voltron, TradeLens, or any other system. 
- Commercial Invoice's get to be factored or financed. DeFi for Enterprise Supply Chain. 
- Commodities - Wheat, LNG, Gas, etc. 
- 


### Workflow Example



```sequence
Shippment Information->Protocol Pipe: Necessary Information
Note right of Protocol: Protocol Parses raw EDI
Protocol Pipe-->Shippment Information: Validates
Protocol Pipe -->Protocol Service: Sends JSON Object
Protocol Service -->Network Service: Generates Contract from JSON Object 
Network Service -->Protocol Service: Contract Deployed
Protocol Service --> Protocol: Instrument Created at 0x....
Note right of Protocol: ABI.json and TDOC.json match and verify
Note left of Shippment Information: Shippment Information Polling (async)
Shippment Information->Protocol: Polling Job Status
Protocol->Shippment Information: Success
```


---
<br />
<br />
<br />
<br />
<br />



---

## Releases 


## Burgess 
Burgess is a library for quickly implementing Business Rule Processes and Decision Notional Models into defined trading channel/trading partner relations to enable seamless counter-party trading.  

Think of Burgess as a library of what you can and can't legally do, what the terms and conditions are, and which choices you can make based on your own defined business policies and regulatory obligations, at scale. 

#### Purpose 
Burgess hosts the standard `consortium` agreements that could be used. Burgess is not the `rulebook` the `rulebook` is the terms and conditions that all network participants must abide by and are legally obligated to observe.

> **Burgess is the source repository of `components` that make up a legal document.**

### Example

See a complete `json` artifact [github.com/freight-trust/burgess/blob/master/logic/Global.json](https://github.com/freight-trust/burgess/blob/master/logic/Global.json)





```
Next Section, TDocs
```
## TDocs 
> note, `.tdocs`/`tdocs` is the actual usage of `TradeDocs`. `TradeDocs` is a separate repository than the one discussed here, you can see it here: [github.com/freight-trust/tradedocs](https://github.com/freight-trust/tradedocs)

###### tags: tradedocs, tdocs, documents of trade, instrumentization 
TradeDocs`StandardDocument` (abbreviated as `StanDoc`) is a family of document models that forms the basis of Freight Trust’s `TradeDoc`, which is used to represent our legal parity digital smart contracts. These documents are based off of ISO Standards and similarly structured standards. It itself is an instantiation of [`BasicDocument` (`BasicDoc`)](https://github.com/metanorma/basicdoc-models)

These are the principle `document structures` that form the basis for our `primitives`.

`TradeDocs` itself is a document model can form the basis of multiple standard document types, including:

-   Bills of Lading
    
-   Warehouse Receipts
    
-   Letters of Credit
    
-   Forward Contracts
    
-   Advanced Manifests
    
-   Bill of Materials
    


The Relaton Bibliography UML models used within these models are located here: [Relaton BibModels](https://github.com/metanorma/relaton-models). This is used for referencing `jursidictional_regimes` from the `burgess` repo. 

See the main `ROOT` Repository
[github.com/freight-trust/tdocs](https://github.com/freight-trust/tdocs)

See the main `toolchain` Repository
[github.com/freight-trust/tdocs/tree/master/src](https://github.com/freight-trust/tdocs/tree/master/src)

[](/wPvjva3OQ7-0uXZ2-9-a4Q#TradeDocs)


### Examples

[See full examples @github/freight-trust/tdocs/master/src](https://github.com/freight-trust/tdocs/tree/master/src)

#### Main Document
![image alt](https://raw.githubusercontent.com/freight-trust/tdocs/master/src/images/StandardDoc_StandardDocument.png "title" =800x500)

####     Sections
![image alt](https://raw.githubusercontent.com/freight-trust/tdocs/master/src/images/StandardDoc_Sections.png "title" =500x300)

### Roadmap
- DMN Integraton (in progress)
- BPMN Automation Workflow
- Library Registry for standards (only VICS supported right now)



```
TradeDocs Form Generator 
```


## TradeDocs: FormGen
###### tags: tradedocs, tdocs, documents of trade, form generation
> Generate PDFs from `json` objects, like `abi.json`

[github/freight-trust/tradedocs-formgen](https://github.com/freight-trust/tradedocs-formgen)

FormGen is a PDF Generator that utilizes `json objects` as an input. For example we can take a `abi.json` from a `smart contract` deployment that meets our schema and definitions as prescribed by the `tradedocs/tdocs` standard. The schema is defined as a `.hbs` file, and the corresponding solidity contracts are found in github.com/freight-trust/instruments

#### Roadmap 
- Integration of `@freight-trust/customs` packages
    - You will be able to `import` a package, for example `@freight-trust/customs-united-states` and it would provide you with the appropriate boilerplate manifests, etc. 


- Support for non-Latin Characters  (major)
    - two phases: 
        - simple font support 
        - document generation support 

- Expanded template generation and create library for boilerplate templates (minor)

- Ingest `.xsd` files to generate or enable porting into `.hbs` format. (major)

- Enable Logging outputs and metrics via logging module or export interface such as Prometheus  (major)

#### Support 
Long Term Support is available (LTS)

**Authenticated Container** can be found at: 
```bash
$ docker pull quay.io/freight/formgen@sha256:e68f61631a2e17934490bcd217ba02ef8e07804ff8a91aceaf96bf5744ad837e
```
[sha256:e68f61631a2e17934490bcd217ba02ef8e07804ff8a91aceaf96bf5744ad837e]

**Source Code** (Apache 2.0)
[github/freight-trust/tradedocs-formgen](https://github.com/freight-trust/tradedocs-formgen)


<br />

```
TradeSure - Automated First Pass Compliance Check
```
## TradeSure
###### tags: tradedocs, tdocs, documents of trade, automation, compliance, regulatory, sanctions, embargo

> Compliance Tool for OFAC/Sanctions
> 
TradeSure combines and indexes Governmental Regulations against Designated Nationals/Legal Entities for compliance.
[Demo Instance Available here](http://44.226.206.38:8084/)

#### Purpose 
A first-pass compliance tool that automates compliance obligations for AML/KYC/KYT regulations. 

#### Example 

1. Go to [TradeSure Demo Instance](http://44.226.206.38:8084/)
2. Enter a search term, for example `abu`
3. See results 

#### Notes:
Currently tradesure pulls from [github.com/freight-trust/tradedocs/tree/master/compliance](https://github.com/freight-trust/tradedocs/tree/master/compliance)

#### Roadmap 
- Ingest non-latin Character searches (major)
- DMN Integration 
    - Boilerplate Rules (FEEL Grammar)
    - Direct Apache NiFi Integration (priority)
- Production Deployment 
    - Branding Layer
    - Crypto Transaction Support 
    - DMN/BPMN Suite







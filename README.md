# Coronavirus Tracker App's Webpage

This was sprint project to hash out the code and release an (Android & iOS) app in their corresponding Play/App Store over the weekend.


### Project Repos

* ***App***: [ylogx/oneplanet-corona-app](https://github.com/ylogx/oneplanet-corona-app)
  - The app is written in dart and works on both Android & iOS. App supports dark & light mode.
* API Backend: [ylogx/oneplanet-corona-backend](https://github.com/ylogx/oneplanet-corona-backend)
  - The app reads data from backend REST API written in GoLang, containerized using [Docker][docker] and deployed on Google Cloud using [Cloud Run][gc-cloud-run].
* App Webpage: [ylogx/oneplanet-corona-web](https://github.com/ylogx/oneplanet-corona-web)
  - The app's webpage is a requirement for submitting the app to App/Play Store ([privacy page][privacy-page], tos page, etc.)

---

### Development

To run unit test, use: `make test`. To build the docker image, use: `make build_gcr_docker`.  
To deploy on Google Cloud Run, use: `make gcr_pipeline`. There are tons of useful commands in Makefile.

### Deployment

Deployed at: [/privacy][privacy-page]

[gc-cloud-run]: https://cloud.google.com/run/
[docker]: https://shubham.chaudhary.xyz/blog/docker/101
[privacy-page]: https://corona.oneplanet.us/privacy

  const CURRENT_URL = new URL(window.location.href)

  let PROTOCOL_HTTP = "https://"
  let DOMAIN_NAME = CURRENT_URL.hostname
  let URL_API = ""
  let URL_APP = ""


  if(DOMAIN_NAME === "localhost"){
    DOMAIN_NAME = DOMAIN_NAME +":3000"
    PROTOCOL_HTTP = "http://"
  }
  URL_API = `${PROTOCOL_HTTP}/${DOMAIN_NAME}/api`
  URL_APP = `${PROTOCOL_HTTP}/${DOMAIN_NAME}`

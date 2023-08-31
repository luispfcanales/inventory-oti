const CURRENT_URL = new URL(window.location.href)

let STREAM_URL = ""
let STREAM_COMPUTERS = ""

let PROTOCOL_HTTP = "https://"
let DOMAIN_NAME = CURRENT_URL.hostname
let URL_API = ""
let URL_APP = ""
let PORT = window.location.port

if(PORT === "3000"){
  PROTOCOL_HTTP = "http://"
  DOMAIN_NAME = DOMAIN_NAME + ":"+PORT
}

URL_API = `${PROTOCOL_HTTP}/${DOMAIN_NAME}/api`
URL_APP = `${PROTOCOL_HTTP}/${DOMAIN_NAME}`

STREAM_URL = `${PROTOCOL_HTTP}/${DOMAIN_NAME}/stream`
STREAM_COMPUTERS = `ws://${DOMAIN_NAME}/stream/admin/computer`


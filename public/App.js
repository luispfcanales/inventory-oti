import router from './router/router.js'
import pageNotFound from './views/error404/error404.js'
import pageAdmin from './views/admin/admin.js'


export function App(){
  const $rootID = document.getElementById("rootID");

  //clear page content
  $rootID.innerHTML=null

  switch (router()) {
    case "":
      window.location = "#/";break;
    case "#/":
      $rootID.appendChild(pageAdmin());break;
    default:
      $rootID.appendChild(pageNotFound());break;
  }
}

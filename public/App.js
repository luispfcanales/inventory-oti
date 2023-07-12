import router from './router/router.js'
import pageNotFound from './views/error404/error404.js'
import pageAdmin from './views/admin/admin.js'
import pageAdminComputers from './views/admin/computers/computers.js'

//scripts
import { AppendScriptAdmin,AppendScript } from './scripts/scripts.js'

export function App(){
  //const $rootID = document.getElementById("rootID");
  const $rootID = document.body

  //clear page content
  $rootID.innerHTML=null

  switch (router().toLowerCase()) {
    case "":
      window.location = "#/admin";break;
    case "#/":
      window.location = "#/admin";break;
    case "#/admin":
      $rootID.appendChild(pageAdmin());
      AppendScript("session")
      break;
    case "#/admin/equipos-informaticos":
      $rootID.appendChild(pageAdminComputers())
      AppendScript("session")
      AppendScriptAdmin("modal_computers")
      AppendScriptAdmin("table_computers")
      break;
    default:
      $rootID.appendChild(pageNotFound());break;
  }
}

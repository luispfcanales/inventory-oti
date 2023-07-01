// PAGE ADMIN RENDER
import {NavHeader , NavBody} from '../../components/navbar/navbar.js'

export default function () {
  const divContent = document.createElement("div")
  divContent.appendChild(NavHeader())
  divContent.appendChild(NavBody())
  return divContent
}

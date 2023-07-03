// PAGE ADMIN/COMPUTERS RENDER
import {NavHeader , NavBody} from '../../../components/navbar/navbar.js'
import {ModalEquipos} from '../../../components/modal/modal.js'
import {ButtonModalElement} from '../../../components/modal/button.js'
import { tableComputer } from '../../../components/table/table.js'

//import variables
import { 
  ID_CONTAINER_COMPUTERS_MODAL,
  ID_CLOSE_COMPUTERS_MODAL,
  ID_BUTTON_OPEN_MODAL,
} from './variables.js'


function sectionComputers(){
  const body = document.createElement("div")
  body.className = "px-4 overflow-y-auto bg-white"

  const container = document.createElement("div")
  container.className = "flex w-full"
  container.appendChild(tableComputer(
    "body_table_computer",
    [
      "key",
      "arquitectura",
      "disco",
      "fabricante",
      "modelo",
      "nombre equipo",
      "procesador",
      "ram",
      "serial",
      "tamano disco",
      "OPCIONES"
    ]
  ))

  body.appendChild(ButtonModalElement(ID_CONTAINER_COMPUTERS_MODAL,ID_BUTTON_OPEN_MODAL))
  body.appendChild(container)
  return body
}

export default function () {
  const divContent = document.createElement("div")
  const wrapper = document.createElement("div")
  wrapper.className = "pt-10 bg-white my-2 mx-10 shadow-md rounded-md"

  divContent.appendChild(NavHeader())
  divContent.appendChild(NavBody())
  wrapper.appendChild(ModalEquipos(ID_CONTAINER_COMPUTERS_MODAL,ID_CLOSE_COMPUTERS_MODAL))
  wrapper.appendChild(sectionComputers())
  divContent.appendChild(wrapper)
  //divContent.appendChild(ButtonModalElement(ID_CONTAINER_COMPUTERS_MODAL,ID_BUTTON_OPEN_MODAL))



  return divContent
}

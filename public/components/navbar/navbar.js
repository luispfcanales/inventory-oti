import button from '../../components/button.js'
import div from '../../components/div.js'
import image from '../../components/image.js'
import span from '../../components/span.js'
import {SvgRound , SvgAccess , SvgMaintenance , SvgPeople , SvgComputer } from '../../components/svg.js'
import link from '../../components/link.js'

let items_primary = [
  image(
    "w-7 h-7 rounded-full",
    "https://static.vecteezy.com/system/resources/previews/002/002/427/original/man-avatar-character-isolated-icon-free-vector.jpg",
  ),
  span("","absolute right-0 -mb-0.5 bottom-0 w-2 h-2 rounded-full bg-rose-500 border border-white dark:border-gray-900")
]

let items = [
  span(
    "",
    "relative flex-shrink-0",
    items_primary,
  ),
  span(
    "Nombre del Usuario",
    "ml-2",
  ),
]

let item_wrap = [
  button(
    "Perfil",
    "px-2 w-full text-left p-1 hover:bg-gray-100 hover:text-gray-700",
  ),
  button(
    "Salir",
    "px-2 w-full text-left p-1 hover:bg-gray-100 hover:text-gray-700",
  ),
]
export function NavHeader() {
  const containerHeader = document.createElement("div")
  containerHeader.className = "h-16 lg:flex w-full border-b border-gray-200 dark:border-gray-800 px-10 sticky top-0 z-50 bg-[var(--header-color)] relative"


  let contentWrap = div(
    "text-sm w-[150px] left-0 text-gray-500 absolute top-[3rem] hidden group-hover:block bg-white shadow-md rounded-md",
    item_wrap,
  )

  let rowIcon = div(
    "absolute top-[4rem] right-5 p-2 flex flex-col gap-2",
  )

  //container
  let divContent = div(
    "ml-auto flex items-center py-4 float-right group relative",
    [
      button(
        "",
        "flex items-center text-white",
        items,
        SvgRound(),
      ),
      contentWrap,
      rowIcon,
    ]
  )

  containerHeader.appendChild(divContent)

  return containerHeader
}

//CreateMenu return html element menu navbar
function CreateMenu(btn,links){
  let containerSubMenu = div(
    "absolute group-hover:block hidden z-10 w-40 origin-top-right divide-y divide-gray-100 rounded-md bg-white shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none",
  )
  containerSubMenu.setAttribute("role","menu")
  containerSubMenu.setAttribute("aria-orientation","vertical")
  containerSubMenu.setAttribute("aria-labelledby","menu-button")
  containerSubMenu.setAttribute("tabindex","-1")

  links.setAttribute("role","none")
  containerSubMenu.appendChild(links)

  let container = div(
    "relative group inline-block text-left",
    [
      div(
        "",
        [btn],
      ),
      containerSubMenu,
    ]
  )
  return container
}

//custom button
function CustomButton(imgButton){
  let btn = button(
    "",
    "hover:text-gray-700 text-gray-500 flex items-center text-sm menu",
  )
  btn.setAttribute("aria-expanded","true")
  btn.setAttribute("aria-haspopup","true")
  btn.innerHTML = imgButton
  return btn
}

function createPathRoute(nameRoute){
  let pathRoute = nameRoute.split(" ").join("-")
  return pathRoute
}

function createLinks(names){
  let links = []
  for(let i=0;i<names.length;i++){
    links.push(
      link(
        "menu-item-2",
        "hover:bg-gray-300 text-gray-700 block px-4 py-2 text-sm",
        names[i],
        `#/${createPathRoute(names[i])}`,
      )
    )
  }
  return links
}
//custom links
function wrapLinks(namesLink) {
  let custom = div(
    "",
    createLinks(namesLink),
  )
  return custom
}

export function NavBody(){
  const componentBody = document.createElement("div")
  componentBody.className = "bg-white shadow-md h-16 flex pl-10 gap-7 items-center"

  let buttonAccess = CustomButton(SvgAccess())
  let linkAccess = wrapLinks(["Roles y privilegios","Usuario de sistema"])
  
  let btnMaintenance = CustomButton(SvgMaintenance())
  let linkMaintenance = wrapLinks(["sedes","dependencias"])

  let btnPeople = CustomButton(SvgPeople())
  let linkPeople = wrapLinks(["Datos personales","Empleados"])


  const customLink = document.createElement("a")
  customLink.className = "hover:text-gray-700 text-gray-500 flex items-center text-sm menu"
  customLink.href = `#/${createPathRoute("Equipos informaticos")}`
  customLink.innerHTML = `${SvgComputer()}Equipos informaticos`
  let menuCPU = div(
    "inline-block text-left",
    [
      div(
        "",
        [ customLink ]
      )
    ],
  )
  componentBody.appendChild(CreateMenu(
    buttonAccess,
    linkAccess,
  ))
  componentBody.appendChild(CreateMenu(
    btnMaintenance,
    linkMaintenance,
  ))
  componentBody.appendChild(CreateMenu(
    btnPeople,
    linkPeople,
  ))
  componentBody.appendChild(menuCPU)

  return componentBody
}


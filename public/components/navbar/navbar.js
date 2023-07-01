import button from '../../components/button.js'
import div from '../../components/div.js'
import image from '../../components/image.js'
import span from '../../components/span.js'
import svg from '../../components/svg.js'

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
export default function () {
  const component = document.createElement("div")

  let contentWrap = div(
    "text-sm w-[150px] left-0 text-gray-500 absolute top-[3rem] hidden group-hover:block bg-white shadow-md rounded-md",
    item_wrap,
  )

  let contentCard = div(
    "absolute top-[4rem] right-5 p-2 flex flex-col gap-2",
  )

  component.className = "ml-auto flex items-center py-4 float-right group relative"
  component.appendChild(
    button(
      "",
      "flex items-center text-white",
      items,
      svg(),
    )
  )
  component.appendChild(contentWrap)
  component.appendChild(contentCard)

  return component
}


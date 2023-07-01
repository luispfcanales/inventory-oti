export default function (id,styleClass,text,pathLink = "#/") {
  const item = document.createElement("a")
  item.id = id
  item.className = styleClass
  item.href = pathLink
  item.innerText = text

  item.setAttribute("role","menuitem")
  item.setAttribute("tab-index","-1")
  return item
}

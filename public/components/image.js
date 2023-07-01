export default function (styleClass,src) {
  const divContent = document.createElement("img")
  divContent.className = styleClass
  divContent.src = src

  return divContent
}

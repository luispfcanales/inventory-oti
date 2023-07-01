export default function (styleClass,itemChilds) {
  const divContent = document.createElement("div")
  divContent.className = styleClass

  if (itemChilds !== undefined){
    for(let i=0;i<itemChilds.length;i++){
      divContent.appendChild(itemChilds[i])
    }
  }
  return divContent
}

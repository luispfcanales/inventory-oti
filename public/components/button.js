export default function (idEvent,text,styleClass,itemChilds,svgImage) {
  const item = document.createElement("button")
  item.className = styleClass

  if(idEvent !== undefined ){
    item.id = idEvent
  }

  if(text !== "" ) {
    item.innerText = text
  } 

  if (itemChilds !== undefined){
    for(let i=0;i<itemChilds.length;i++){
      item.appendChild(itemChilds[i])
    }
  }

  if(svgImage !== undefined){
    item.innerHTML += svgImage
  }

  return item
}

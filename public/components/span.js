export default function (idNameUser,text,styleClass,itemChilds) {
  const divContent = document.createElement("span")
  divContent.className = styleClass

  if(idNameUser !== "") {
    divContent.id=idNameUser
  }

  if(text !== "") {
    divContent.innerText = text
  }

  if (itemChilds !== undefined){
    for(let i=0;i<itemChilds.length;i++){
      divContent.appendChild(itemChilds[i])
    }
  }

  return divContent
}

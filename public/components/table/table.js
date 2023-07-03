
//createTH return element of row table 
const createTH =(TEXT,CLASS_NAME)=>{
  let th = document.createElement("th")
  th.className = CLASS_NAME
  th.innerText = TEXT

  return th
}

//createTRhead return element of row table 
const createTRhead =(TH_ELEMENTS)=>{
  const itemTR = document.createElement("tr")
  itemTR.className= "text-gray-500 text-sm"

  for(let i=0;i<TH_ELEMENTS.length;i++){

    let th = createTH(TH_ELEMENTS[i],"px-3 pt-0 pb-3 border-b border-gray-200 dark:border-gray-800")

    itemTR.appendChild(th)
  }

  return itemTR
}
export function tableComputer(KEY_BODY,HEADER_ELEMENTS) {
  const table = document.createElement("table")
  table.className = "w-full text-left overflow-x-auto pb-5"

  const thead = document.createElement("thead")
  thead.className = "bg-gray-100"

  const tbody =document.createElement("tbody")
  tbody.id = KEY_BODY

  thead.appendChild(createTRhead(HEADER_ELEMENTS))
  table.appendChild(thead)
  table.appendChild(tbody)
  return table
}

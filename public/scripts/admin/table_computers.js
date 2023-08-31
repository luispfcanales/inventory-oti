const loadDataComputer=async()=>{

  const url = `${URL_API}/computers?token=123412341234`
  const data = await fetch(url).then((response)=> response.json()).then(data => {return data.data})

  $("#body_table_computer").html(
    data.map(data=>{ 
      return`
        <tr class="hover:bg-gray-100">
          <td class="border-b">${data.key}</td>
          <td class="border-b">${data.architecture}</td>
          <td class="border-b">${data.disk}</td>
          <td class="border-b">${data.maker}</td>
          <td class="border-b">${data.model}</td>
          <td class="border-b">${data.name}</td>
          <td class="border-b">${data.processor}</td>
          <td class="border-b">${data.ram}</td>
          <td class="border-b">${data.serial}</td>
          <td class="border-b">${data.size_disk}</td>
        </tr>
    `})
  )
}

loadDataComputer()

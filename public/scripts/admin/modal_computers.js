if (typeof idmodal === "undefined") {
  const idmodal = document.getElementById("open_modal_computers")
  idmodal.addEventListener("click",()=>{
    document.querySelector("#container_modal_computers").classList.toggle('hidden')
  })
}

if (typeof id_close_modal === "undefined") {
  const id_close_modal = document.getElementById("close_modal_computers")
  id_close_modal.addEventListener("click",()=>{
    document.querySelector("#container_modal_computers").classList.toggle('hidden')
  })
}

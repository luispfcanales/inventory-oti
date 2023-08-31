const open_modal_pc = document.getElementById("open_modal_computers")
open_modal_pc.addEventListener("click",()=>{
  document.querySelector("#container_modal_computers").classList.toggle('hidden')
})
const close_modal_pc = document.getElementById("close_modal_computers")
close_modal_pc.addEventListener("click",()=>{
  document.querySelector("#container_modal_computers").classList.toggle('hidden')
})

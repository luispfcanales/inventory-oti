
export function ButtonModal(ModalID,KeyID) {
  return `
        <div class="flex justify-between mb-4">
          <span class="text-lg font-medium text-gray-700">Roles De Usuarios</span>
          <button id=${KeyID}
          class="bg-[var(--new-button)] py-2 px-4 text-white rounded-md hover:bg-[var(--hover-new-button)]"
            data-modal-target=${ModalID}
            data-modal-toggle=${ModalID}
            type="button"
          >
            +&nbsp;Nuevo
          </button>
        </div>
  `
}

export function ButtonModalElement(MODAL_ID,KEY_ID) {
  const item = document.createElement("div")
  item.className = "flex justify-between mb-4"

  const span = document.createElement("span")
  span.className = "text-lg font-medium text-gray-700"
  span.innerText = "Equipos informaticos"

  const btn = document.createElement("button")
  btn.id = KEY_ID
  btn.className = "bg-[var(--new-button)] py-2 px-4 text-white rounded-md hover:bg-[var(--hover-new-button)]"
  btn.setAttribute("data-modal-target",MODAL_ID)
  btn.setAttribute("data-modal-toggle",MODAL_ID)
  btn.innerHTML = "+&nbsp;Nuevo"

  item.appendChild(span)
  item.appendChild(btn)
  return item
}

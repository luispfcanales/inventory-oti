function modal(ID_STATUS_MODAL,ID_CLOSE_MODAL) {
  return `
    <div
      class="relative w-full h-auto max-w-2xl max-h-full my-2"
      style="z-index: 99999"
    >
      <!-- Modal content -->
      <form
        action="#"
        class="relative bg-white rounded-lg shadow dark:bg-gray-700 h-full"
      >
        <!-- Modal header -->
        <div class="flex items-start justify-between p-4 border-b rounded-t dark:border-gray-600">
          <h3 class="text-xl font-semibold text-gray-900 dark:text-white">Nueva Area</h3>
          <button
            id=${ID_CLOSE_MODAL}
            type="button"
            class="text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm p-1.5 ml-auto inline-flex items-center dark:hover:bg-gray-600 dark:hover:text-white"
            data-modal-hide=${ID_STATUS_MODAL}
          >
            <svg
              aria-hidden="true"
              class="w-5 h-5"
              fill="currentColor"
              viewBox="0 0 20 20"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                fill-rule="evenodd"
                d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
                clip-rule="evenodd"
              ></path>
            </svg>
          </button>
        </div>
        <!-- Modal body -->
        <div class="p-6 space-y-6">
          <div class="flex flex-col gap-3">
            <div class="col-span-6 sm:col-span-3">
              <label
                for="department"
                class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                >Departamento</label
              >
              <select name="" id="select_departments" class="w-full p-2 rounded-md border-gray-300 bg-gray-50">

              </select>
            </div>
            <div class="">
              <label
                for="name"
                class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                >Nombre</label
              >
              <input
                type="text"
                name="first-name"
                id="name_area"
                class="shadow-sm bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-600 focus:border-blue-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                placeholder="Desarrollo"
                required=""
              />
            </div>
            <div class="">
              <label
                for="last"
                class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                >Descripcion</label
              >
              <textarea
                name="last-name"
                id="description_area"
                class="shadow-sm bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-600 focus:border-blue-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                placeholder="Green"
                required=""
              ></textarea>
            </div>
          </div>
        </div>
        <!-- Modal footer -->
        <div class="flex items-center p-6 space-x-2 border-t border-gray-200 rounded-b dark:border-gray-600 justify-end">
          <button
            id="add_area"
            type="button"
            class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
          >
            Guardar
          </button>
        </div>
      </form>
    </div>
  `
}
export function ModalEquipos(ID_STATUS_MODAL,ID_CLOSE_MODAL) {
  const container = document.createElement("div")

  container.id = ID_STATUS_MODAL
  container.className = "hidden fixed top-0 left-0 backdrop-blur-sm right-0 z-50 items-center justify-center w-full p-4 overflow-y-auto md:inset-0 max-h-full h-[calc(100vh)] flex "
  container.style = "background: rgba(0, 0, 0, 0.082)"
  container.setAttribute("tabindex","-1")
  container.setAttribute("aria-hidden","true")

  container.innerHTML=modal(ID_STATUS_MODAL,ID_CLOSE_MODAL)

  return container
}

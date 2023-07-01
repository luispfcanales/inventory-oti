// PAGE ADMIN RENDER
import navbar from '../../components/navbar/navbar.js'

export default function () {
  const divContent = document.createElement("div")
  divContent.className = "h-16 lg:flex w-full border-b border-gray-200 dark:border-gray-800 px-10 sticky top-0 z-50 bg-[var(--header-color)] relative"
  divContent.appendChild(navbar())
  return divContent
}

export function AppendScriptAdmin(name){
  const script = document.createElement('script');
  script.src = `/public/scripts/admin/${name}.js`;
  document.body.appendChild(script);
}

export function AppendScript(name){
  const script = document.createElement('script');
  script.src = `/public/scripts/${name}.js`;
  document.body.appendChild(script);
}

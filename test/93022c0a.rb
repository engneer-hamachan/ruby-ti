a = JS.document

b = a.innerHTML

a.appendChild b.style
a.appendChild 1


a.setAttribute 'a', 'b'
a.setAttribute 'a', b

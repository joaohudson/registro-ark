async function fetchDinos(region = '', locomotion = '', food = '', name = ''){
    const response = await fetch('/api/dinos?region=' + region + '&locomotion=' + locomotion + '&food=' + food + '&name=' + name);

    if(!response.ok){
        throw await response.text();
    }

    return await response.json()
}

async function loadData(){
    let data;
    try{
        data = await fetchDinos();
    }catch(e){
        alert(e);
        return;
    }

    itemsDiv.innerHTML = '';

    const table = document.createElement('table');

    const categoryNames = [
      { name: "name", label: "Nome" },
      { name: "food", label: "Alimentação" },
      { name: "locomotion", label: "Locomoção" },
      { name: "region", label: "Região" }
    ];

    const trTitle = document.createElement("tr");
    for(const f of categoryNames){
        const th = document.createElement('th');
        th.innerText = f.label;
        trTitle.appendChild(th);
    }
    table.appendChild(trTitle);
    
    for(const obj of data){
        const tr = document.createElement('tr');
        tr.onclick = () => {
            location.href = '/dino?id=' + obj.id;
        };
        for(const f of categoryNames){
            const td = document.createElement('td');
            td.innerText = obj[f.name];
            tr.appendChild(td);
        }
        table.appendChild(tr);
    }

    itemsDiv.appendChild(table);
}

loadData();
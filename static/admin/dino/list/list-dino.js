async function fetchDinos(region = '', locomotion = '', food = '', name = ''){
    return await adminFetch('/api/adm/dinos?region=' + region + '&locomotion=' + locomotion + '&food=' + food + '&name=' + name, {}, true);
}

async function deleteDino(id){
    return await adminFetch('/api/dino?id='+id, {method: 'DELETE'});
}

async function loadData(){
    let data;
    try{
        data = await fetchDinos();
    }catch(e){
        dialog.showMessage(e);
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
    const actionsTh = document.createElement('th');
    actionsTh.innerText = 'Ações';
    trTitle.appendChild(actionsTh);
    table.appendChild(trTitle);
    
    for(const obj of data){
        const tr = document.createElement('tr');
        for(const f of categoryNames){
            const td = document.createElement('td');
            td.innerText = obj[f.name];
            tr.appendChild(td);
        }
        table.appendChild(tr);

        const actionsTd = document.createElement('td');
        
        const seeButton = document.createElement('button');
        seeButton.onclick = () => location.href = '/dino?id=' + obj.id;
        seeButton.innerText = '🔍';
        seeButton.title = 'Detalhes do dino';
        seeButton.className = 'actionButton';
        actionsTd.appendChild(seeButton);

        const deleteButton = document.createElement('button');
        deleteButton.onclick = async () => {
            try{
                const ok = await dialog.showConfirm('Tem certeza que dejesa deletar ' + obj.name +'?');
                if(ok)
                    await deleteDino(obj.id);
            }
            catch(e){
                dialog.showMessage(e);
                return;
            }

            loadData();
        };
        deleteButton.innerText = '❌';
        deleteButton.title = 'Remover dino';
        deleteButton.className = 'actionButton';
        actionsTd.appendChild(deleteButton);

        tr.appendChild(actionsTd);
        table.appendChild(tr);
    }

    itemsDiv.appendChild(table);
}

loadData();
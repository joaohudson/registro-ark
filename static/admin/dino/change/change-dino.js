(function(){
    const regionField = document.getElementById('regionField');
    const locomotionField = document.getElementById('locomotionField');
    const foodField = document.getElementById('foodField');
    const utitlityField = document.getElementById('utilityField');
    const trainingField = document.getElementById('trainingField');
    const nameField = document.getElementById('nameField');
    const createDinoButton = document.getElementById('createDinoButton');
    
    async function fetchCategories(){
        const response = await fetch('/api/dino/categories');
    
        if(!response.ok){
            throw await response.text();
        }
    
        return await response.json();
    }
    
    async function fetchDino(id){
        const response = await fetch('/api/dino?id=' + id);
    
        if(!response.ok){
            throw await response.text();
        }
    
        return await response.json();
    }
    
    async function putDino(dino, id){
        return adminFetch('/api/dino?id='+id, {method: 'PUT', body: JSON.stringify(dino)});
    }
    
    async function changeDino(id){
        const dino = {
            name: nameField.value,
            utility: utitlityField.value,
            training: trainingField.value,
            regionId: new Number(regionField.value),
            locomotionId: new Number(locomotionField.value),
            foodId: new Number(foodField.value)
        };
    
        try{
            await putDino(dino, id);
        }
        catch(e){
            console.log(e);
            dialog.showMessage(e);
            return;
        }
    
        dialog.showMessage('Dino alterado com sucesso!');
    }
    
    function pupulateDropdown(data, dropdown){
        const keys = data.map(k => k.name);
        const values = data.map(v => v.id);
        const optionsKeys =  keys;
        const optionsValues = values;
    
        for(let i = 0; i < optionsKeys.length; i++){
            const key = optionsKeys[i];
            const value = optionsValues[i];
            const opt = document.createElement('option');
            opt.innerText = key;
            opt.value = value;
            dropdown.appendChild(opt);
        }
    }
    
    function setDropdownValue(dropdown, value){
        const {options} = dropdown;
        for(const opt of options){
            opt.selected = opt.value == value;
        }
    }
    
    async function main(){
        let categories;
        let dino;
        let id;
        const params = new URLSearchParams(location.search);
    
        try{
            if(!params.has('id')){
                throw 'Houve um erro ao carregar a página, volte para a página anterior!';
            }
            id = params.get('id');
            categories = await fetchCategories();
            dino = await fetchDino(id);
        }
        catch(e){
            dialog.showMessage(e);
            return;
        }
    
        //carregando dropdrowns
        pupulateDropdown(categories.regions, regionField);
        pupulateDropdown(categories.locomotions, locomotionField);
        pupulateDropdown(categories.foods, foodField);
        
        //carregando parâmetros
        setDropdownValue(regionField, dino.region.id);
        setDropdownValue(locomotionField, dino.locomotion.id);
        setDropdownValue(foodField, dino.food.id);
        nameField.value = dino.name;
        utitlityField.value = dino.utility;
        trainingField.value = dino.training;
    
        createDinoButton.onclick = async () => await changeDino(id);
    }
    
    main();
    
})();
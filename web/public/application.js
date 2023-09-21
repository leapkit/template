function openConfirm(el){
    const title = el.dataset.title
    const description = el.dataset.description
    const deleteText = el.dataset.delete
    const cancelText = el.dataset.cancel

    return Swal.fire({
        title: title, 
        text: description, 
        showCancelButton: true, 
        showCloseButton: true, 
        customClass:{
                container: 'top-0 left-0 right-0 bottom-0', 
                title: 'text-lg font-semibold leading-6 text-gray-900 mb-1 inline-block flex', 
                cancelButton: 'bg-white mt-3 w-full rounded-md border border-gray-300 shadow-sm px-4 py-2 text-base font-medium focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:mt-0 sm:w-auto sm:text-sm', 
                actions: 'container-actions justify-end w-full rounded mt-7',  
                confirmButton: `bg-red-600 w-full rounded-md border border-transparent shadow-sm px-4 py-2 text-base font-medium text-white focus:outline-none sm:ml-3 sm:w-auto sm:text-sm`
            }, 
        buttonsStyling: false,
        reverseButtons: true, 
        confirmButtonText: deleteText, 
        cancelButtonText: cancelText
    })
}

function selectFile(e) {
    const fileInput = e.target;
    const file = fileInput.files[0];

    if (file) {
        let imageValue = e.target.files[0]
        document.getElementById('selectFile').classList.add('hidden')
        document.getElementById('fileSelected').classList.remove('hidden')
        document.getElementById('fileName').innerHTML = imageValue.name
        document.getElementById('action').value = "create"
    }
}

function clearSelection() {
    document.getElementById('selectFile').classList.remove('hidden')
    document.getElementById('fileSelected').classList.add('hidden')
    document.getElementById('document-file-upload').value = ''
    document.getElementById('fileName').innerHTML = ""
    document.getElementById('action').value = "delete"
}

function profile(e) {
    const fileInput = e.target;
    const file = fileInput.files[0];

    if (file) {
        let errorFileImage = document.getElementById("errorFile")
        if (file.size > 5 * 1024 * 1024) {
            errorFileImage.innerHTML = "The image must be less than 5MB in size"
            fileInput.value = '';
            return
        }
        let imageValue = e.target.files[0]
        errorFileImage.innerHTML = ""
        document.getElementById('imageProfile').src = this.src(imageValue)
        document.getElementById('selectImage').classList.add('hidden')
        document.getElementById('imageSelected').classList.remove('hidden')
        document.getElementById('imageSelected').classList.add('flex')
        document.getElementById('fileName').innerHTML = imageValue.name
        document.getElementById('action').value = "create"
    }
}

function clear() {
    document.getElementById('selectImage').classList.remove('hidden')
    document.getElementById('imageSelected').classList.add('hidden')
    document.getElementById('document-file-upload').value = ''
    document.getElementById('imageProfile').src = ""
    document.getElementById('fileName').innerHTML = ""
    document.getElementById('action').value = "delete"
}

function src(file){
    let regex = new RegExp('[^.]+$');
    let ext = regex.exec(file.name)[0]

    const src = (ext == 'svg' || ext == 'png' || ext == 'jpg' || ext == 'jpeg') ?  URL.createObjectURL(file):ext
    return src
}

function selectEmployee(e){
    document.getElementById('empSelected').value = e.dataset.id;
    let paintEmployee = document.getElementById('employeeSelected');
    paintEmployee.innerHTML = e.innerHTML
    let inputSelected = document.getElementById('employeeSelected')
    inputSelected.parentNode.parentNode.classList.add('flex')
    inputSelected.parentNode.parentNode.classList.remove('hidden')
}

function resetValue(){
    document.getElementById('empSelected').value = '';
    let paintEmployee = document.getElementById('employeeSelected');
    paintEmployee.innerHTML = '';
    let inputSelected = document.getElementById('employeeSelected')
    inputSelected.parentNode.parentNode.classList.add('hidden')
    document.getElementById('search-div').classList.remove('hidden')
}
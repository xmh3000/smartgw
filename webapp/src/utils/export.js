export const exportFile = (data, type, res, name = "") => {
  let blob = new Blob([data], {
    type: type,
  });

  let downloadElement = document.createElement("a");
  let href = window.URL.createObjectURL(blob); //创建下载的链接
  downloadElement.href = href;
  downloadElement.download =
    decodeURI(res.headers["content-disposition"].split("filename=")[1]) || name; //下载后文件名
  document.body.appendChild(downloadElement);
  downloadElement.click(); //点击下载
  document.body.removeChild(downloadElement); //下载完成移除元素
  window.URL.revokeObjectURL(href); //释放掉blob对象
};
export const exportCsv = (data, name)=>{
  // “\ufeff” BOM头
  var uri = 'data:text/csv;charset=utf-8,\ufeff' + encodeURIComponent(data);
  var downloadLink = document.createElement("a");
  downloadLink.href = uri;
  downloadLink.download = (name+".csv")||"tsl.csv";
  document.body.appendChild(downloadLink);
  downloadLink.click();
  document.body.removeChild(downloadLink);
}

export const exportJson = (data, name)=>{
  // “\ufeff” BOM头
  var uri = 'data:text/json;charset=utf-8,' + encodeURIComponent(data);
  var downloadLink = document.createElement("a");
  downloadLink.href = uri;
  downloadLink.download = (name+".json")||"devices.json";
  document.body.appendChild(downloadLink);
  downloadLink.click();
  document.body.removeChild(downloadLink);
}

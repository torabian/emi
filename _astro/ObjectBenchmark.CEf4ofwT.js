import{r as j}from"./index.CO9X3OiW.js";var m={exports:{}},l={};var d;function f(){if(d)return l;d=1;var i=Symbol.for("react.transitional.element"),t=Symbol.for("react.fragment");function r(c,e,s){var o=null;if(s!==void 0&&(o=""+s),e.key!==void 0&&(o=""+e.key),"key"in e){s={};for(var u in e)u!=="key"&&(s[u]=e[u])}else s=e;return e=s.ref,{$$typeof:i,type:c,key:o,ref:e!==void 0?e:null,props:s}}return l.Fragment=t,l.jsx=r,l.jsxs=r,l}var p;function g(){return p||(p=1,m.exports=f()),m.exports}var a=g();class N{str="";num=0;obj={};arr=[];setStr(t){this.str=typeof t=="string"?t:String(t)}getStr(){return this.str}setNum(t){if(typeof t=="number")this.num=t;else{const r=Number(t);this.num=isNaN(r)?0:r}}getNum(){return this.num}setObj(t){this.obj=typeof t=="object"&&!Array.isArray(t)?t:{value:t}}getObj(){return this.obj}setArr(t){this.arr=Array.isArray(t)?t:[t]}getArr(){return this.arr}}function v(){const[i,t]=j.useState(1e5),[r,c]=j.useState(null),e=()=>{let o=performance.now(),u=performance.now();const b=u-o,n=new N;o=performance.now();for(let x=0;x<i;x++)n.setStr("Hello"),n.setNum(42),n.setObj({k:"v"}),n.setArr([1,2,3]),n.getStr(),n.getNum(),n.getObj(),n.getArr();u=performance.now();const h=u-o;c({simple:b,classBased:h})},s=`// simple object
const obj: any = {};
for (let i = 0; i < ${i}; i++) {
  obj.str = "Hello";
  obj.num = 42;
  obj.obj = { k: "v" };
  obj.arr = [1, 2, 3];
  const x1 = obj.str;
  const x2 = obj.num;
  const x3 = obj.obj;
  const x4 = obj.arr;
}

// class object
const o = new TestClass();
for (let i = 0; i < ${i}; i++) {
  o.setStr("Hello");
  o.setNum(42);
  o.setObj({ k: "v" });
  o.setArr([1, 2, 3]);
  const x1 = o.getStr();
  const x2 = o.getNum();
  const x3 = o.getObj();
  const x4 = o.getArr();
}`;return a.jsxs("div",{className:"p-4 space-y-4",children:[a.jsxs("div",{className:"flex gap-2 items-center",children:[a.jsx("input",{type:"number",className:"border p-2 rounded",value:i,onChange:o=>t(Number(o.target.value))}),a.jsx("button",{onClick:e,className:"px-4 py-2 bg-blue-600 text-white rounded",children:"Run"})]}),r&&a.jsxs("div",{className:"space-y-2",children:[a.jsxs("div",{children:["Object property: ",r.simple.toFixed(2)," ms"]}),a.jsxs("div",{children:["Class get/set: ",r.classBased.toFixed(2)," ms"]})]}),a.jsx("pre",{className:"bg-gray-100 p-3 rounded overflow-x-auto text-sm",children:s})]})}export{v as default};

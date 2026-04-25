[如何使用nvm切换Node.js和npm版本-CSDN博客](https://blog.csdn.net/weixin_43031220/article/details/130536030)

一个将数据渲染为html视图的开源JavaScript库。就是js代码所有适用于js的在这里都适用！

**npx create-react-app my-app（项目名）：创建React项目**



npx create-next-app@latest：Next.js 的页面路由，是一个全栈的 React 框架，添加了服务器端渲染（SSR）和静态网站生成（SSG）的能力。

npx create-remix：Remix是一个具有嵌套路由的全栈式 React 框架

npx create-gatsby：Gatsby是一个快速的支持 CMS 的网站的 React 框架

npx create-expo-app：Expo是一个 React 框架，可以让你创建具有真正原生 UI 的应用，包括 Android、iOS，以及 Web 应用



在项目文件夹下执行npm init（npm init -y 全部接受）可以初始化项目为npm管理的项目。



react应用：在创建好的react项目里，执行npm install react react-dom，在文件里面 import { createRoot } from 'react-dom/client;就可以写代码了

使用vscode打开某文件夹 ：进入这个文件夹运行 coed  ./ 实际上就是在code后面加文件路径。



bable：ES6转ES5、jsx转js

**js基础知识：**

判断this指向

class类

ES6语法规范：没事就看看：[export - JavaScript | MDN (mozilla.org)](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Statements/export)熟悉一下ES6

npm包管理器

原型、原型链：[彻底搞懂JavaScript原型和原型链-CSDN博客](https://blog.csdn.net/gaotlantis/article/details/139152875)

数组常用方法

模块化



定义虚拟DOM不用单引号，在标签使用js时要用{}

样式的类名指定不要用class，要用className：\<h1  className='title'……>……    .title{样式}

内联的style要写style={{ 样式……key:value}}

只有一个根标签

标签首字母：1、如果以小写字母开头，就转化为对应的html标签，如果没有对应的html标签就报错；2、React 组件必须以大写字母开头，而 HTML 标签则必须是小写字母。以大写字母开头的，react就去渲染对应的组件，如果没有对应的组件，就报错。



区分js表达式（一个表达式会产生一个值，可以放在如何需要值的地方）和js代码块。所以给数组或其他东西就要叫用函数及将其加工成字符串。



可以使用 `className` 来指定一个 CSS 的 class。它与 HTML 的 `class`( 属性的工作方式相同



react路由：

[React Router 使用教程 - 阮一峰的网络日志 (ruanyifeng.com)](https://ruanyifeng.com/blog/2016/05/react_router.html)

[ReactRouter——路由配置、路由跳转、带参跳转、新route配置项_react-router-dom 6.x 路由调跳转-CSDN博客](https://blog.csdn.net/DogEgg_001/article/details/139449907)



渲染列表：

假设你有一个产品数组：

```jsx
const products = [
  { title: 'Cabbage', id: 1 },
  { title: 'Garlic', id: 2 },
  { title: 'Apple', id: 3 },
];
```

在你的组件中，使用 `map()` 函数将这个数组转换为 `<li>` 标签构成的列表:

```jsx
const listItems = products.map(product =>
  <li key={product.id}>
    {product.title}
  </li>
);

return (
  <ul>{listItems}</ul>
);
```

注意， `<li>` 有一个 `key` 属性。对于列表中的每一个元素，你都应该传递一个字符串或者数字给 `key`，用于在其兄弟节点中唯一标识该元素。通常 key 来自你的数据，比如数据库中的 ID。如果你在后续插入、删除或重新排序这些项目，React 将依靠你提供的 key 来思考发生了什么。



响应事件：

你可以通过在组件中声明 **事件处理** 函数来响应事件：

```jsx
function MyButton() {
  function handleClick() {
    alert('You clicked me!');
  }

  return (
    <button onClick={handleClick}>
      Click me
    </button>
  );
}
```

注意，`onClick={handleClick}` 的结尾没有小括号！加括号是调用；你只需 **把函数传递给事件** 绑定即可。当用户点击按钮时 React 会调用你传递的事件处理函数



React 提供了一个名为 `useState` 的特殊函数，可以从组件中调用它来让它“记住”一些东西import { useState } from 'react';的const [count, setCount] = useState(0);可以记录状态，这个东西可以放在任何一个代码里面，也可以作为参数传递

```javascript
export default function MyApp() {
  return (
    <div>
      <h1>Counters that update together</h1>
      <MyButton count={count} onClick={handleClick} /> // 参数这样传递。
      <MyButton count={count} onClick={handleClick} />
    </div>
  );
}

function MyButton({ count, onClick }) {
  return (
    <button onClick={onClick}>
      Clicked {count} times
    </button>
  );
}
```



const [enabled, setEnabled] = useState(false);// 会重用作为初始 state 传入的值以确定值的类型，推断类型为 "boolean"

// 显式设置类型为 "boolean"

const [enabled, setEnabled] = useState\<boolean>(false);



使用 Chrome 浏览器，则可以使用键盘快捷键 **Shift + Ctrl + J**（在 Windows/Linux 上）或 **Option + ⌘ + J**（在 macOS 上）查看控制台，ctrl+L：清空控制台。

------



react项目的标准结构：

```bash
your-react-project/
├── node_modules/       # 项目依赖包（自动生成，不用管）
├── public/             # 静态资源,这里的资源不能用import引入，用 /路径访问
├── src/                # 核心源码（99% 开发都在这里）
│   ├── assets/         # 需要打包的静态资源：图片、字体、全局样式、图标。可以 import 引入
│   ├── components/     # 公共可复用组件（Button、Card、Modal等）通用、无业务逻辑、可复用
│   ├── layouts/        # 页面骨架（Header、Footer）避免每个页面重复写导航侧边栏。
│   ├── pages/          # 页面组件，一个文件对应一个页面（Home、About、Login）
│   ├── router/         # 路由配置（React Router 路由表）
│   ├── store/          # 全局状态管理（Redux/Zustand/Jotai/Context）跨页面共享数据，用户信息、主题、全局弹窗、购物车等
│   ├── hooks/          # 自定义 Hooks（useFetch、useUser、useModal等）让组件代码更干净，逻辑可复用
│   ├── utils/          # 工具函数（时间格式化、请求封装、校验函数、防抖）
│   ├── services/       # API 请求接口（axios 封装、接口统一管理）
│   ├── constants/      # 常量定义（枚举、配置项、固定文本）
│   ├── types/          # TS 类型定义（只在 TypeScript 项目中）
│   ├── App.jsx         # 根组件（路由入口）
│   └── main.jsx        # 项目入口文件（渲染 App 到 DOM）
├── .env                # 环境变量（开发环境）
├── .env.production     # 环境变量（生产环境）
├── .gitignore          # Git 忽略文件
├── package.json        # 项目配置、依赖、脚本命令
├── vite.config.js      # Vite 配置文件（CRA 则是 craco.config.js）
└── README.md           # 项目说明文档
```





## 组件的导入导出

```jsx
默认导出和具名导出export default和export，它们之间的主要区别在于如何导入它们。

export default function App():

同一文件中，有且仅有一个默认导出，但可以有多个具名导出！
当你导入时，你可以给这个导出命名为任何名称，而不需要使用花括号。
示例：import App from './App';//这个就是默认导入，导入他export default的组件

export function App():
这允许你导出一个或多个命名导出。一个模块可以包含多个命名导出。
当你导入命名导出时，你需要使用花括号并且必须使用相同的名称（除非使用as关键字重命名）。
这对于导出和导入多个函数或变量很有用。
示例导入：import { App as myapp } from './App';//这个就是具名导入，可以同名也可以重命名。
```





## JSX语法

JSX 是一种语法扩展，而 React 则是一个 JavaScript 的库。

### return规范

在组件retuen时，如果只返回一个东西，那么要紧跟在return后面。

return \<img src="\https:\//i.imgur.com/jA8hHMpm.jpg" alt="Katsuko Saruhashi" />;

要不然就用括号括起来

 return (
    \<img 
      src="https:\//i.imgur.com/jA8hHMpm.jpg" 
      alt="Katsuko Saruhashi" 
    />
  );

### {}插值表达式

1. jsx只能返回一个根元素，如果要返回多个可以使用div或者\<></>来将他们包裹起来。因为：JSX 虽然看起来很像 HTML，但在底层其实被转化为了 JavaScript 对象，你不能在一个函数中返回多个对象，除非用一个数组把他们包装起来。这就是为什么多个 JSX 标签必须要用一个父元素或者 Fragment 来包裹
2. 标签必须闭合。
3. 使用驼峰式命名法给大部分属性命名，因为return的标签会被转化为js对象，里面一些关键字像class是不能出现的。属性参考https://zh-hans.react.dev/reference/react-dom/components/common
4. 在jsx里面可以使用{}来写js代码，但标签内的h1……这些不能使用{}，大括号让你可以将 JavaScript 的逻辑和变量带入到标签中。
5. 出现双大括号是最外面是插值表达式，里面的大括号是js的对象！只是在插值表达式里写了一个对象。在使用内联样式是经常style={{样式}}。

### props

Props 是只读的时间快照：每次渲染都会收到新版本的 props。

像img标签的src、alt这些就是props，我们自定义的组件也可以自定义props（就是函数的参数而已）：

```jsx
//React 组件函数接受一个参数，一个 props 对象
function Avatar({ person, size=0 }) {
    //如果你想在没有指定值的情况下给 prop 一个默认值，你可以通过在参数后面写 = 和默认值来进行解构，如果你想在没有指定值的情况下给 prop 一个默认值，你可以通过在参数后面写 = 和默认值来进行解构：
  // 在这里 person 和 size 是可访问的
}
//在声明 props 时， 不要忘记 ( 和 ) 之间的一对花括号 { 和 }  ，这种语法被称为 “解构”，等价于于从函数参数中读取属性

export default function Profile() {
  return (
    <Avatar
      person={{ name: 'Lin Lanying', imageId: '1bX5QH6' }}
      size={100}
    />
  );
}
//——————————————————————————————————————————————————————
//可以使用序列解包的方法，其中props是一个对象，可以使用 <Avatar {...props} /> JSX 展开语法转发所有 props，但不要过度使用它！	
function Profile(props) {
  return (
    <div className="card">
      <Avatar {...props} />
    </div>
  );
}
```

在自己的组件里嵌套组件：

childern是一个特殊的属性，用于接受<>children\</>里面的内容！

```jsx
import Avatar from './Avatar.js';

function Card({ children }) {
  return (
    <div className="card">
      {children}
    </div>
  );
}

export default function Profile() {
  return (
    <Card>
      <Avatar
        size={100}
        person={{ 
          name: 'Katsuko Saruhashi',
          imageId: 'YfeOqp2'
        }}
      />
    </Card>
  );
}
//——————————————————————————————————————————————
//props的参数可以来自内敛的和children
function Card({ children, title }) {
  return (
    <div className="card">
      <div className="card-content">
        <h1>{title}</h1>
        {children}
      </div>
    </div>
  );
}

export default function Profile() {
  return (
    <div>
      <Card title="About">
        <p>Aklilu Lemma was a distinguished Ethiopian scientist who discovered a natural treatment to schistosomiasis.</p>
      </Card>
    </div>
  );
}
```

### 条件渲染

使用js的`if` 语句、`&&` 和 `? :` 运算符来选择性地渲染 JSX。

- 在 JSX 中，`{cond ? <A /> : <B />}` 表示 *“当 `cond` 为真值时, 渲染 `<A />`，否则 `<B />`”*。
- 在 JSX 中，`{cond && <A />}` 表示 *“当 `cond` 为真值时, 渲染 `<A />`，否则不进行渲染”*

```jsx
//if比较直接
function Item({ name, isPacked }) {
  if (isPacked) {
    return <li className="item">{name} ✔</li>;
  }
  return <li className="item">{name}</li>;
}
//使用三目运算符
return (
  <li className="item">
    {isPacked ? name + ' ✔' : name}
  </li>
);
//使用&&运算符，当 isPacked 为真值时，则（&&）渲染勾选符号，否则，不渲染。
function Item({ name, isPacked }) {
  return (
    <li className="item">
      {name} {isPacked && '✔'}//可以看成‘√’是true，所以结果取决于前面的布尔值。
    </li>
  );
/*
切勿将数字放在 && 左侧.
JavaScript 会自动将左侧的值转换成布尔类型以判断条件成立与否。然而，如果左侧是 0，整个表达式将变成左侧的值（0），React 此时则会渲染 0 而不是不进行渲染。
一个常见的错误是 messageCount && <p>New messages</p>。其原本是想当 messageCount 为 0 的时候不进行渲染，但实际上却渲染了 0。
为了更正，可以将左侧的值改成布尔类型：messageCount > 0 && <p>New messages</p>。
*/
```

### 渲染列表

使用 `filter()`筛选需要渲染的组件和使用 `map()`把数组转换成组件数组。注意：!==的使用。

```jsx
const people = [{
  id: 0,
  name: '凯瑟琳·约翰逊',
  profession: '数学家',
}, 
……
                ]

//使用filter选出化学家
const chemists = people.filter(person =>
  person.profession === '化学家'
);//箭头函数会隐式地返回位于 => 之后的表达式，所以你可以省略 return 语句。

const listItems = chemists.map(person =>
  <li>
     <img
       src={getImageUrl(person)}
       alt={person.name}
     />
     <p>
       <b>{person.name}:</b>
       {' ' + person.profession + ' '}
       因{person.accomplishment}而闻名世界
     </p>
  </li>
);
return <ul>{listItems}</ul>;


//——————————————————————————块函数体————————————————————————————
如果你的 => 后面跟了一对花括号 { ，那你必须使用 return 来指定返回值！

const listItems = chemists.map(person => { // 花括号
  return <li>...</li>;
});
箭头函数 => { 后面的部分被称为 “块函数体”，块函数体支持多行代码的写法，但要用 return 语句才能指定返回值。假如你忘了写 return，那这个函数什么都不会返回！


//————————————————还可以直接写在插值表达式里面
export default function RecipeList() {
  return (
    <div>
      <h1>菜谱</h1>
          //这里是第一个插值表达式
      {recipes.map(recipe =>
        <div key={recipe.id}>
          <h2>{recipe.name}</h2>
          <ul>
              //这里是第二个插值表达式
            {recipe.ingredients.map(ingredient =>
              <li key={ingredient}>
                {ingredient}
              </li>
            )}
          </ul>
        </div>
      )}
    </div>
  );
}

```

一个精心选择的 key 值所能提供的信息远远不止于这个元素在数组中的位置。即使元素的位置在渲染的过程中发生了改变，它提供的 `key` 值也能让 React 在整个生命周期中一直认得它。



### 保持组件存粹

在 React 中，你可以在渲染时读取三种输入：props，state 和 context。你应该始终将这些输入视为只读。

当你想要更改数组的任意项时，必须先对其进行拷贝。

记住数组上的哪些操作会修改原始数组、哪些不会，这非常有帮助。例如，`push`、`pop`、`reverse` 和 `sort` 会改变原始数组，但 `slice`、`filter` 和 `map` 则会创建一个新数组。

```jsx
//这是一段有bug的代码，功能不纯粹，因为他修改了传进来的props，导致props一直更新，所以每次时钟信号到来的时候就调用这个组件渲染，然后又往里面添加，无限循环，导致页面出现无尽的添加的东西
export default function StoryTray({ stories }) {
  stories.push({
    id: 'create',
    label: 'Create Story'
  });

  return (
    <ul>
      {stories.map(story => (
        <li key={story.id}>
          {story.label}
        </li>
      ))}
    </ul>
  );
}
//解决办法是：
export default function StoryTray({ stories }) {
  // 复制数组！
  let storiesToDisplay = stories.slice();

  // 不影响原始数组：
  storiesToDisplay.push({
    id: 'create',
    label: 'Create Story'
  });

  return (
    <ul>
      {storiesToDisplay.map(story => (
        <li key={story.id}>
          {story.label}
        </li>
      ))}
    </ul>
  );
}
```

### 绑定事件

- 通常在你的组件 **内部** 定义。
- 名称以 `handle` 开头，后跟事件名称。

```jsx
export default function Button() {
  function handleClick() {
    alert('你点击了我！');
  }

  return (
    <button onClick={handleClick}>
      点我
    </button>
  );
}

//JSX 中定义一个内联的事件处理函数：
<button onClick={function handleClick() {
  alert('你点击了我！');
}}>
    
//lambda表达式
<button onClick={() => {
  alert('你点击了我！');
}}>    
```

```
<button onClick={handleClick}>//是传递一个函数
<button onClick={handleClick()}>//是调用一个函数
在第一个示例中，handleClick 函数作为 onClick 事件处理函数传递。这会让 React 记住它，并且只在用户点击按钮时调用你的函数。
在第二个示例中，handleClick() 中最后的 () 会在 渲染 过程中 立即 触发函数，即使没有任何点击。这是因为在 JSX { 和 } 之间的 JavaScript 会立即执行。

<button onClick={() => alert('...')}>//这个也是传递一个函数
<button onClick={alert('...')}>//这个就是调用函数了，会在渲染时立即执行
```

#### 将事件事件作为props传递

事件处理函数 props 应该以 `on` 开头，后跟一个大写字母。

```jsx
function Button({ onClick, children }) {
  return (
    <button onClick={onClick}>
      {children}
    </button>
  );
}

function PlayButton({ movieName }) {
  function handlePlayClick() {
    alert(`正在播放 ${movieName}！`);
  }

  return (
    <Button onClick={handlePlayClick}>
      播放 "{movieName}"
    </Button>
  );
}

function UploadButton() {
  return (
    <Button onClick={() => alert('正在上传！')}>
      上传图片
    </Button>
  );
}

export default function Toolbar() {}
```

阻值事件冒泡：阻止触发绑定在外层标签上的事件处理函数

```jsx
function Button({ onClick, children }) {
  return (
    <button onClick={e => {
      e.stopPropagation();//这段代码就是阻值事件冒泡的代码
      onClick();
    }}>
      {children}
    </button>
  );
}

export default function Toolbar() {
  return (
    <div className="Toolbar" onClick={() => {
      alert('你点击了 toolbar ！');
    }}>
      <Button onClick={() => alert('正在播放！')}>
        播放电影
      </Button>
      <Button onClick={() => alert('正在上传！')}>
        上传图片
      </Button>
    </div>
  );
}

```

阻值浏览器默认行为：

```jsx
export default function Signup() {
  return (
    <form onSubmit={e => {
      e.preventDefault();//这里就会阻值浏览器的默认行为，阻值表单提交。
      alert('提交表单！');
    }}>
      <input />
      <button>发送</button>
    </form>
  );
}
```

通过bodystyle选中调用函数的组件：

```jsx
export default function LightSwitch() {
  function handleClick() {
    let bodyStyle = document.body.style;
    if (bodyStyle.backgroundColor === 'black') {
      bodyStyle.backgroundColor = 'white';
    } else {
      bodyStyle.backgroundColor = 'black';
    }
  }

  return (
    <button onClick={handleClick}>
      切换背景
    </button>
  );
}
```

### state组件的记忆

```jsx
像这样渲染是不行的，因为index是局部变量，当触发函数时会重新渲染这个组件，index还是从0开始。
export default function Gallery() {
  let index = 0;

  function handleClick() {
    index = index + 1;
  }

  let sculpture = sculptureList[index];
  return (
    <>
      <button onClick={handleClick}>
        Next
      </button>
      <h2>
        <i>{sculpture.name} </i> 
```

import { useState } from 'react';

1. **state 变量** (`index`) 会保存上次渲染的值。
2. **state setter 函数** (`setIndex`) 可以更新 state 变量并触发 React 重新渲染组件。

将let  index=0；改为const [index, setIndex] = useState(0);

其中`index` 是一个 state 变量，`setIndex` 是对应的 setter 函数，这里的 `[` 和 `]` 语法称为[数组解构](https://zh.javascript.info/destructuring-assignment)，它允许你从数组中读取值。 `useState` 返回的数组总是正好有两项。

下面是修改后的代码：

```sx
import { useState } from 'react';	

export default function Gallery() {
  const [index, setIndex] = useState(0);//这个例子里，你希望 React 记住 index。useState 的唯一参数是 state 变量的初始值。在这个例子中，index 的初始值被useState(0)设置为 0

  function handleClick() {
    setIndex(index + 1);//通过调用setter触发函数来对index进行修改
  }

  let sculpture = sculptureList[index];
  return (
    <>
      <button onClick={handleClick}>
        Next
      </button>
      <h2>
        <i>{sculpture.name} </i> 
```



在 React 中，`useState` 以及任何其他以“`use`”开头的函数都被称为 **Hook**。Hook 是特殊的函数，只在 React 渲染时有效。**Hooks ——以 `use` 开头的函数——只能在组件或自定义 Hook的最顶层调用**

应用启动时会第一次渲染。

- **对于初次渲染，** React 会使用` appendChild()`DOM API 将其创建的所有 DOM 节点放在屏幕上。
- **对于重渲染，** React 将应用最少的必要操作（在渲染时计算！），以使得 DOM 与最新的渲染输出相互匹配。

**React 仅在渲染之间存在差异时才会更改 DOM 节点。** 例如，有一个组件，它每秒使用从父组件传递下来的不同属性重新渲染一次。注意，你可以添加一些文本到 `<input>` 标签，更新它的 `value`，但是文本不会在组件重渲染时消失

#### state快照：

```jsx
import { useState } from 'react';

export default function Counter() {
  const [number, setNumber] = useState(0);
  return (
    <>
      <h1>{number}</h1>
      <button onClick={() => {
        setNumber(number + 5);
        alert(number);
      }}>+5</button>
    </>
  )
}
//这段代码的运行结果是提示0
因为运行这段代码是state（number）的值是0，里面alert的number也是这个值，因为这时候函数还没运行完，state（number）的值还没有提交，在没提交前他读到的是没提交前的值。
将alert异步延时后就看见提交后的了
 setTimeout(() => {
          alert(number);
        }, 3000);
```



**调用state的setXxxx才会触发组件的重新渲染，如果是直接修改state的值是不能重新渲染的。而且提交时间是在state所在的组件运行完**



### 对state对象的修改

#### 常规修改

当有多个值要渲染时，我们调用setXxxx时还要复制原来的值到这个对象中，再将修改的值覆盖，可以使用对象解包：<font  color=blue>`...` 展开语法本质是是“浅拷贝”——它只会复制一层。这使得它的执行速度很快，但是也意味着当你想要更新一个嵌套属性时，你必须得多次使用展开语法。</font>

```jsx
setPerson({
  ...person, // 复制上一个 person 中的所有字段
  firstName: e.target.value // 但是覆盖 firstName 字段 
});//对于大型表单，将所有数据都存放在同一个对象中是非常方便的

//代码使用
  function handleEmailChange(e) {
    setPerson({
      ...person,
      email: e.target.value
    });
  }
```

动态修改字段：使用 `[` 和 `]` 括号来实现属性的动态命名，实现一个事件处理函数处理多个事件。

```jsx
export default function Form() {
  const [person, setPerson] = useState({
    firstName: 'Barbara',
    lastName: 'Hepworth',
    email: 'bhepworth@sculpture.com'
  });

  function handleChange(e) {
    setPerson({
      ...person,
      [e.target.name]: e.target.value//这里的e.target就是指调用这个函数的标签，.name就是他的peops的值，所以这个[e.target.name]解析出来就可以对应person对象的成员。
    });
  }

  return (
    <>
      <label>
        First name:
        <input
          name="firstName"
          value={person.firstName}
          onChange={handleChange}
        />
      </label>
      <label>
        Last name:
        <input
          name="lastName"
          value={person.lastName}
          onChange={handleChange}
        />
      </label>
      <label>
        Email:
        <input
          name="email"
          value={person.email}
          onChange={handleChange}
        />
      </label>
      <p>
        {person.firstName}{' '}
        {person.lastName}{' '}
        ({person.email})
      </p>
    </>
  );
}

```

<font color=red>在JavaScript中，`xxx: yyy`这种格式通常指的是对象字面量（object literal）中的键值对，其中`xxx`是键（key），`yyy`是值（value）。</font>

嵌套对象修改成员的方法：

```jsx
const nextArtwork = { ...person.artwork, city: 'New Delhi' };//将嵌套变量单独复制出来修改
const nextPerson = { ...person, artwork: nextArtwork };//修改完后把覆盖原来的
```



#### 使用 Immer:

由 Immer 提供的 `draft` 是一种特殊类型的对象，被称为 Proxy，它会记录你用它所进行的操作。这就是你能够随心所欲地直接修改对象的原因所在！从原理上说，Immer 会弄清楚 `draft` 对象的哪些部分被改变了，并会依照你的修改创建出一个全新的对象。

1. 运行 `npm install use-immer` 添加 Immer 依赖
2. 用 `import { useImmer } from 'use-immer'` 替换掉 `import { useState } from 'react'`

```jsx
import { useImmer } from 'use-immer';

export default function Form() {
  const [person, updatePerson] = useImmer({
    name: 'Niki de Saint Phalle',
    artwork: {
      title: 'Blue Nana',
      city: 'Hamburg',
      image: 'https://i.imgur.com/Sd1AgUOm.jpg',
    }
  });
// updatePerson的draft会自动绑定对应的state
  function handleNameChange(e) {
    updatePerson(draft => {
      draft.name = e.target.value;
    });
  }

……

  function handleImageChange(e) {
    updatePerson(draft => {
      draft.artwork.image = e.target.value;
    });
  }
```

注意：

```jsx
  function handleMove(dx, dy) {
    shape.position.x += dx;
    shape.position.y += dy;
  }
//这个是直接修改了类里面的值，如果这是一个外部的引用，他就会触发mutation，要是其他组件使用他就会出现意想不到的结果！！
  function handleMove(dx, dy) {
    setShape({
      ...shape,
      position: {
        x: shape.position.x + dx,//这才是正宗的用类的字面量拷贝后覆盖。
        y: shape.position.y + dy,
      }
    });
  }
```



### 更新数组并且不产生mutation

|          | 避免使用 (会改变原始数组)     | 推荐使用 (会返回一个新数组）  |
| -------- | ----------------------------- | :---------------------------: |
| 添加元素 | `push`，`unshift`             | `concat`，`[...arr]` 展开语法 |
| 删除元素 | `pop`，`shift`，`splice`      |       `filter`，`slice`       |
| 替换元素 | `splice`，`arr[i] = ...` 赋值 |             `map`             |
| 排序     | `reverse`，`sort`             |       先将数组复制一份        |

- `slice` 让你可以拷贝数组或是数组的一部分。
- `splice` **会直接修改** 原始数组（插入或者删除元素）



```jsx
//———————————————————————————————————map来加工数组——————————————————————————————————————
let initialShapes = [
  { id: 0, type: 'circle', x: 50, y: 100 },
  { id: 1, type: 'square', x: 150, y: 100 },
  { id: 2, type: 'circle', x: 250, y: 100 },
];

export default function ShapeEditor() {
  const [shapes, setShapes] = useState(
    initialShapes
  );

  function handleClick() {
    const nextShapes = shapes.map(shape => {//返回一个加工后的数组赋值给另一个对象，加{}了所以要使用return
      if (shape.type === 'square') {//里面对数组每个元素遍历和加工。
        // 不作改变
        return shape;
      } else {
        // 返回一个新的圆形，位置在下方 50px 处
        return {
          ...shape,//复制一个新的数组，然后用字面量覆盖
          y: shape.y + 50,
        };
      }
    });
    // ————————————————————使用新的数组进行重渲染！！！
    setShapes(nextShapes);
  }

  return (
    <>
      <button onClick={handleClick}>
        所有圆形向下移动！
      </button>
      {shapes.map(shape => (
       ……
      ))}
    </>
  );
}
//计数器
  function handleIncrementClick(index) {
    const nextCounters = counters.map((c, i) => {//第一个是第n个变量，第二个是对应的索引n
      if (i === index) {
        // 递增被点击的计数器数值
        return c + 1;
      } else {
        // 其余部分不发生变化
        return c;
      }
    });
    setCounters(nextCounters);
  }

//————————————————————————————————像数组插入元素：就是切片操作——————————————————————————————
  function handleClick() {
    const insertAt = 1; // 可能是任何索引
    const nextArtists = [
      // 插入点之前的元素：
      ...artists.slice(0, insertAt),
      // 新的元素：
      { id: nextId++, name: name },
      // 插入点之后的元素：
      ...artists.slice(insertAt)
    ];
    setArtists(nextArtists);
    setName('');
  }
//——————————————————————————————————对象数组：
  function handleToggleMyList(artworkId, nextSeen) {
    setMyList(myList.map(artwork => {//对myList里的每个artwork对象遍历
      if (artwork.id === artworkId) {
        // 创建包含变更的*新*对象
        return { ...artwork, seen: nextSeen };//拷贝一个数组，然后覆盖
      } else {
        // 没有变更
        return artwork;
      }
    }));
  }
//用Immer改进
 function handleToggleMyList(id, nextSeen) {
    updateMyList(draft => {
      const artwork = draft.find(a =>
        a.id === id
      );
      artwork.seen = nextSeen;
    });
  }
//当使用 Immer 时，类似 artwork.seen = nextSeen 这种会产生 mutation 的语法不会再有任何问题了：
```

 const nextList = [...list];潜拷贝一个数组，reverse() 和 sort()

### 状态管理

```jsx
在组件间共享状态
有时候，你希望两个组件的状态始终同步更改。要实现这一点，可以将相关 state 从这两个组件上移除，并把 state 放到它们的公共父级，再通过 props 将 state 传递给这两个组件。这被称为“状态提升”，这是编写 React 代码时常做的事。
```

React 在移除一个组件时，也会销毁它的 state。

react会给UI构建渲染树！可以使用key来标识组件让react认出他们。

**只要一个组件还被渲染在 UI 树的相同位置，React 就会保留它的 state**（[对 state 进行保留和重置 – React 中文文档](https://zh-hans.react.dev/learn/preserving-and-resetting-state)）对 React 来说重要的是组件在 UI 树中的位置,而不是在 JSX 中的位置！渲染在UI树的同一个位置的组件会使用之前这个位置保留下来的state。不同组件就会使state重置。

```jsx
//不同位置
export default function App() {
  const counter = <Counter />;
  return (
    <div>
      {counter}
      {counter}//虽然是系统的标签，但这是渲染在UI树的不同位置！所以有各自的state！！
    </div>
  );
}
//不同位置，当没被渲染时state会随之消失！！！
<Counter />
{showB && <Counter />} 

//虽然有两个，但渲染只占同一个位置，state会被保留下来，系统组件才会保留stste，要是组件不相同就不能保留，eg把其中一个换成<p>state就会消失。或者使用其他标签包裹<Counter>也会被看成是不同的标签。
 {isPlayerA ? (
        <Counter person="Taylor" />
      ) : (
        <Counter person="Sarah" />
      )}
//————————————————————————————————————不同位置——————————————————————————————————————
	  {isPlayerA &&
        <Counter person="Taylor" />
      }
      {!isPlayerA &&
        <Counter person="Sarah" />
      }
//你可能在渲染列表时见到过 key。但 key 不只可以用于列表！还可以使用 key 来让 React 区分任何组件！！
{isPlayerA ? (
        <Counter key="Taylor" person="Taylor" />
      ) : (
        <Counter key="Sarah" person="Sarah" />
      )}//虽然这段代码看起来是在渲染同一个位置的同一个组件，但key不同就是不同组件。
//请记住 key 不是全局唯一的。它们只能指定 父组件内部 的顺序。

```

#### 为被移除的组件保留 state 的策略

在真正的聊天应用中，你可能会想在用户再次选择前一个收件人时恢复输入 state。对于一个不可见的组件，有几种方法可以让它的 state “活下去”：

- 与其只渲染现在这一个聊天，你可以把 **所有** 聊天都渲染出来，但用 CSS 把其他聊天隐藏起来。这些聊天就不会从树中被移除了，所以它们的内部 state 会被保留下来。这种解决方法对于简单 UI 非常有效。但如果要隐藏的树形结构很大且包含了大量的 DOM 节点，那么性能就会变得很差。
- 你可以进行 [状态提升](https://zh-hans.react.dev/learn/sharing-state-between-components) 并在父组件中保存每个收件人的草稿消息。这样即使子组件被移除了也无所谓，因为保留重要信息的是父组件。这是最常见的解决方法。
- 除了 React 的 state，你也可以使用其他数据源。例如，也许你希望即使用户不小心关闭页面也可以保存一份信息草稿。要实现这一点，你可以让 `Chat` 组件通过读取 [`localStorage`](https://developer.mozilla.org/zh-CN/docs/Web/API/Window/localStorage) 对其 state 进行初始化，并把草稿保存在那里
- **PageWrapper 组件**：这个组件负责在组件卸载时保存状态，并在组件挂载时恢复状态。

#### 迁移状态逻辑到reducer

使用 reducers 管理状态与直接设置状态略有不同。它不是通过设置状态来告诉 React “要做什么”，而是通过事件处理程序 dispatch 一个 “action” 来指明 “用户刚刚做了什么”。（而状态更新逻辑则保存在其他地方！）因此，我们不再通过事件处理器直接 “设置 `task`”，而是 dispatch 一个 “添加/修改/删除任务” 的 action。



reducer 函数就是你放置状态逻辑的地方。它接受两个参数，分别为当前 state 和 action 对象，并且返回的是更新后的 state：

```jsx
function yourReducer(state, action) {
  // 给 React 返回更新后的状态
  //可以使用switch case或者if来写
}
/*建议将每个 case 块包装到 { 和 } 花括号中，这样在不同 case 中声明的变量就不会互相冲突。此外，case 通常应该以 return 结尾。如果你忘了 return，代码就会 进入 到下一个 case，这就会导致错误*/
```

##### 在组件中使用 reducer 

就是将绑定事件的函数实现逻辑提取出来装入reducer，这些事件的函数只用来传参，有hook来……

需要将 `tasksReducer` 导入到组件中。记得先从 React 中导入 `useReducer` Hook：

```
import { useReducer } from 'react';
```

只需要像下面这样使用 `useReducer`:

```
const [tasks（状态值）, dispatch] = useReducer(tasksReducer（reducer函数）, initialTasks（初始化state）);
```

将之前state的状态逻辑就是哪些setXxx的地方改成dispatch

`useReducer` 和 `useState` 很相似——你必须给它传递一个初始状态，它会返回一个有状态的值和一个设置该状态的函数（在这个例子中就是 dispatch 函数）。但是，它们两个之间还是有点差异的。

`useReducer` 钩子接受 2 个参数：

1. 一个 reducer 函数
2. 一个初始的 state

它返回如下内容：

1. 一个有状态的值
2. 一个 dispatch 函数（用来 “派发” 用户操作给 reducer

```jsx
import { useReducer } from 'react';
import AddTask from './AddTask.js';
import TaskList from './TaskList.js';

export default function TaskApp() {
  const [tasks, dispatch] = useReducer(tasksReducer, initialTasks);//声明reducer

  function handleAddTask(text) {
    dispatch({//这里就给reducer传参的
      type: 'added',//reducer的类型，在reducer函数要用到
      id: nextId++,//其他是在reducer另外的参数
      text: text,
    });
  }

  function handleDeleteTask(taskId) {
  ……;
  }

  return (
 ……
  );
}

function tasksReducer(tasks, action) {
  switch (action.type) {
    case 'added': {//判断是什么动作！！
      return [
        ...tasks,
        {
          id: action.id,
          text: action.text,
          done: false,
        },
      ];
    }

    case 'deleted': {
    ……
    }
  }
}

let nextId = 3;//js的语法这里都能用，因为这里就是一个js文件，全局变量
const initialTasks = [
  {id: 0, text: '参观卡夫卡博物馆', done: true},
  ……
];
```

甚至可以把reducer放在其他文件只要export default /export就行!

##### 使用immer简化

```jsx
import { useImmerReducer } from 'use-immer';//导入immer
import AddTask from './AddTask.js';
import TaskList from './TaskList.js';

export default function TaskApp() {
  const [tasks, dispatch] = useImmerReducer(tasksReducer, initialTasks);

  function handleAddTask(text) {
    dispatch({
      type: 'added',
      id: nextId++,
      text: text,
    });
  }

  function handleChangeTask(task) {
    dispatch({
      type: 'changed',
      task: task,
    });
  }

  function handleDeleteTask(taskId) {
    dispatch({
……
  }

  return (
……
  );
}
function tasksReducer(draft, action) {//这里的参数我觉得没什么特别的，至于他是怎么知道改变哪个组件和这个函数的参数无关，里面只是根据action.type的值来确定哪个具体的function，里面的dispatch就可以确定哪个useImmerReducer，然后做出响应！
  switch (action.type) {
    case 'added': {
      draft.push({//
        id: action.id,
        text: action.text,
        done: false,
      });
      break;
    }
    case 'changed': {
      const index = draft.findIndex((t) => t.id === action.task.id);
      draft[index] = action.task;
      break;
    }
    case 'deleted': {
……
    }
  }
}
let nextId = 3;
const initialTasks = [
  {id: 0, text: '参观卡夫卡博物馆', done: true},
……
];

```

#### Context深层传递参数

Context 可以让父节点，甚至是很远的父节点都可以为其内部的整个组件树提供数据。

建议看[使用 Context 深层传递参数 – React 中文文档](https://zh-hans.react.dev/learn/passing-data-deeply-with-context)和[使用 Reducer 和 Context 拓展你的应用 – React 中文文档](https://zh-hans.react.dev/learn/scaling-up-with-reducer-and-context)

使用方法：

1. 通过 `export const MyContext = createContext(defaultValue)` 创建并导出 context。
2. 在无论层级多深的任何子组件中，把 context 传递给 `useContext(MyContext)` Hook 来读取它。
3. 在父组件中把 children 包在 `<MyContext.Provider value={...}>` 中来提供 context。

```jsx
//————————————————————————文件LevelContext.js——————————————————————
import { createContext } from 'react';
export const LevelContext = createContext(1);//在这里创建Context！
//—————————————————————————文件Heading.js——————————————————————————
import { useContext } from 'react';
import { LevelContext } from './LevelContext.js';//引入那个Context（全局生效）
export default function Heading({ children }) {
  const level = useContext(LevelContext);//在LevelContext获取那个我们想要的Context
  switch (level) {
    case 1:
      return <h1>{children}</h1>;
……
    case 6:
      return <h6>{children}</h6>;
    default:
      throw Error('未知的 level：' + level);
  }
} //Heading 使用 useContext(LevelContext) 访问上层最近的 LevelContext 提供的值。

//—————————————————————————文件Section.js——————————————————————————
import { useContext } from 'react';
import { LevelContext } from './LevelContext.js';
export default function Section({ children }) {
  const level = useContext(LevelContext);
  return (
    <section className="section">
      <LevelContext.Provider value={level + 1}>
        {children}
      </LevelContext.Provider>
    </section>
  );
}
/*LevelContext.Provider标签告诉 React：“如果在 <Section> 组件中的任何子组件请求 LevelContext，给他们这个 level。”组件会使用 UI 树中在它上层最近的那个 <LevelContext.Provider> 传递过来的值。*/

//—————————————————————————文件App.js——————————————————————————
import Heading from './Heading.js';
import Section from './Section.js';
export default function Page() {
  return (
    <Section>//这些就是children！
      <Heading>主标题</Heading>
      <Section>
        <Heading>副标题</Heading>
        <Heading>副标题</Heading>
        <Heading>副标题</Heading>
        <Section>
          <Heading>子标题</Heading>
          <Heading>子标题</Heading>
          <Heading>子标题</Heading>
          <Section>
            <Heading>子子标题</Heading>
            <Heading>子子标题</Heading>
            <Heading>子子标题</Heading>
          </Section>
        </Section>
      </Section>
    </Section>
  );
}
```



### 脱围机制

用ref组件不会在每次改变时重新渲染。 与 state 一样，React 会在每次重新渲染之间保留 ref。但是，设置 state 会重新渲染组件，更改 ref 不会！

当一条信息仅被事件处理器需要，并且更改它不需要重新渲染时，使用 ref 可能会更高效。

| ref                                                     | state                                                        |
| ------------------------------------------------------- | ------------------------------------------------------------ |
| `useRef(initialValue)`返回 `{ current: initialValue }`  | `useState(initialValue)` 返回 state 变量的当前值和一个 state 设置函数 ( `[value, setValue]`) |
| 更改时不会触发重新渲染                                  | 更改时触发重新渲染。                                         |
| 可变 —— 你可以在渲染过程之外修改和更新 `current` 的值。 | “不可变” —— 你必须使用 state 设置函数来修改 state 变量，从而排队重新渲染。 |
| 你不应在渲染期间读取（或写入） `current` 值。           | 你可以随时读取 state。但是，每次渲染都有自己不变的 state 快照。 |

每当你的组件重新渲染时（例如当你设置 state 时），组件重新渲染时，组件里的局部变量会重新初始化！所有局部变量都会从头开始初始化。如果组件在运行期间改变某些变量，下一次渲染时想使用他就需要使用ref，在组件里面定义ref。

#### 使用ref操作DOM（使用map的键值对维护）

```jsx
通过const myRef = useRef(null)和<div ref={myRef}>就可以获得DOM的引用。类似 querySelectorAll 
import { useRef } from 'react';

export default function Form() {
  const inputRef = useRef(null);

  function handleClick() {
    inputRef.current.focus();
  }	
  return (
    <>
      <input ref={inputRef} />//这告诉 React 将这个 <input> 的 DOM 节点放入
      <button onClick={handleClick}>
        聚焦输入框
      </button>
    </>
  );
}

//———————————————————————————————————————————————————————————————————————————————————————
ref 的数量是预先确定的。但有时候，你可能需要为列表中的每一项都绑定 ref ，而你又不知道会有多少项。像下面这样做是行不通的：

<ul>
  {items.map((item) => {
    // 行不通！
    const ref = useRef(null);
    return <li ref={ref} />;
  })}
</ul>
这是因为 Hook 只能在组件的顶层被调用。不能在循环语句、条件语句或 map() 函数中调用 useRef 。

解决方法是：解决方案是将函数传递给 ref 属性。这称为 ref 回调。当需要设置 ref 时，React 将传入 DOM 节点来调用你的 ref 回调，并在需要清除它时传入 null 。这使你可以维护自己的数组或 Map，并通过其索引或某种类型的 ID 访问任何 ref。
import { useRef, useState } from "react";
export default function CatFriends() {
    const itemsRef = useRef(null);
    const [catList, setCatList] = useState(setupCatList);
    function scrollToCat(cat) {
        const map = getMap();
        const node = map.get(cat);
        node.scrollIntoView({
            behavior: "smooth",
            block: "nearest",
            inline: "center",
        });
    }

    function getMap() {
        if (!itemsRef.current) {
            // 首次运行时初始化 Map。
            itemsRef.current = new Map();//直接给ref的值赋值
        }
        return itemsRef.current;
    }

    return (
        <>
            <nav>
                <button onClick={() => scrollToCat(catList[0])}>Tom</button>
                <button onClick={() => scrollToCat(catList[5])}>Maru</button>
                <button onClick={() => scrollToCat(catList[9])}>Jellylorum</button>
            </nav>
            <div>
                <ul>
                    {catList.map((cat) => (
                        <li
                            key={cat}
                            /*ref={(node) => {...}}: ref属性接收一个回调函数，这个回调函数的参数node是对应的DOM元素。当组件挂载（mount）或更新（update）时，React会调用这个回调函数，并传入DOM元素作为参数，如果组件卸载（unmount），则传入null。*/
                            ref={(node) => {
                                const map = getMap();
                                if (node) {
                                    map.set(cat, node);
                                } else {
                                    map.delete(cat);
                                }
                            }}
                            //通过维护节点信息和引用的map就可以在渲染列表时对列表的项进行跟踪。
                        >
                            <img src={cat} />
                        </li>
                    ))}
                </ul>
            </div>
        </>
    );
}

function setupCatList() {
……
}
```

上面的itemsRef.current = new Map();//直接给ref的值赋值

#### ref用于自己的组件

默认情况下，组件不暴露其 DOM 节点。 您可以通过使用 `forwardRef` 并将第二个 `ref` 参数传递给特定节点来暴露 DOM 节点

\<MyInput ref={inputRef} />会报错，默认只能用于html的标签。

非要用：

```jsx
import { forwardRef, useRef } from 'react';
//使用 forwardRef API：
const MyInput = forwardRef((props, ref) => {
  return <input {...props} ref={ref} />;
});
/*
	<MyInput ref={inputRef} /> 告诉 React 将对应的 DOM 节点放入 inputRef.current 中。但是，这取决于 MyInput 组件是否允许这种行为， 默认情况下是不允许的。
	MyInput 组件是使用 forwardRef 声明的。 这让从上面接收的 inputRef 作为第二个参数 ref 传入组件，第一个参数是 props 。
	MyInput 组件将自己接收到的 ref 传递给它内部的 <input>。
*/
export default function Form() {
  const inputRef = useRef(null);

  function handleClick() {
    inputRef.current.focus();
  }

  return (
    <>
      <MyInput ref={inputRef} />
      <button onClick={handleClick}>
        聚焦输入框
      </button>
    </>
  );
}
```

#### 用 flushSync 同步更新 state

在 React 中，每次更新都分为 [两个阶段](https://zh-hans.react.dev/learn/render-and-commit#step-3-react-commits-changes-to-the-dom)：

- 在 **渲染** 阶段， React 调用你的组件来确定屏幕上应该显示什么。
- 在 **提交** 阶段， React 把变更应用于 DOM。

所以在ref，state是在提交之后才更新的！在没提交之前是看不到更改后的状态的，所以需要异步更新才可以在没提交前看见更改状态。

有bug的代码：它的目的是添加一个新的待办事项，并将屏幕向下滚动到列表的最后一个子项。但因为上面的渲染和更新机制，它总是滚动到最后一个添加 **之前** 的待办事项。

```jsx
import { useState, useRef } from 'react';

export default function TodoList() {
  const listRef = useRef(null);
  const [text, setText] = useState('');
  const [todos, setTodos] = useState(
    initialTodos
  );

  function handleAdd() {
    const newTodo = { id: nextId++, text: text };
    setText('');
      //问题就是出在这里，他先渲染，此时已经将最新添加的渲染了。但是还没有提交
    setTodos([ ...todos, newTodo]);
      //所以下面的代码还是处于上一个
    listRef.current.lastChild.scrollIntoView({
      behavior: 'smooth',
      block: 'nearest'
    });
  }

  return (
    <>
      <button onClick={handleAdd}>
        添加
      </button>
      <input
        value={text}
        onChange={e => setText(e.target.value)}
      />
      <ul ref={listRef}>
        {todos.map(todo => (
          <li key={todo.id}>{todo.text}</li>
        ))}
      </ul>
    </>
  );
}

let nextId = 0;
let initialTodos = [];
for (let i = 0; i < 20; i++) {
  initialTodos.push({
    id: nextId++,
    text: '待办 #' + (i + 1)
  });
}

```

解决方法：

```jsx
flushSync(() => {
  setTodos([ ...todos, newTodo]);
});
listRef.current.lastChild.scrollIntoView();//指示 React 当封装在 flushSync 中的代码执行后，立即提交并同步更新 DOM
```

[使用 ref 操作 DOM – React 中文文档](https://zh-hans.react.dev/learn/manipulating-the-dom-with-refs)第三题，图像轮播！



### Effect

useEffect的执行是依靠回调的，不依赖任何顺序，即使写在return的前面，也是只有在触发回调的事件发生后才会执行。

事件处理函数只有在你再次执行同样的交互时才会重新运行。Effect 和事件处理函数不一样，它只有在读取的 props 或 state 值和上一次渲染不一样时才会重新同步。Effects 会在渲染后运行一些代码，以便可以将组件与 React 之外的某些系统同步。

```jsx
import { useEffect } from 'react';
function MyComponent() {
  useEffect(() => {
    // 每次渲染后都会执行此处的代码
  });
  return <div />;
}
//—————————————————————————Effect会在渲染完成后调用—————————————————————————————————
import { useState, useRef, useEffect } from 'react';

function VideoPlayer({ src, isPlaying }) {
  const ref = useRef(null);

  useEffect(() => {
    if (isPlaying) {
      console.log('调用 video.play()');
      ref.current.play(); // 没有useEffect时渲染期间不能调用 `play()`，因为第一次渲染时，ref还没被同步到这里，是一个null所以调用找不到对象。 
    } else {
      console.log('调用 video.pause()');
      ref.current.pause();// 同样，没有useEffect时调用 `pause()` 也不行。
    }//但是这里使用useEffect就可以使用play和pause
  },[isPlaying]);/*第二个参数是依赖数组 ，让effect的执行依赖于某个变量！
  指定 [isPlaying] 会告诉 React，如果 isPlaying 在上一次渲染时与当前相同，它应该跳过重新运行 Effect。通过这个改变，输入框的输入不会导致 Effect 重新运行，但是按下播放/暂停按钮会重新运行 Effect。*/

  return <video ref={ref} src={src} loop playsInline />;
}
/*当 VideoPlayer 组件渲染时（无论是否为首次渲染），都会发生以下事情。首先，React 会刷新屏幕，确保 <video> 元素已经正确地出现在 DOM 中；然后，React 将运行 Effect；最后，Effect 将根据 isPlaying 的值调用 play() 或 pause()。*/
export default function App() {
  const [isPlaying, setIsPlaying] = useState(false);
  const [text, setText] = useState('');
  return (
    <>
      <input value={text} onChange={e => setText(e.target.value)} />
      <button onClick={() => setIsPlaying(!isPlaying)}>
        {isPlaying ? '暂停' : '播放'}
      </button>
      <VideoPlayer
        isPlaying={isPlaying}
        src="https://interactive-examples.mdn.mozilla.net/media/cc0-videos/flower.mp4"
      />
    </>
  );
}
```

同一个组件里使用useState和useEffect，而且在effect里使用setXxx就会进入死循环。

effect会在每次渲染完成后执行，但有的时候不需要每次都执行，就像在输入框输入时onChange每次回触发渲染，但里面的effect不必每次都执行。

```jsx
useEffect(() => {
  // 没有依赖项数组，所以每次组件渲染后都会执行。里面的函数随便写
});

useEffect(() => {
  // 有一个空的依赖项数组，告诉 Effect 不依赖于响应式值，里面的代码没有响应式，在挂载后执行一次，后面有渲染不执行，return的东西在组件卸载后才执行。
}, []);//使用[]是effect代码块里面不能有state！否则要指明依赖哪个state，就像下面那样。

useEffect(() => {
  //这里的代码只会在每次渲染后，并且 a 或 b 的值与上次渲染不一致时执行。注意：多个变量中只要一个和上次不同就会被触发！
}, [a, b]);
```

#### 清理函数

每次重新执行 Effect 之前，React 都会调用清理函数；组件被卸载时，也会调用清理函数

为了解决连接没断开的问题可以在 Effect 中返回一个 **清理（cleanup）** 函数。

```jsx
  useEffect(() => {
    const connection = createConnection();
    connection.connect();
      //就是下面这一段，清理函数。
    return () => {
      connection.disconnect();
    };
  }, []);
```



**开发环境中effect会被react调用两次！**

在开发环境下，出现额外的连接、断连时，这是 React 正在调试你的代码。这是很正常的现象，可以选择关闭 [严格模式](https://zh-hans.react.dev/reference/react/StrictMode) 来关闭开发环境下特有的行为，但我们建议保留它。这可以帮助发现许多上面这样的错误。在开发环境中，React 有意重复挂载你的组件，以查找像上面示例中的错误。**正确的态度是“如何修复 Effect 以便它在重复挂载后能正常工作”，而不是“如何只运行一次 Effect”**

```jsx
//对于普通组件调用多次没什么区别，但某些 API 可能不允许连续调用两次。例如，内置的 <dialog> 元素的 showModal 方法在连续调用两次时会抛出异常，此时实现清理函数并使其关闭对话框：
useEffect(() => {
  const dialog = dialogRef.current;
  dialog.showModal();
  return () => dialog.close();
}, []);

//如果 Effect 订阅了某些事件，清理函数应该退订这些事件：
useEffect(() => {
  function handleScroll(e) {
    console.log(window.scrollX, window.scrollY);
  }
  window.addEventListener('scroll', handleScroll);
  return () => window.removeEventListener('scroll', handleScroll);
}, []);
//如果 Effect 对某些内容加入了动画，清理函数应将动画重置：
useEffect(() => {
  const node = ref.current;
  node.style.opacity = 1; // 触发动画
  return () => {
    node.style.opacity = 0; // 重置为初始值
  };
}, []);
```

#### 不需要effect的情况：

1. **如果一个值可以基于现有的 props 或 state 计算得出，[不要把它作为一个 state](https://zh-hans.react.dev/learn/choosing-the-state-structure#avoid-redundant-state)，而是在渲染期间直接计算这个值**。
2. [你可能不需要 Effect – React 中文文档](https://zh-hans.react.dev/learn/you-might-not-need-an-effect)

### 生命周期

每个 React 组件都经历相同的生命周期：

- 当组件被添加到屏幕上时，它会进行组件的 **挂载**。
- 当组件接收到新的 props 或 state 时，通常是作为对交互的响应，它会进行组件的 **更新**。
- 当组件从屏幕上移除时，它会进行组件的 **卸载**。



响应式的就是值随时可能改变的

```jsx
function ChatRoom({ roomId }) { // roomId 是响应式的
  const [serverUrl, setServerUrl] = useState('https://localhost:1234'); // serverUrl 是响应式的
  useEffect(() => {
    const connection = createConnection(serverUrl, roomId);
    connection.connect();
    return () => {
      connection.disconnect();
    };
  }, [serverUrl, roomId]); // ✅ 声明的所有依赖
  // ...
}

//——————————————————————————————————————————————————————————————————————————————————
const serverUrl = 'https://localhost:1234'; // serverUrl 不是响应式的
const roomId = 'general'; // roomId 不是响应式的

function ChatRoom() {
  useEffect(() => {
    const connection = createConnection(serverUrl, roomId);
    connection.connect();
    return () => {
      connection.disconnect();
    };
  }, []); // ✅ 声明的所有依赖
  // ...
}
```

如果组件里有多个响应式的值而且effect带[]的要在effect那里指明依赖哪个。

[响应式 Effect 的生命周期 – React 中文文档](https://zh-hans.react.dev/learn/lifecycle-of-reactive-effects)挑战第三题：

```jsx
export default function App() {
  const [position, setPosition] = useState({ x: 0, y: 0 });
  const [canMove, setCanMove] = useState(true);

  function handleMove(e) {
    if (canMove) {
      setPosition({ x: e.clientX, y: e.clientY });
    }
  }

  useEffect(() => {
    window.addEventListener('pointermove', handleMove);
    return () => window.removeEventListener('pointermove', handleMove);
  }, []);

  return (
    <>
      <label>
        <input type="checkbox"
          checked={canMove}
          onChange={e => setCanMove(e.target.checked)} 
        />
        是否允许移动
      </label>
      <hr />
      <div style={{
        position: 'absolute',
        backgroundColor: 'pink',
……
      }} />
    </>
  );
}
```

//这段代码失效的原是effect使用[]，在第一次渲染时canMove就是true的，<font color=red>[]在state更新渲染时不会重新渲染，所以这个值一直是true</font>卸载时才会return。所以要给effect添加依赖。



[将事件从 Effect 中分开 – React 中文文档](https://zh-hans.react.dev/learn/separating-events-from-effects)挑战第四题，在effect里面的延时函数没有传参在函数里使用到参数就会在组件里读取最新的，如果想使用旧的就得将旧的变量作为参数传给延时里的函数。



抑制 linter就是不规范使用[]而想要响应式生效。

[移除 Effect 依赖 – React 中文文档](https://zh-hans.react.dev/learn/removing-effect-dependencies)挑战第二题有动画渐变的代码，第三题展示了对象里面的值解包作为参数传给effect会更好，因为对象随时可能变化

### 自定义Hook

[使用自定义 Hook 复用逻辑 – React 中文文档](https://zh-hans.react.dev/learn/reusing-logic-with-custom-hooks)里面有一大堆在线离线检查。



hook的名称要以use开头，然后紧跟一个大写字母，就是驼峰命名。

自定义 Hook 共享的只是状态逻辑，不是状态本身。

你可以将响应值从一个 Hook 传到另一个，并且他们会保持最新。

每次组件重新渲染时，所有的 Hook 会重新运行。





hook就是对代码的复用，在多处功能相同的地方就可以提取代码使用hook。

但是每当你写 Effect 时，考虑一下把它包裹在自定义 Hook 是否更清晰。把 Effect 包裹进自定义 Hook 可以准确表达你的目标以及数据在里面是如何流动的。

```jsx
export default function Counter() {
  const [count, setCount] = useState(0);
  useEffect(() => {
    const id = setInterval(() => {
      setCount(c => c + 1);
    }, 1000);
    return () => clearInterval(id);
  }, []);
  return <h1>Seconds passed: {count}</h1>;
}
//——————————————————————————————————提取后————————————————————————————————————————————
将他用到的state和effect一起提取出来。
import { useState, useEffect } from 'react';

export function useCounter() {
  const [count, setCount] = useState(0);
  useEffect(() => {
    const id = setInterval(() => {
      setCount(c => c + 1);
    }, 1000);
    return () => clearInterval(id);
  }, []);
  return count;
}
//——————————
import { useCounter } from './useCounter.js';

export default function Counter() {
  const count = useCounter();
  return <h1>Seconds passed: {count}</h1>;
}
```

#### 跟随移动特效hook+effect+延时

```jsx
import { useState, useEffect } from 'react';

export function usePointerPosition() {
  const [position, setPosition] = useState({ x: 0, y: 0 });
  useEffect(() => {
    function handleMove(e) {
      setPosition({ x: e.clientX, y: e.clientY });
    }
    window.addEventListener('pointermove', handleMove);
    return () => window.removeEventListener('pointermove', handleMove);
  }, []);
  return position;
}
//——————————————————————————————————核心代码——————————————————————————————————————————
import { useState, useEffect } from 'react';
import { usePointerPosition } from './usePointerPosition.js';

function useDelayedValue(value, delay) {
  const [delayedValue, setDelayedValue] = useState(value);

  useEffect(() => {
    setTimeout(() => {
      setDelayedValue(value);//在依赖改变触发后，将旧值以参数的形式传给他，再次触发延时后就会使用旧值执行，不会到state读取最新的值。
    }, delay);
  }, [value, delay]);

  return delayedValue;
}

export default function Canvas() {
  const pos1 = usePointerPosition();
  const pos2 = useDelayedValue(pos1, 100);
  const pos3 = useDelayedValue(pos2, 200);
  const pos4 = useDelayedValue(pos3, 100);
  const pos5 = useDelayedValue(pos3, 50);
  return (
    <>
      <Dot position={pos1} opacity={1} />
      <Dot position={pos2} opacity={0.8} />
      <Dot position={pos3} opacity={0.6} />
      <Dot position={pos4} opacity={0.4} />
      <Dot position={pos5} opacity={0.2} />
    </>
  );
}

//opacity：不透明度
function Dot({ position, opacity }) {
  return (
    <div style={{
      position: 'absolute',
      backgroundColor: 'pink',
      borderRadius: '50%',
……
    }} />
  );
}
```



[React 项目中使用 `@` 符号代替 `src` 目录 - 掘金 (juejin.cn)](https://juejin.cn/post/6892636779850137608)



## 样式

**内联样式**：直接在JSX元素上使用`style`属性添加CSS样式。

```jsx
return (
    <>
        <img src="flowers/lily1.png" alt="百合花" style={{ width: '100px', height: 'auto' }} />
        {isLogin ? <Signin /> : <Register />}
    </>
)
```



**CSS样式表**：创建一个CSS文件，并在组件中导入它。

```jsx
/* SigninPage.css */
.imgStyle {
    width: 100px;
    height: auto;
}

import './SigninPage.css';

return (
    <>
        <img src="flowers/lily1.png" alt="百合花" className="imgStyle" />
        {isLogin ? <Signin /> : <Register />}
    </>
)
```



**CSS模块**：使用CSS模块技术来避免样式冲突。

```jsx
/* SigninPage.module.css */
.imgStyle {
    width: 100px;
    height: auto;
}

import styles from './SigninPage.module.css';

return (
    <>
        <img src="flowers/lily1.png" alt="百合花" className={styles.imgStyle} />
        {isLogin ? <Signin /> : <Register />}
    </>
)
```



**Styled-components**：使用`styled-components`库来在JavaScript中编写CSS。

```jsx
import styled from 'styled-components';

const StyledImg = styled.img`
    width: 100px;
    height: auto;
`;

return (
    <>
        <StyledImg src="flowers/lily1.png" alt="百合花" />
        {isLogin ? <Signin /> : <Register />}
    </>
)
```

mainfest.json文件

一个Web应用程序清单，它为Web应用提供了一种方式来告诉浏览器关于应用的信息，并且如何表现当它被添加到用户的主屏幕上。这个文件是JSON格式的，可以包含多种信息，比如应用的名称、图标、启动URL、显示模式以及主题颜色等。这些信息有助于浏览器将Web应用表现得更像一个原生应用。

以下是manifest.json文件中可能包含的一些关键字段及其作用：

- **`short_name`** 和 **`name`**: 应用的名称。`short_name`是应用的简短名称，主要在主屏幕显示；`name`是应用的完整名称，可能会在应用安装横幅或应用商店中使用。
- **`icons`**: 一组图标，这些图标用于在设备的主屏幕、任务切换器、应用商店等地方表示应用。可以指定多个大小以适应不同的设备和上下文。
- **`start_url`**: 指定当用户从主屏幕启动应用时，应用启动的URL。
- **`display`**: 控制浏览器界面的显示模式，比如是否显示为全屏、是否显示浏览器导航栏等。
- **`theme_color`** 和 **`background_color`**: 指定应用的主题颜色和背景颜色。这些颜色在浏览器界面中使用，比如地址栏的颜色。

通过提供这个文件，开发者可以控制应用添加到主屏幕后的外观和行为，提升用户体验，使Web应用更接近原生应用的体验。

## 箭头函数

箭头函数（Arrow functions）是JavaScript ES6中引入的一种写法，它提供了一种更简洁的方式来书写函数表达式。箭头函数在某些方面与传统函数有所不同，特别是在`this`的绑定方面。关于箭头函数的`return`，以下是几种情况：
### 当箭头函数没有`return`语句时：
1. **没有花括号**：如果箭头函数没有花括号，那么它默认返回表达式右侧的值。
   ```javascript
   let func = () => 'Hello World';  // 等同于 return 'Hello World';
   console.log(func()); // 输出：Hello World
   ```
2. **有花括号但没有return**：如果箭头函数有花括号但没有`return`语句，那么它默认返回`undefined`。
   ```javascript
   let func = () => {
       let a = 'Hello World';
   };
   console.log(func()); // 输出：undefined
   ```
### 当箭头函数有`return`语句时：
1. **有花括号和return**：如果箭头函数有花括号并且包含`return`语句，那么它会返回`return`后面的表达式的值。
   ```javascript
   let func = () => {
       return 'Hello World';
   };
   console.log(func()); // 输出：Hello World
   ```
2. **返回对象字面量**：如果你想要返回一个对象字面量，需要将对象字面量包裹在括号中，否则花括号会被解释为函数体的开始和结束。
   ```javascript
   let func = () => ({ key: 'value' });
   console.log(func()); // 输出：{ key: 'value' }
   ```
   了解箭头函数的这些`return`行为，可以帮助你更准确地控制函数的输出。



## public文件夹

以`/`开头的路径通常指的是相对于网站根目录的绝对路径。所以位于public文件夹下的文件直接使用，路径用'/'开头  "/path/to/file"

其他的使用相对路径，

- `./file` 是一个相对路径，表示当前目录下的 `file` 文件。`.` 符号代表当前目录，所以 `./file` 明确指出 `file` 文件位于当前目录。
- `file` 也是一个相对路径

脚手架的内置服务器devServer开启的服务就是localhost:3000，react的脚手架通过webpack配置public文件夹就是locahost:3000这台服务器的根路径。

## 路由

[快速学习Ant Design-布局-腾讯云开发者社区-腾讯云 (tencent.com)](https://cloud.tencent.com/developer/article/1583564)

一般组件使用路由组件api使用\<withRouter>

浏览器前进后退用到的就是history！

前端就是一个映射关系，通过改url的path来实现。前端路由就是根据path映射组件，后端路由根据path映射函数。

默认是模糊匹配，给多了可以匹配到，少了不行。要精准匹配，在注册路由时exact={true}，严格匹配不要随便开启，又时会导致不能匹配二级路由在一级路由下面，如果二级路由在二级路由映射的组件里，严格匹配导致没渲染一级组件，自然二级也不行。

react的路由注册是有顺序的，每次匹配时按注册的顺序匹配，所以多级路由要/xxx/yyy要些完整的路径。

#### 路由刷新样式丢失（路由多级结构）

检查样式的相对路径是否出现./...

[082_尚硅谷_react教程_解决样式丢失问题_哔哩哔哩_bilibili](https://www.bilibili.com/video/BV1wy4y1D7JT?p=82&spm_id_from=pageDriver&vd_source=a4cceb92ed79c160022ee4a50a573634)

#### 强制刷新

按shift+刷新：不走缓存，强制刷新。

内置组件

\<BrowserRotter>
\<HashRouter>
\<Route>
\<Redirect>
\<Link>
\<NavLink>
\<Switch>
其它
history 对象
match 对象
withRouter函数

要在路由前添加代码进行某些判断，你可以使用`react-router-dom的` `Route`组件或者是更高级的特性，如`wrappers`或`guards`，来实现条件渲染或执行某些操作。这里，我将展示如何使用一个简单的守卫（guard）组件来在路由渲染前进行判断。

步骤如下：

1. **创建一个守卫（Guard）组件**：这个组件将根据你的条件来决定是否渲染目标组件或重定向到其他页面。
2. **在路由配置中使用守卫组件**：将目标路由的[`element`](vscode-file://vscode-app/d:/Program/vscode/Microsoft VS Code/resources/app/out/vs/code/electron-sandbox/workbench/workbench.html)属性设置为守卫组件，并通过props传递目标组件。

### 守卫组件示例

```jsx
// src/components/RouteGuard.jsx
import React from 'react';
import { Navigate } from 'react-router-dom';

const RouteGuard = ({ condition, children, redirectTo }) => {
  if (!condition) {
    // 如果条件不满足，重定向到指定的路由
    return <Navigate to={redirectTo} replace />;
  }
  // 条件满足，渲染子组件
  return children;
};



// src/router/index.jsx
import { createBrowserRouter } from "react-router-dom";
import HomePage from "../hall/homePage";
import Login from "../login/loginPage";
import RouteGuard from "../components/RouteGuard";

// 假设有一个函数来判断用户是否登录
const isUserLoggedIn = () => {
  // 实现检查逻辑
  return true; // 或 false
};

const router = createBrowserRouter([
    {
        path: '/home',
        element: (
          <RouteGuard condition={isUserLoggedIn()} redirectTo="/">
            <HomePage />
          </RouteGuard>
        ),
    },
    {
        path: '/',
        element: <Login />,
    }
]);

export default router;
```

### 配置路由的方式（两种）

配置React Router的方式主要有两种，根据不同的版本和需求，可以选择不同的方式来实现路由功能。以下是详细说明：

### 1. 使用 `<BrowserRouter>` 或 `<HashRouter>` 组件（适用于 React Router v6 之前的版本）
这是传统的方式，通过在应用的顶层使用 `<BrowserRouter>` 或 `<HashRouter>` 组件来配置路由。

示例代码：

```javascript
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import HomePage from './HomePage';
import AboutPage from './AboutPage';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="about" element={<AboutPage />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
```

说明：

- **`<BrowserRouter>`** 使用 HTML5 的 history API，推荐用于常规的 web 应用。
- **`<HashRouter>`** 使用 URL 的 hash 部分（即 `window.location.hash`），适用于需要支持旧版本浏览器或者静态文件服务器的应用。

### 2. 使用 `<RouterProvider>` 和 `createBrowserRouter` 或 `createHashRouter`（适用于 React Router v6.4 及以上版本）
这是更现代的方式，通过创建一个路由器实例并使用 `<RouterProvider>` 提供这个路由器上下文。

示例代码：

```javascript
import { RouterProvider, createBrowserRouter } from 'react-router-dom';
import HomePage from './HomePage';
import AboutPage from './AboutPage';

const router = createBrowserRouter([
  {
    path: '/',
    element: <HomePage />,
  },
  {
    path: 'about',
    element: <AboutPage />,
  },
]);

function App() {
  return (
    <RouterProvider router={router} />
  );
}

export default App;
```

说明：

- **`createBrowserRouter`** 创建一个基于 HTML5 history API 的路由器。
- **`createHashRouter`** 创建一个基于 URL hash 部分的路由器。
- **`<RouterProvider>`** 组件用于提供路由器上下文，使得整个应用中的路由器配置可用。



### 不推荐同时使用

在一个React应用中，通常不需要同时使用`<RouterProvider>`和`<Routers>`组件。这是因为它们是不同的方式来配置和使用React Router。

- **`<RouterProvider>`**: 这是React Router v6.4+中的一种新方式，用来提供一个配置好的路由器给你的应用。你通常使用`createBrowserRouter`或`createHashRouter`来创建路由器，然后通过`<RouterProvider>`来提供这个路由器。

- **`<BrowserRouter>`** 和 **`<HashRouter>`**: 这些是React Router v6之前版本常用的路由组件，直接包裹在你的应用中来提供路由功能。

你不能同时使用`<RouterProvider>`和`<BrowserRouter>`，因为它们都提供了路由器的上下文，会造成冲突。通常建议选择一种方式并坚持使用它。

如果你使用的是React Router v6.4+，建议使用`<RouterProvider>`和`createBrowserRouter`来配置路由，因为它提供了一种更现代化、更灵活的方式来定义路由配置。如果你使用的是React Router v6之前的版本，继续使用`<BrowserRouter>`和`<Routes>`来配置路由。

### const

在JavaScript和React中，使用`const`关键字声明组件并不意味着组件是静态的。`const`关键字仅仅意味着该变量（在这种情况下是组件）的引用不可变，即你不能将`const`声明的变量重新赋值为另一个值或对象。然而，组件的行为（包括其渲染的内容）可以根据其接收的`props`或其内部状态（如果使用了状态钩子如`useState`）动态变化。



### 页面重定向/navigate组件重定向/Outlet嵌套路由

快速入门：[【React Router 6 快速上手一】重定向Navigate / useRoutes路由表 / 嵌套路由Outlet_useroutes 嵌套路由-CSDN博客](https://blog.csdn.net/xuxuii/article/details/126337516)

react专栏：[React_蜡笔雏田学前端的博客-CSDN博客](https://blog.csdn.net/xuxuii/category_11838680.html)

关于组件库：[React UI组件库——如何快速实现antd的按需引入和自定义主题_antd react定制主题-CSDN博客](https://xuxuii.blog.csdn.net/article/details/125831464)

路由的：[一文带你搞懂React路由（详解版）_react 路由-CSDN博客](https://blog.csdn.net/xuxuii/article/details/125691593?spm=1001.2014.3001.5502)

有关路由的组件或更深的子组件要被\<RouterProvider router={router} />（挂载到组件上的	）、\<BrowserRouter>等这些包围，路由操作才会生效。如果存在两个\<BrowserRouter>的话，这是两个路由，实际上整个应用要用一个路由管理。可以在root.render( <React.StrictMode>用\<BrowserRouter>包裹整个App就行。

写的时候做好将路由组件（靠路由导航的组件）和一般组件分开	

一个路径对应多个组件，那么组件都会渲染。单一匹配用\<Switch>\</Switch>包裹时，多个组件会渲染第一个



1. #### 编程式/命令式导航，使用useNavigate重定向页面

```jsx
import useNavigate  from 'react-router-dom';//导入重定向组件
export default function SigninPage() {

    //TODO：页面重定向的方法。
    const navigate = useNavigate();
……
/*已经改成个箭头函数！
  function redirectToHome() {
        navigate('/home', { replace: true }); // 使用navigate函数进行重定向
    }
*/
    return (
        <span style={{ position: "fixed", display: 'flex', alignItems: 'center' }}>
……
                <br />{/* 使用重定向的地方 */}
                <button onClick={() => navigate('/home', { replace: true })}>重定向到hall</button>

        </span>

    )
}
```

路由导航传参;

```jsx
searchParams传参
navigate('/article?id=1001&name=jack')
//组件从路径中回去参数。
const Article=()=>{
    const [params]=useSearchParams
    const id= params.get('id')
    const name = params.get('name')
    return <div>我是文章页{id}-{ngme}</div>
}

Params传参
//在路由配置占位符
{
    path:'/home/:id',
    element:<HomePage/>
}
//在组件获取
const Article=()={
    const params =useParams()
    const id = params.id
    return <div>我是文章页{id}</div>
}
```

传递state参数,state是一个对象所以{{}}，第一个{}是js表达式，第二个{}才是对象。

```jsx
{/*向路由组件传递state参数 */}
<Link to={{pathname:'/home/message/detail',state:{id:msgobj.id}}}>{msgobj.title}</Link>
```



2. #### url控制组件重定向

```jsx
// src/router/index.jsx
//创建一个路由配置
import { createBrowserRouter } from "react-router-dom";
import HomePage from "../hall/homePage";
import Login from "../login/loginPage";
import ErrorPage from "./errorPage";
const router = createBrowserRouter([
    {
        path: '/home',
        element: <HomePage />,
    },
    {
        path: '/',
        element: <Login />,
        errorElement: <ErrorPage />,
        //配置子路由
        children: [
            {
                path: "content",
                element: <ErrorPage />,
            },
        ],
    }
    ,
    {
        path: '*', // 通配符 代表除上面提到的路径，其他都会直接显示errorPage页面
        element: <ErrorPage />
    }

])

export default router

//————————————————————————————————将他挂载到某个组件————————————————————————————————
import React from 'react';
import { RouterProvider } from 'react-router-dom';
import router from './index.jsx'; // 确保这里的路径是指向你的路由配置文件的正确路径

function RouterApp() {
    return (
        <div className="App">
            <RouterProvider router={router} />//就是这段代码
        </其他组件>
        </div>
    );
}

export default RouterApp;
//——————————————————————————————渲染阶段——————————————————————————————————
import RouterApp from ……
root.render(
  <React.StrictMode>
    {/* < SigninPage /> */}
    {/* <ThreeScene /> */}
    {/* <HomePage /> */}
    <RouterApp />//这里
    <h6 style={{ textAlign: 'center', width: '100%' }}>作者：随便乱取</h6>
  </React.StrictMode>
)
/**这个是根组件，只能有一个，所以要用根组件</React.StrictMode>包裹所有的组件，这样才能渲染出来，里面的Router组件是路由组件，根据浏览器url的不同渲染
  不同的组件，这个组件是独立的，如果还写了其他组件的话，会两个一起渲染出来，只是路由组件会随着url的变化而变化，其他组件是不会变化的
  功能独立，各司其职*/

如果不用<BrowserRouter>或者<HashRouter>包裹，就挂在在某个组件上就用 <RouterProvider router={router} />写在<React.StrictMode>里面！
```

- 像下面这张图，也是根据路由来渲染组件\<Route path=" xxx"  component={不带<>的组件名}/>，redirect在全部没匹配时的重定向。多个\<Route>组件可以根据路径渲染其中符合条件的component，如果有多个符合就会去一起渲染，使用\<Switch>组件可以避免这个问题，只会渲染第一个符合条件的component。

![image-20240716180048278](D:\Program\Typora\My picture library\image-20240716180048278.png)





3. #### Link:声明式导航（在配置好路由表之后才生效！）
   
    声明式导航是指通过在模版中通过`<Link/>`组件描述出要跳转到哪里去，比如后台管理系统的左侧菜单通常使用这种方式进行
    `<Link to="/article">文章</Link>`
    语法说明:通过给组件的to属性指定要跳转到路由path，组件会被渲染为浏览器支持的a链接，如果需要传参直接通过字符串拼接的方式拼接参数即可
4. #### NavLink：Link的升级版

`<NavLink activeClassName="自定义样式" className="其他class" to="/PATH" >XXX</NavLink> `：这个默认点击导航到这个路由时样式设置为active，可以自定义样式，记得将权限提到最高

```css
.demo{
    background-clor:xxxxx  !impoortant;
    color:white !important;	
}
```

封装

```jsx
创建一个名为 `MyNavLink` 的自定义组件，该组件扩展了 `Component` 类，并且使用了 `NavLink` 组件来封装一个导航链接。

 `to` 属性应该从 `this.props` 中解构出来，并且传递给 `NavLink`。
 `About` 看起来是硬编码的文本，如果你想要它可配置，应该将其作为一个prop传递给 `MyNavLink`。


import React, { Component } from 'react';
import { NavLink } from 'react-router-dom'; // 确保导入了NavLink
export default class MyNavLink extends Component {
  render() {
    const { to, children } = this.props; // 解构to和children（或其他需要的props）
    
    return (
      <NavLink
        activeClassName="atguigu" // 当NavLink处于活动状态时应用的类
        className="list-group-item" // 基础类名
        to={to} // 链接的目标位置
      >
        {children} {/* 使用children prop来显示传递给MyNavLink的子元素 */}
      </NavLink>
    );
  }
}
现在，你可以这样使用 `MyNavLink` 组件：

<MyNavLink to="/about">About</MyNavLink>
这将渲染一个导航链接，当链接处于活动状态时，会添加 `atguigu` 类，并且始终具有 `list-group-item` 类。`About` 文本将作为子元素显示。如果你想要传递其他props到 `NavLink`，你可以在 `MyNavLink` 组件中相应地传递它们。
```

传递的参数可以用解构{...this.popps}

3. 嵌套路由

    ![image-20240716110555508](D:\Program\Typora\My picture library\image-20240716110555508.png)

    ![image-20240716111326182](D:\Program\Typora\My picture library\image-20240716111326182.png)

    访问一级路由时，二级路由可默认渲染，只要将path去掉换成index：true就行。

    ![image-20240716114448998](D:\Program\Typora\My picture library\image-20240716114448998.png)

    

    5. #### Navigate实现路由重定向：当渲染这个组件时就触发重定向！修改路径，更换视图。

    `<Navigate to="/xxx" replace={true} /> `：

    to：前往路由  

    replace：是否替换历史堆栈中的当前条目，浏览器浏览记录是栈结构！replace的true替换栈顶，false时push历史记录

    - 当 `replace` 设置为 `true` 时，导航会将当前的历史记录条目替换为新的条目。这意味着用户无法通过浏览器的前进和后退按钮返回到被替换的页面。
    - 当 `replace` 设置为 `false`（默认值）时，导航会在历史记录中添加一个新的条目，用户可以通过浏览器的前进和后退按钮在新的页面和之前的页面之间切换。
        简而言之，`replace={true}` 的作用是在导航到新页面时，不保留当前页面的历史记录，相当于使用 `history.replace()` 方法而不是 `history.push()` 方法。这在某些场景下很有用，比如在用户登录后，你可能不希望用户能够回到登录页面。

既可以在组件处重定向，也可以在router结构里重定向！

在router配置里重定向,引用2的代码，当输入/tohome就会但想到/home

```jsx
import { createBrowserRouter } from "react-router-dom";
import HomePage from "../hall/homePage";
import Login from "../login/loginPage";
import ErrorPage from "./errorPage";
import { Navigate } from "react-router-dom";
const router = createBrowserRouter([
    {
        path: '/tohome',
        element: <Navigate to="/home" replace={false} />,
        errorElement: <ErrorPage />,
    },
    {
        path: '/home',
        element: <HomePage />,
    },
……

])

export default router
```

在组件重定向（控制渲染时机）

```jsx
export default function SigninPage() {

    const [reden, setReden] = useState(false);
……
    return (
        <span style={{ position: "fixed", display: 'flex', alignItems: 'center' }}>
……
                <button onClick={() => setReden(true)}>渲染重定向到hall</button>
                <Navigate to={reden ? '/home' : '/'} replace={true} />
//或者： {reden ? <Navigate to='/home' replace={false} /> : null}//要将三目运算写在{}里，要不然不会生效！
        </span>

    )
}

//——————————————————————————————————————————————————使用return也可以！

  if (shouldRedirect) {
    // 使用 Navigate 组件进行跳转，这里跳转到 '/new-location'
    return <Navigate to="/new-location" replace />;
  }
```

## antd

文档3.x最详细。

暴露配置，看package.json中"eject": "react-scripts eject"，命令行输入npm  eject就可以看见配置文件。

[095_尚硅谷_react教程_antd样式的按需引入_哔哩哔哩_bilibili](https://www.bilibili.com/video/BV1wy4y1D7JT?p=95&vd_source=a4cceb92ed79c160022ee4a50a573634)





onChange={(e) => setPassword(e.target.value)

在JavaScript中，`e.target`是一个事件属性，它返回触发事件的元素。在React的事件处理中，`e`代表事件对象，而`e.target`通常用于获取触发事件的DOM元素。例如，在一个输入框（`<input>`）的`onChange`事件中，`e.target`指的就是这个输入框元素本身，因此`e.target.value`就是输入框当前的值。



在React应用中定义和使用环境变量通常涉及以下步骤：
## 创建`.env`文件
在你的React项目根目录（和src同级）下创建一个名为`.env`的文件。在这个文件中，你可以定义你的环境变量，格式为`变量名=值`。

文件名规则：

- `.env`： 默认。
- `.env.local`：本地覆盖。**为除测试之外的所有环境加载此文件。**
- `.env.development`, `.env.test`, `.env.production`：特定于环境的设置。
- `.env.development.local`, `.env.test.local`, `.env.production.local`：本地覆盖特定于环境的设置。

不同文件名，在执行命令时的优先级，从左到右，`左边优先局最高`：

- `$ npm start`: `.env.development.local` > `.env.local` > `.env.development` > `.env`
- `$ npm run build`: `.env.production.local` > `.env.local` > `.env.production` > `.env`
- `$ npm test`: `.env.test.local` > `.env.test` > `.env`（注意 `.env.local` 不包含了）

```env
REACT_APP_MY_VARIABLE="my value"
```
注意：变量名必须以`REACT_APP_`为前缀，这是因为Create React App（一个常用的React应用脚手架工具）默认只加载以`REACT_APP_`开头的环境变量到`process.env`。

在代码中可以通过：consloe.log( process.env.REACT_APP_VARIABLE );来访问到环境变量。

[React .env 环境变量（详细使用、命名方式） - 掘金 (juejin.cn)](https://juejin.cn/post/7221717935613624378)

<font color=red size=4>每次对环境变量修改后都要重启项目更改才能生效</font>

在 .env 文件中，环境变量的值无需强制使用引号包围，除非值中包含空格。在您提供的例子中，值 `http://localhost:8881/upload` 不包含空格，因此不需要使用引号。但是，使用引号包围值在某些情况下可以增加可读性，尤其是当值中包含特殊字符时。如果您想要为了一致性或可读性而使用引号，这是完全可以的，但不是必需的。



## 生命周期

[【React】深入理解React组件生命周期----图文详解（含代码）_react官方生命周期图解在哪-CSDN博客](https://blog.csdn.net/xuxuii/article/details/125053304)

[一看就懂的React生命周期 - 掘金 (juejin.cn)](https://juejin.cn/post/7285540804734468150)

[尚硅谷React生命周期](D:/the_files_at/Assignment/笔记资料/尚硅谷区块链/以太坊理论资料/以太坊理论资料/pdf/11_React生命周期简介.pdf)	




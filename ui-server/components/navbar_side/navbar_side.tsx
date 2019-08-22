import React from 'react'
import MenuItem from './components/menu_item';

const NavbarSide: React.SFC<{}> = () => {
  return <div className="fixed left-0 top-0 lg:flex flex-col hidden h-screen w-56 px-3 bg-white border shadow z-50">
    <h3 className="w-full mb-8 text-3xl font-bold text-indigo-700 leading-relaxed text-center select-none">
      phenopod
    </h3>
    <ul>
      <li className="my-3"><MenuItem icon="explore" name="discover" active/></li>
      <li className="my-3"><MenuItem icon="feed" name="feed"/></li>
      <li className="my-3"><MenuItem icon="heart" name="subscriptions"/></li>
      <li className="my-3"><MenuItem icon="user-solid-circle" name="profile"/></li>
    </ul>
  </div>
}

export default NavbarSide
import React from 'react'

const NavbarBottom: React.SFC<{}> = () => (
  <section className="menu lg:hidden xs:flex justify-around h-full pt-1 pb-0 bg-white z-50"> 
    <button className="w-24 text-green-600">
        <svg className="h-5 w-auto mx-auto fill-current" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20"><path d="M10 20a10 10 0 1 1 0-20 10 10 0 0 1 0 20zM7.88 7.88l-3.54 7.78 7.78-3.54 3.54-7.78-7.78 3.54zM10 11a1 1 0 1 1 0-2 1 1 0 0 1 0 2z"/></svg>
        <h4 className="capitalize text-xs text-center leading-loose underline">discover</h4> 
    </button>
    <button className="w-24 text-gray-700">
        <svg className="h-5 w-auto mx-auto fill-current" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20"><path d="M6 6V2c0-1.1.9-2 2-2h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-4v4a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V8c0-1.1.9-2 2-2h4zm2 0h4a2 2 0 0 1 2 2v4h4V2H8v4zM2 8v10h10V8H2z"/></svg>
        <h4 className="capitalize text-xs text-center leading-loose">feed</h4> 
    </button>
    <button className="w-24 text-gray-700">
        <svg className="h-5 w-auto mx-auto fill-current" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20"><path d="M10 3.22l-.61-.6a5.5 5.5 0 0 0-7.78 7.77L10 18.78l8.39-8.4a5.5 5.5 0 0 0-7.78-7.77l-.61.61z"/></svg>
        <h4 className="capitalize text-xs text-center leading-loose">subscriptions</h4> 
    </button>
    <button className="w-24 text-gray-700">
        <svg className="h-5 w-auto mx-auto fill-current" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20"><path d="M10 20a10 10 0 1 1 0-20 10 10 0 0 1 0 20zM7 6v2a3 3 0 1 0 6 0V6a3 3 0 1 0-6 0zm-3.65 8.44a8 8 0 0 0 13.3 0 15.94 15.94 0 0 0-13.3 0z"/></svg>
        <h4 className="capitalize text-xs text-center leading-loose">profile</h4> 
    </button>
  </section>
)

export default NavbarBottom
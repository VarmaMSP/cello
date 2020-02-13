import About from 'components/about'
import React from 'react'
import Categories from './categories'
import Recommended from './recommended'

const HomeView: React.FC<{}> = () => {
  return (
    <>
      <Recommended />
      <Categories />
      <div className="md:hidden text-center">
        <hr className="my-5" />
        <About />
      </div>
    </>
  )
}

export default HomeView

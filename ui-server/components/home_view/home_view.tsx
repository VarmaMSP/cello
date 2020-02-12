import About from 'components/about'
import CategoryList from 'components/category_list'
import React from 'react'
import Recommended from './recommended'

const HomeView: React.FC<{}> = () => {
  return (
    <div>
      <Recommended />

      <h1 className="pb-1 text-xl tracking-wide font-semibold">
        {'Browse By Categories'}
      </h1>
      <hr className="mb-3" />
      <div className="category-list">
        <CategoryList className="md:w-1/4 w-full" />
      </div>

      <div className="md:hidden text-center">
        <hr className="my-5" />
        <About />
      </div>
    </div>
  )
}

export default HomeView

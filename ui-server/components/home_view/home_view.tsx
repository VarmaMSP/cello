import About from 'components/about'
import CategoryList from 'components/category_list'
import React from 'react'
import Recommended from './recommended'

const HomeView: React.FC<{}> = () => {
  return (
    <div>
      <Recommended />

      <h1 className="pb-1 text-xl tracking-wide text-gray-800 font-medium">
        {'Browse By Categories'}
      </h1>
      <hr className="mb-3" />
      <div className="flex flex-col flex-wrap" style={{ height: '50rem' }}>
        <CategoryList className="w-1/4" />
      </div>

      <div className="md:hidden text-center">
        <hr className="my-5" />
        <About />
      </div>
    </div>
  )
}

export default HomeView

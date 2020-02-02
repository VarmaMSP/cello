import { ChartLink } from 'components/link'
import React from 'react'
import { Curation } from 'types/app'

export interface StateToProps {
  categories: Curation[]
}

const Categories: React.FC<StateToProps> = ({ categories }) => {
  return (
    <>
      <h1 className="pb-1 text-xl tracking-wide text-gray-800 font-medium">
        {'Browse By Categories'}
      </h1>
      <hr className="mb-3" />
      <div className="md:w-2/3 w-full">
        <ul className="flex flex-wrap">
          {categories.map((c) => (
            <ChartLink key={c.id} chartId={c.id}>
              <a className="block h-10 lg:w-1/2 w-full pl-6 py-1 rounded-full hover:bg-gray-200">
                <li className="text-lg text-gray-900">{c.title}</li>
              </a>
            </ChartLink>
          ))}
        </ul>
      </div>
    </>
  )
}

export default Categories

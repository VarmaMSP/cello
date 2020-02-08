import classnames from 'classnames'
import { ChartLink } from 'components/link'
import React from 'react'
import { Category } from 'types/models'
import { formatCategoryTitle } from 'utils/format'

export interface StateToProps {
  categories: Category[]
}

export interface OwnProps {
  className: string
}

const CategoryList: React.FC<StateToProps & OwnProps> = ({
  categories,
  className,
}) => {
  const byId = categories.reduce<{ [categoryId: string]: Category }>(
    (acc, c) => ({ ...acc, [c.id]: c }),
    {},
  )
  const byParentId = categories.reduce<{ [categoryId: string]: string[] }>(
    (acc, p) =>
      !!p.parentId
        ? acc
        : {
            ...acc,
            [p.id]: categories.reduce<string[]>(
              (acc, c) => (c.parentId !== p.id ? acc : [...acc, c.id]),
              [],
            ),
          },
    {},
  )

  return (
    <>
      {Object.keys(byParentId).map((parentId) => {
        const parent = byId[parentId]
        const childIds = byParentId[parentId]

        if (childIds.length === 0) {
          return (
            <ChartLink
              chartId={formatCategoryTitle(parent.name)}
              key={parentId}
            >
              <a className={classnames('my-2 font-medium', className)}>
                {parent.name}
              </a>
            </ChartLink>
          )
        }

        return (
          <div key={parentId} className={classnames('my-2', className)}>
            <ChartLink
              chartId={formatCategoryTitle(parent.name)}
              key={parentId}
            >
              <a className="block font-medium mb-1">{parent.name}</a>
            </ChartLink>

            <ul className="list-disc list-inside">
              {childIds.map((childId) => (
                <li key={childId} className="pl-4 text-xs leading-relaxed">
                  <ChartLink chartId={formatCategoryTitle(byId[childId].name)}>
                    <a>{byId[childId].name}</a>
                  </ChartLink>
                </li>
              ))}
            </ul>
          </div>
        )
      })}
    </>
  )
}

export default CategoryList

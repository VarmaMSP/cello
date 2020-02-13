import classnames from 'classnames'
import { ChartLink } from 'components/link'
import { Category } from 'models'
import React from 'react'

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
            <ChartLink chartUrlParam={parent.urlParam} key={parentId}>
              <a
                className={classnames(
                  'block my-2 text-black tracking-wide font-medium',
                  className,
                )}
              >
                {parent.name}
              </a>
            </ChartLink>
          )
        }

        return (
          <div key={parentId} className={classnames('my-2', className)}>
            <ChartLink chartUrlParam={parent.urlParam} key={parentId}>
              <a className="block text-black tracking-wide font-medium mb-1">
                {parent.name}
              </a>
            </ChartLink>

            <ul className="list-disc list-inside">
              {childIds.map((childId) => (
                <li
                  key={childId}
                  className="pl-4 text-sm leading-relaxed tracking-wide"
                >
                  <ChartLink chartUrlParam={byId[childId].urlParam}>
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

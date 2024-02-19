-- name: FindAllMeals :many
select * from meal;

-- name: FindMealsWithName :many
select * from meal m where m.name like $1;

-- name: FindOneMealById :one
select * from meal m where m.id = $1;

-- name: FindMealVariants :many
select mv.*, m.* from meal_variant mv inner join macro m on mv.id = m.meal_variant_id where mv.meal_id = $1;

-- name: FindMacro :one
select * from macro where meal_variant_id = $1;

-- name: FindIngredients :many
select i.amount, i.unit, i.snack, p.name from ingredient i inner join product p on p.id = i.product_id where i.meal_variant_id = $1;

-- name: FindAllProducts :many
select * from product;

-- name: CreatePlan :one
insert into plan(date) values($1) returning *;

-- name: AddMealToPlan :exec
insert into plan_meal(plan_id, meal_variant_id) values($1, $2);

-- name: FindPlanByDate :one
select * from plan where date = $1;

-- name: FindMealByProductName :many
select meal.*, mv.kcal from meal
left join meal_variant mv on meal.id = mv.meal_id
left join ingredient i on mv.id = i.meal_variant_id
left join product p on i.product_id = p.id
where lower(p.name) like $1;

-- name: FindMealVariantsByDate :many
select pm.meal_variant_id, mv.kcal, m.proteins, m.fats, m.carbs, m.fiber, meal.name, meal.description from plan
left join public.plan_meal pm on plan.id = pm.plan_id
left join public.meal_variant mv on pm.meal_variant_id = mv.id
left join public.meal on mv.meal_id = meal.id
left join public.macro m on mv.id = m.meal_variant_id
where plan.date = $1;

-- name: FindIngredientsByMealVariantId :many
select p.name, i.amount, i.unit, i.snack from ingredient i
left join public.product p on i.product_id = p.id
where i.meal_variant_id = $1;

-- name: FindAllMealsVariantsWithKcal :many
select m.name, mv.id as meal_variant_id, mv.kcal from meal m
inner join public.meal_variant mv on m.id = mv.meal_id ORDER BY m.id;
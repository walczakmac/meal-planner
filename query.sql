-- name: FindAllMeals :many
select * from meal;

-- name: FindMealVariants :many
select * from meal_variant where meal_id = $1;

-- name: FindMacro :one
select * from macro where meal_variant_id = $1;

-- name: FindIngredients :many
select i.*, p.name as product_name from ingredient i left join product p on p.id = i.product_id where i.meal_variant_id = $1;

-- name: FindAllProducts :many
select * from product;

-- name: FindMealsWithName :many
select * from meal m where m.name like $1;

-- name: CreatePlan :one
insert into plan(date) values($1) returning *;

-- name: AddMealToPlan :exec
insert into plan_meal(plan_id, meal_variant_id) values($1, $2);

-- name: FindPlanByDate :one
select * from plan where date = $1;
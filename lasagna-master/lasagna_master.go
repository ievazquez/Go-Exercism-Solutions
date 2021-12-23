package lasagna

//const
const (
	noodlesPerLayer  = 50
	saucePerLayer    = 0.2
	standardPortions = 2.0
)

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, minutes int ) int {
    if minutes == 0 {
        minutes = 2
    }
    return len(layers) * minutes
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string ) (int, float64) {
    var grams int = 0
    var liters float64
    for _, ingredient:= range layers {
        switch ingredient {
            case "sauce":
                 liters += saucePerLayer
            case "noodles":
                 grams += noodlesPerLayer
        }
    }
    return grams,liters
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendlist, ownlist []string ) []string {
    return append(ownlist, friendlist[len(friendlist)-1])
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe( quantities []float64, portions int) []float64 {
    var newQuantities = make([]float64, len(quantities))
    copy(newQuantities, quantities)
    var multiplier = float64(portions) / standardPortions
    for i:= 0 ; i < len(newQuantities) ; i++{
        newQuantities[i] *= multiplier
    }
    return newQuantities
}

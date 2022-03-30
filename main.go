package main

import (
	"gsdc/letsfix/controllers"
	//"gsdc/letsfix/models"
	"gsdc/letsfix/repository"
	"gsdc/letsfix/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	userRepository repository.UserRepository = repository.NewUserRepository()
	userService service.UserService = service.New(userRepository)
	userController controllers.UserController = controllers.New(userService)

	deviceRepository repository.DeviceRepository = repository.NewDeviceRepository()
	deviceService service.DeviceService = service.NewDeviceService(deviceRepository)
	deviceController controllers.DeviceController = controllers.NewDeviceController(deviceService)

	ownershipRepository repository.OwnershipRepository = repository.NewOwnershipRepository()
	ownershipService service.OwnershipService = service.NewOwnershipService(ownershipRepository)
	ownershipController controllers.OwnershipController = controllers.NewOwnershipController(ownershipService)

	typeRepository repository.TypeRepository = repository.NewTypeRepository()
	typeService service.TypeService = service.NewTypeService(typeRepository)
	typeController controllers.TypeController = controllers.NewTypeController(typeService)

	brandRepository repository.BrandRepository = repository.NewBrandRepository()
	brandService service.BrandService = service.NewBrandService(brandRepository)
	brandController controllers.BrandController = controllers.NewBrandController(brandService)

	recyclerRepository repository.RecyclerRepository = repository.NewRecyclerRepository()
	recyclerService service.RecyclerService = service.NewRecyclerService(recyclerRepository)
	recyclerController controllers.RecyclerController = controllers.NewRecyclerController(recyclerService)

	repairerRepository repository.RepairerRepository = repository.NewRepairerRepository()
	repairerService service.RepairerService = service.NewRepairerService(repairerRepository)
	repairerController controllers.RepairerController = controllers.NewRepairerController(repairerService)

	rctypeRepository repository.RctypeRepository = repository.NewRctypeRepository()
	rctypeService service.RctypeService = service.NewRctypeService(rctypeRepository)
	rctypeController controllers.RctypeController = controllers.NewRctypeController(rctypeService)

	rptypeRepository repository.RptypeRepository = repository.NewRptypeRepository()
	rptypeService service.RptypeService = service.NewRptypeService(rptypeRepository)
	rptypeController controllers.RptypeController = controllers.NewRptypeController(rptypeService)

)

func main() {
	r := gin.Default()

	repository.DatabaseConnect()

	// Users CRUD
	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, userController.FindAll())
	})

	r.POST("/users", func(ctx *gin.Context) {
		err := userController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "User Input is Valid!"})
		}
	})
	
	r.GET("/users/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, userController.FindByUserId(ctx))
	})

	r.PUT("/users/:id", func(ctx *gin.Context) {
		err := userController.Update(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "User Update is Valid!"})
		}
	})

	// Devices CRUD

	r.GET("/devices", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, deviceController.FindAllDevices())
	})

	r.GET("/devices/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, deviceController.FindDevice(ctx))
	})

	r.POST("/devices", func(ctx *gin.Context) {
		err := deviceController.SaveDevice(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "User Input is Valid!"})
		}
	})

	// Ownership CRUD
	r.GET("/ownerships", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, ownershipController.FindAllOwnerships())
	})

	r.GET("/ownerships/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, ownershipController.FindOwnershipByUserId(ctx))
	})

	r.GET("/ownerships/devices/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, ownershipController.FindDevicesByUserId(ctx))
	})

	r.POST("/ownerships", func(ctx *gin.Context) {
		err := ownershipController.SaveOwnership(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "User Input is Valid!"})
		}
	})	

	// Type CRUD
	r.GET("/types", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, typeController.FindAllTypes())
	})

	r.GET("/types/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, typeController.FindByTypeId(ctx))
	})

	r.POST("/types", func(ctx *gin.Context) {
		err := typeController.SaveType(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "User Input is Valid!"})
		}
	})

	// Brand CRUD
	r.GET("/brands", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, brandController.FindAllBrands())
	})

	r.GET("/brands/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, brandController.FindByBrandId(ctx))
	})

	r.POST("/brands", func(ctx *gin.Context) {
		err := brandController.SaveBrand(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "User Input is Valid!"})
		}
	})

	// Recycler CRUD
	r.GET("/recyclers", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, recyclerController.FindAllRecyclers())
	})

	r.GET("/recyclers/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, recyclerController.FindByRecyclerId(ctx))
	})

	r.POST("/recyclers", func(ctx *gin.Context) {
		err := recyclerController.SaveRecycler(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "User Input is Valid!"})
		}
	})

	// Repairer CRUD
	r.GET("/repairers", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, repairerController.FindAllRepairers())
	})

	r.GET("/repairers/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, repairerController.FindByRepairerId(ctx))
	})

	r.POST("/repairers", func(ctx *gin.Context) {
		err := repairerController.SaveRepairer(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "User Input is Valid!"})
		}
	})

	// Recycler type CRUD
	r.GET("/recyclertypes", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, rctypeController.FindAllRctypes())
	})

	r.GET("/recyclertypes/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, rctypeController.FindRctypeById(ctx))
	})

	r.GET("/recyclertypes/types/:type_name", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, rctypeController.FindRctypeByType(ctx))
	})

	r.POST("/recyclertypes", func(ctx *gin.Context) {
		err := rctypeController.SaveRctype(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "User Input is Valid!"})
		}
	})

	// Repairer type CRUD
	r.GET("/repairertypes", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, rptypeController.FindAllRptypes())
	})

	r.GET("/repairertypes/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, rptypeController.FindRptypeById(ctx))
	})

	r.GET("/repairertypes/repairers/:repairer", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, rptypeController.FindRptypeByRepairer(ctx))
	})

	r.GET("/repairertypes/types/:type_name", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, rptypeController.FindRptypeByType(ctx))
	})

	r.GET("/repairertypes/brands/:brand_name", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, rptypeController.FindRptypeByBrand(ctx))
	})

	r.GET("/repairertypes/components/:component", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, rptypeController.FindRptypeByComponent(ctx))
	})

	r.POST("/repairertypes", func(ctx *gin.Context) {
		err := rptypeController.SaveRptype(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "User Input is Valid!"})
		}
	})

	r.Run()
}

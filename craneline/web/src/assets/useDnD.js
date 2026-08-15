import { useVueFlow } from '@vue-flow/core'
import { ref, watch } from 'vue'

let id = 0

/**
 * @returns {string} - A unique id.
 */
function getId() {
  return `dndnode_${id++}`
}

/**
 * In a real world scenario you'd want to avoid creating refs in a global scope like this as they might not be cleaned up properly.
 * @type {{draggedType: Ref<string|null>, isDragOver: Ref<boolean>, isDragging: Ref<boolean>}}
 */
const state = {
  /**
   * The type of the node being dragged.
   */
  draggedType: ref(null),
  isDragOver: ref(false),
  isDragging: ref(false),
}

export default function useDragAndDrop() {
  const { draggedType, isDragOver, isDragging } = state

  const { addNodes, screenToFlowCoordinate, onNodesInitialized, updateNode } = useVueFlow()

  watch(isDragging, (dragging) => {
    document.body.style.userSelect = dragging ? 'none' : ''
  })

  function onDragStart(event, type) {
    if (event.dataTransfer) {
      event.dataTransfer.setData('application/vueflow', type)
      event.dataTransfer.effectAllowed = 'move'
    }

    draggedType.value = type
    isDragging.value = true

    document.addEventListener('drop', onDragEnd)
  }

  /**
   * Handles the drag over event.
   *
   * @param {DragEvent} event
   */
  function onDragOver(event) {
    event.preventDefault()

    if (draggedType.value) {
      isDragOver.value = true

      if (event.dataTransfer) {
        event.dataTransfer.dropEffect = 'move'
      }
    }
  }

  function onDragLeave() {
    isDragOver.value = false
  }

  function onDragEnd() {
    isDragging.value = false
    isDragOver.value = false
    draggedType.value = null
    document.removeEventListener('drop', onDragEnd)
  }

  /**
   * Handles the drop event.
   *
   * @param {DragEvent} event
   */
  // function onDrop(event) {
  //   const position = screenToFlowCoordinate({
  //     x: event.clientX,
  //     y: event.clientY,
  //   })

  //   const nodeId = getId()

  //   const newNode = {
  //     id: nodeId,
  //     type: draggedType.value,
  //     position,
  //     data: { label: nodeId },
  //   }

  //   /**
  //    * Align node position after drop, so it's centered to the mouse
  //    *
  //    * We can hook into events even in a callback, and we can remove the event listener after it's been called.
  //    */
  //   const { off } = onNodesInitialized(() => {
  //     updateNode(nodeId, (node) => ({
  //       position: { x: node.position.x - node.dimensions.width / 2, y: node.position.y - node.dimensions.height / 2 },
  //     }))

  //     off()
  //   })

  //   addNodes(newNode)
  // }
  function onDrop(event) {
  const raw = event.dataTransfer?.getData('application/vueflow')
  if (!raw) return

  const image = JSON.parse(raw)
  // const position = screenToFlowCoordinate({ x: event.clientX - left, y: event.clientY - top })
  const position = screenToFlowCoordinate({
      x: event.clientX,
      y: event.clientY,
    })

  // taille dynamique selon le contenu, ex: en fonction du nombre de ports à afficher
  const height = 60 + image.ports.length * 20

  const newNode = {
    id: getId(),
    type: 'dockerImage', // un seul type de composant custom pour toutes les images
    position,
    data: {
      label: image.name,
      tag: image.tag,
      ports: image.ports,
      category: image.category,
      description: image.description,
    },
    style: {
      width: '200px',
      height: `${height}px`,
    },
  }

  addNodes(newNode)
}

  return {
    draggedType,
    isDragOver,
    isDragging,
    onDragStart,
    onDragLeave,
    onDragOver,
    onDrop,
  }
}
